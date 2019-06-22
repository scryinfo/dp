pragma solidity ^0.4.24;

import "./common.sol";
import "../ScryToken.sol";

library verification {
    event RegisterVerifier(string seqNo, address[] users);
    event Vote(string seqNo, uint256 transactionId, bool judge, string comments, uint8 state, uint8 index, address[] users);
    event VerifierDisable(string seqNo, address verifier, address[] users);

    function register(
        common.Verifiers storage self,
        common.Configuration storage conf,
        string seqNo,
        ERC20 token) public {

        common.Verifier storage v = getVerifier(self, msg.sender);
        require( v.addr == 0x00, "Address already registered");

        //deposit
        if (conf.verifierDepositToken > 0) {
            require(token.balanceOf(msg.sender) >= conf.verifierDepositToken, "No enough balance");
            require(token.transferFrom(msg.sender, address(this), conf.verifierDepositToken), "Pay deposit failed");
        }

        self.list[self.list.length++] = common.Verifier(msg.sender, conf.verifierDepositToken, 0, 0, true);
        self.validVerifierCount++;

        address[] memory users = new address[](1);
        users[0] = msg.sender;
        emit RegisterVerifier(seqNo, users);
    }

    function vote(
        common.Verifiers storage self,
        common.TransactionData storage txData,
        common.VoteData storage voteData,
        common.Configuration storage conf,
        string seqNo,
        uint txId,
        bool judge,
        string comments,
        ERC20 token) public {

        common.TransactionItem storage txItem = txData.map[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.state == common.TransactionState.Created || txItem.state == common.TransactionState.Voted, "Invalid transaction state");

        bool valid;
        uint8 index;
        common.Verifier storage verifier = getVerifier(self, msg.sender);
        (valid, index) = verifierValid(verifier, txItem.verifiers);
        require(valid, "Invalid verifier");

        if (!voteData.map[txId][msg.sender].used) {
            payToVerifier(txItem, conf, verifier.addr, token);
        }
        voteData.map[txId][msg.sender] = common.VoteResult(judge, comments, true);

        txItem.state = common.TransactionState.Voted;
        txItem.creditGived[index] = false;

        address[] memory users = new address[](1);
        users[0] = txItem.buyer;
        emit Vote(seqNo, txId, judge, comments, uint8(txItem.state), index+1, users);

        users[0] = msg.sender;
        emit Vote(seqNo, txId, judge, comments, uint8(txItem.state), 0, users);
    }

    function creditsToVerifier(
        common.Verifiers storage self,
        common.PublishedData storage pubData,
        common.TransactionData storage txData,
        common.Configuration storage conf,
        string seqNo,
        uint256 txId,
        uint8 verifierIndex,
        uint8 credit) public {

        //validate
        require(credit >= conf.creditLow && credit <= conf.creditHigh, "0 <= credit <= 5 is valid");

        common.TransactionItem storage txItem = txData.map[txId];
        require(txItem.used, "Transaction does not exist");

        common.Verifier storage verifier = getVerifier(self, txItem.verifiers[verifierIndex]);
        require(verifier.addr != 0x00, "Verifier does not exist");

        common.DataInfoPublished storage data = pubData.map[txItem.publishId];
        require(data.used, "Publish data does not exist");
        require(txItem.needVerify, "Transaction can't verify");

        bool valid;
        uint256 index;
        (valid, index) = verifierValid(verifier, txItem.verifiers);
        require(valid, "Invalid verifier");
        require(!txItem.creditGived[index], "This verifier is credited");

        verifier.credits = (uint8)((verifier.credits * verifier.creditTimes + credit)/(verifier.creditTimes+1));
        verifier.creditTimes++;
        txItem.creditGived[index] = true;

        address[] memory users = new address[](1);
        users[0] = address(0x00);
        //disable verifier and forfeiture deposit while credit <= creditThreshold
        if (verifier.credits <= conf.creditThreshold) {
            verifier.enable = false;
            verifier.deposit = 0;
            self.validVerifierCount--;
            require(self.validVerifierCount >= 1, "Invalid verifier count");

            emit VerifierDisable(seqNo, verifier.addr, users);
        }
    }


    function chooseVerifiers(common.Verifiers storage self, common.Configuration storage conf) internal view returns (address[] memory) {
        uint256 len = self.list.length;
        address[] memory chosenVerifiers = new address[](conf.verifierNum);

        for (uint8 i = 0; i < conf.verifierNum; i++) {
            uint256 index = uint256(keccak256(abi.encodePacked(now, msg.sender))) % len;
            common.Verifier memory v = self.list[index];

            //loop if invalid verifier was chosen until get valid verifier
            address vb = v.addr;
            while (!v.enable || addressExist(chosenVerifiers, v.addr)) {
                v = self.list[(++index) % len];
                require(v.addr != vb, "Disordered verifiers");
            }

            chosenVerifiers[i] = v.addr;
        }

        return chosenVerifiers;
    }

    function addressExist(address[] addrArray, address addr) internal pure returns (bool exist) {
        for (uint8 i = 0; i < addrArray.length; i++) {
            if (addr == addrArray[i]) {
                exist = true;
                break;
            }
        }

        return ;
    }

    function verifierValid(common.Verifier v, address[] arr) pure internal returns (bool, uint8) {
        bool exist;
        uint8 index;

        (exist, index) = addressIndex(arr, v.addr);
        return (v.enable && exist, index);
    }

    function verifierExist(address addr, address[] arr) pure internal returns (bool) {
        bool exist;
        (exist, ) = addressIndex(arr, addr);

        return exist;
    }

    function addressIndex(address[] addrArray, address addr) pure internal returns (bool, uint8) {
        for (uint8 i = 0; i < addrArray.length; i++) {
            if (addrArray[i] == addr) {
                return (true, i);
            }
        }

        return (false, 0);
    }

    function payToVerifier(common.TransactionItem storage txItem, common.Configuration storage conf, address verifier, ERC20 token) internal {
        if (txItem.buyerDeposit >= conf.verifierBonus) {
            txItem.buyerDeposit -= conf.verifierBonus;

            if (!token.transfer(verifier, conf.verifierBonus)) {
                txItem.buyerDeposit += conf.verifierBonus;
                require(false, "Failed to pay to verifier");
            }
        } else {
            require(false, "Low deposit value for paying to verifier");
        }
    }

    function getVerifier(common.Verifiers storage self, address v) internal view returns (common.Verifier storage){
        for (uint256 i = 0; i < self.list.length; i++) {
            if (self.list[i].addr == v) {
                return self.list[i];
            }
        }

        return self.list[0];
    }

    function chooseArbitrators(common.Verifiers storage self, common.Configuration storage conf, address[] vs) internal view returns (address[] memory) {
        uint256[] memory shortlistIndex = new uint256[](self.list.length - conf.verifierNum);
        uint256 count;

        for (uint256 i = 0;i < self.list.length;i++) {
            common.Verifier storage a = self.list[i];
            if (arbitratorValid(a, conf, vs)) {
                shortlistIndex[count] = i;
                count++;
            }
        }
        require(count >= conf.arbitratorNum, "No enough arbitrators");

        uint256 shortlistLen = count;

        address[] memory shortlist = new address[](count);
        while (count > 0) {
            count--;
            shortlist[count] = self.list[shortlistIndex[count]].addr;
        }

        address[] memory chosenArbitrators = new address[](conf.arbitratorNum);
        for (i = 0;i < conf.arbitratorNum;i++) {
            uint256 index = uint256(keccak256(abi.encodePacked(now, msg.sender))) % shortlistLen;
            address ad = shortlist[index];
            while (addressExist(chosenArbitrators, ad)) {
                ad = shortlist[(++index) % shortlistLen];
            }
            chosenArbitrators[i] = ad;
        }

        return chosenArbitrators;
    }

    function arbitratorValid(common.Verifier storage a, common.Configuration storage conf, address[] vs) internal view returns (bool) {
        bool notVerifier = true;
        for (uint8 i = 0;i < vs.length;i++) {
            if (a.addr == vs[i]) {
                notVerifier = false;
                break;
            }
        }
        return notVerifier && a.enable && (a.credits >= conf.arbitrateCredit);
    }

    function arbitrate(common.TransactionData storage txData, common.ArbitratorData storage abData,
        common.Configuration storage conf, uint256 txId, bool judge, ERC20 token) internal {
        common.TransactionItem storage txItem = txData.map[txId];
        require(txItem.used, "Transaction does not exist");

        bool exist;
        uint8 index;
        (exist, index) = addressIndex(txItem.arbitrators, msg.sender);
        require(exist, "Invalid arbitrator");

        require(!abData.map[txId][index].used, "Address already arbitrated");

        payToArbitrator(txItem, conf, msg.sender, token);
        abData.map[txId][index] = common.ArbitratorResult(msg.sender, judge, true);
    }

    function arbitrateFinished(common.ArbitratorData storage abData, common.Configuration storage conf, uint256 txId) internal returns (bool) {
        bool finish = true;
        for (uint8 i = 0; i < conf.arbitratorNum; i++) {
            if (!abData.map[txId][i].used) {
                finish = false;
                break;
            }
        }

        return finish;
    }

    function payToArbitrator(common.TransactionItem storage txItem, common.Configuration storage conf, address arbitrator, ERC20 token) internal {
        if (txItem.buyerDeposit >= conf.arbitratorBonus) {
            txItem.buyerDeposit -= conf.arbitratorBonus;

            if (!token.transfer(arbitrator, conf.arbitratorBonus)) {
                txItem.buyerDeposit += conf.arbitratorBonus;
                require(false, "Failed to pay to verifier");
            }
        } else {
            require(false, "Low deposit value for paying to verifier");
        }
    }
}
