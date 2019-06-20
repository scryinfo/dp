pragma solidity ^0.4.24;

import "./common.sol";
import "../ScryToken.sol";

library verification {
    event RegisterVerifier(string seqNo, address[] users);
    event VerifiersChosen(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, common.TransactionState state, address[] users);
    event Vote(string seqNo, uint256 transactionId, bool judge, string comments, common.TransactionState state, uint8 index, address[] users);
    event VerifierDisable(string seqNo, address verifier, address[] users);

    function register(
        common.Verifiers storage self,
        common.Configuration storage conf,
        string seqNo,
        ERC20 token) external {

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
        string seqNo,
        uint txId,
        bool judge,
        string comments,
        ERC20 token) external {

        common.TransactionItem storage txItem = txData.map[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.state == common.TransactionState.Created || txItem.state == common.TransactionState.Voted, "Invalid transaction state");

        bool valid;
        uint8 index;
        common.Verifier storage verifier = getVerifier(self, msg.sender);
        (valid, index) = verifierValid(verifier, txItem.verifiers);
        require(valid, "Invalid verifier");

        if (!voteData.map[txId][msg.sender].used) {
            payToVerifier(txItem, verifier.addr, token);
        }
        voteData.map[txId][msg.sender] = common.VoteResult(judge, comments, true);

        txItem.state = common.TransactionState.Voted;
        txItem.creditGived[index] = false;

        address[] memory users = new address[](1);
        users[0] = txItem.buyer;
        emit Vote(seqNo, txId, judge, comments, txItem.state, index+1, users);
    }

    function creditsToVerifier(
        common.Verifiers storage self,
        common.PublishedData storage pubData,
        common.TransactionData storage txData,
        common.Configuration storage conf,
        string seqNo,
        uint256 txId,
        uint8 verifierIndex,
        uint8 credit) external {

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

        (exist, index) = getVerifierIndex(v.addr, arr);
        return (v.enable && exist, index);
    }

    function verifierExist(address addr, address[] arr) pure internal returns (bool) {
        bool exist;
        (exist, ) = getVerifierIndex(addr, arr);

        return exist;
    }

    function getVerifierIndex(address verifier, address[] arrayVerifier) pure internal returns (bool, uint8) {
        for (uint8 i = 0; i < arrayVerifier.length; i++) {
            if (arrayVerifier[i] == verifier) {
                return (true, i);
            }
        }

        return (false, 0);
    }

    function payToVerifier(common.TransactionItem storage txItem, address verifier, ERC20 token) internal {
        if (txItem.buyerDeposit >= txItem.verifierBonus) {
            txItem.buyerDeposit -= txItem.verifierBonus;

            if (!token.transfer(verifier, txItem.verifierBonus)) {
                txItem.buyerDeposit += txItem.verifierBonus;
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
}
