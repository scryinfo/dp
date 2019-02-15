pragma solidity ^0.4.24;

import "./ScryToken.sol";

contract ScryProtocol {
    enum TransactionState {Begin, Created, Voted, Buying, ReadyForDownload, Arbitrating, Payed, Closed}

    struct DataInfoPublished {
        uint256 price;
        bytes metaDataIdEncSeller;
        bytes32[] proofDataIds;
        uint256 numberOfProof;
        string despDataId;
        address seller;
        bool supportVerify;
        bool used;
    }

    struct TransactionItem {
        TransactionState state;
        address buyer;
        address seller;
        address[] verifiers;
        bool[] creditGived;
        string publishId;
        bytes meteDataIdEncBuyer;
        bytes metaDataIdEncSeller;
        uint256 buyerDeposit;
        bool used;
    }

    struct Verifier {
        address addr;
        uint256 deposit;
        uint8 credits;  //credit: 0-5
        uint256 creditTimes;
        bool enable;
    }

    uint8 validVerifierCount = 0;
    uint8 verifierNum = 2;
    uint256 verifierDepositToken = 10000;
    uint8 creditLow = 0;
    uint8 creditHigh = 5;
    uint8 creditThreshold = 2;
    uint8 arbitratorNum = 1;
    uint8 arbitrateCredit = 0;

    struct VoteResult {
        bool judge;
        string comments;
        bool used;
    }

    struct ArbitrateResult {
        bool judge;
    }

    Verifier[] verifiers;
    mapping (string => DataInfoPublished) mapPublishedData;
    mapping (uint256 => TransactionItem) mapTransaction;
    mapping (uint256 => mapping(address => VoteResult)) mapVote;
    mapping (uint256 => mapping(address => ArbitrateResult)) mapArbitrate;
    mapping (uint256 => bool[]) mapCount;

    event RegisterVerifier(string seqNo, address[] users);
    event DataPublish(string seqNo, string publishId, uint256 price, string despDataId, address[] users);
    event TransactionCreate(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bool supportVerify, address[] users);
    event VerifiersChosen(string seqNo, uint256 transactionId, bytes32[] proofIds, address[] users);
    event Vote(string seqNo, uint256 transactionId, bool judge, string comments, address[] users);
    event ArbitratorsChosen(string seqNo,uint256 transactionId,address[] users);
    event Arbitrate(string seqNo,uint256 txId,address[] users);
    event Buy(string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, address[] users);
    event ReadyForDownload(string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, address[] users);
    event TransactionClose(string seqNo, uint256 transactionId, address[] users);
    event VerifierDisable(string seqNo, address verifier, address[] users);

    uint256 transactionSeq = 0;
    uint256 encryptedIdLen = 32;

    address token_address = 0x0;
    address owner         = 0x0;
    ERC20   token;

    constructor (address _token) public {
        require(_token != 0x0);

        owner = msg.sender;
        token_address = _token;
        token = ERC20(_token);

        //the first element used for empty usage
        verifiers[verifiers.length++] = Verifier(0x00, 0, 0, 0, false);
    }

    function registerAsVerifier(string seqNo) external {
        Verifier storage v = getVerifier(msg.sender);
        require( v.addr == 0x00, "The verifier is already registered");

        //deposit
        if (verifierDepositToken > 0) {
            require(token.balanceOf(msg.sender) >= verifierDepositToken, "No enough balance");
            require(token.transferFrom(msg.sender, address(this), verifierDepositToken), "Failed to transfer token from caller");
        }

        verifiers[verifiers.length++] = Verifier(msg.sender, verifierDepositToken, creditHigh, 0, true);
        validVerifierCount++;

        address[] memory users = new address[](1);
        users[0] = msg.sender;
        emit RegisterVerifier(seqNo, users);
    }

    function publishDataInfo(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller,
        bytes32[] proofDataIds, string despDataId, bool supportVerify) external {
        address[] memory users = new address[](1);
        users[0] = address(0x00);

        DataInfoPublished storage data = mapPublishedData[publishId];
        require(!data.used, "Duplicate publish id");

        mapPublishedData[publishId] = DataInfoPublished(price, metaDataIdEncSeller, proofDataIds, proofDataIds.length, despDataId, msg.sender, supportVerify, true);

        emit DataPublish(seqNo, publishId, price, despDataId, users);
    }

    function isPublishedDataExisted(string publishId) internal view returns (bool) {
        DataInfoPublished storage data = mapPublishedData[publishId];
        return data.used;
    }

    function createTransaction(string seqNo, string publishId) external {
        //published data info
        DataInfoPublished memory data = mapPublishedData[publishId];
        require(data.used, "Publish data does not exist");
        require((token.balanceOf(msg.sender)) >= data.price, "No enough balance");
        uint256 fee = data.price;
        if (data.supportVerify) {
            fee += 0;   // should add the tokens which for awarding verifiers and arbitrators.
        }
        require(token.transferFrom(msg.sender, address(this), data.price), "Failed to transfer token from caller");

        //create transaction
        uint txId = getTransactionId();
        bytes memory metaDataIdEncBuyer = new bytes(encryptedIdLen);

        address[] memory selectedVerifiers;
        bool[] memory creditGived;
        if (data.supportVerify) {
            //choose verifiers randomly
            selectedVerifiers = chooseVerifiers(verifierNum);
            creditGived = new bool[](verifierNum);
            emit VerifiersChosen(seqNo, txId,data.proofDataIds, selectedVerifiers);

//            address[] memory selectedArbitrators;
//            selectedArbitrators = chooseArbitrators(arbitratorNum, selectedVerifiers);
//            emit ArbitratorsChosen(seqNo, txId, selectedArbitrators);
        }

        address[] memory users = new address[](1);
        users[0] = msg.sender;
        emit TransactionCreate(seqNo, txId, publishId, data.proofDataIds, data.supportVerify, users);

        mapTransaction[txId] = TransactionItem(TransactionState.Created, msg.sender, data.seller, selectedVerifiers, creditGived,
            publishId, metaDataIdEncBuyer, data.metaDataIdEncSeller, data.price, true);
    }

    function chooseVerifiers(uint8 num) internal view returns (address[] memory) {
        require(num < validVerifierCount, "No enough valid verifiers");
        address[] memory chosenVerifiers = new address[](num);

        for (uint8 i = 0; i < num; i++) {
            uint index = getRandomNumber(verifiers.length) % verifiers.length;
            Verifier storage v = verifiers[index];

            //loop if invalid verifier was chosen until get valid verifier
            address vb = v.addr;
            while (!v.enable || verifierExist(v.addr, chosenVerifiers)) {
                v = verifiers[(++index) % verifiers.length];
                require(v.addr != vb, "Disordered verifiers");
            }

            chosenVerifiers[i] = v.addr;
        }

        return chosenVerifiers;
    }

    function verifierValid(Verifier v, address[] arr) pure internal returns (bool, uint256) {
        bool exist;
        uint256 index;

        (exist, index) = getVerifierIndex(v.addr, arr);
        return (v.enable && exist, index);
    }

    function verifierExist(address addr, address[] arr) pure internal returns (bool) {
        bool exist;
        (exist, ) = getVerifierIndex(addr, arr);

        return exist;
    }

    function getVerifierIndex(address verifier, address[] arrayVerifier) pure internal returns (bool, uint256) {
        for (uint256 i = 0; i < arrayVerifier.length; i++) {
            if (arrayVerifier[i] == verifier) {
                return (true, i);
            }
        }

        return (false, 0);
    }

    function getTransactionId() internal returns(uint) {
        return transactionSeq++;
    }

    function getRandomNumber(uint mod) internal view returns (uint) {
        return uint(keccak256(now, msg.sender)) % mod;
    }

    function vote(string seqNo, uint txId, bool judge, string comments) external {
        TransactionItem storage txItem = mapTransaction[txId];
        require(txItem.used, "Transaction does not exist");

        bool valid;
        uint256 index;
        Verifier storage verifier = getVerifier(msg.sender);
        (valid, index) = verifierValid(verifier, txItem.verifiers);
        require(valid, "Invalid verifier");

        mapVote[txId][msg.sender] = VoteResult(judge, comments, true);

        txItem.state = TransactionState.Voted;
        txItem.creditGived[index] = false;

        address[] memory users = new address[](1);
        users[0] = txItem.buyer;
        emit Vote(seqNo, txId, judge, comments, users);
    }

    function chooseArbitrators(uint8 num, address[] vs) internal view returns (address[] memory) {
        address[] memory shortlist = new address[](num * 3);
        address[] memory chosenArbitrators = new address[](num);
        uint256 count;

        for (uint256 i = 0;i < verifiers.length && count < num * 3;i++) {
            Verifier storage a = verifiers[i];
            if (arbitratorValid(a.addr,vs)) {
                shortlist[count] = a.addr;
                count++;
            }
        }
        require(count >= num, "Not enough arbitrators");

        for (i = 0;i < num;) {
            uint index = getRandomNumber(shortlist.length) % shortlist.length;
            address ad = shortlist[index];
            if (!arbitratorExist(ad,shortlist)) {
                chosenArbitrators[i] = a.addr;
                i++;
            }
        }

        return chosenArbitrators;
    }

    function arbitratorValid(address addr, address[] vs) internal pure returns (bool) {
        bool notVerifier = true;
        for (uint8 i = 0;i < vs.length;i++) {
            if (addr == vs[i]) {
                return false;
            }
        }
//        return notVerifier && a.enable && (a.credits >= arbitrateCredit);
        return notVerifier;
    }

    function arbitratorExist(address addr, address[] shortlist) internal pure returns (bool) {
        for (uint256 i = 0;i < shortlist.length;i++) {
            if (addr == shortlist[i]) {
                return true;
            }
        }
        return false;
    }

    function buyData(string seqNo, uint256 txId) external {
        //validate
        TransactionItem storage txItem = mapTransaction[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.buyer == msg.sender, "Invalid buyer");

        DataInfoPublished memory data = mapPublishedData[txItem.publishId];
        require(data.used, "Publish data does not exist");

        if(!data.supportVerify) {
            require(txItem.state == TransactionState.Created, "Invalid transaction state");
        } else {
            require(txItem.state == TransactionState.Voted, "Invalid transaction state");
        }

        txItem.state = TransactionState.Buying;

        address[] memory users = new address[](1);
        users[0] = txItem.seller;
        emit Buy(seqNo, txId, txItem.publishId, txItem.metaDataIdEncSeller, users);
    }

    function submitMetaDataIdEncWithBuyer(string seqNo, uint256 txId, bytes encryptedMetaDataId) external {
        //validate
        TransactionItem storage txItem = mapTransaction[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.seller == msg.sender, "Invalid seller");
        require(txItem.state == TransactionState.Buying, "Invalid transaction state");

        txItem.meteDataIdEncBuyer = encryptedMetaDataId;
        txItem.state = TransactionState.ReadyForDownload;

        //ReadyForDownload event
        address[] memory users = new address[](1);
        users[0] = txItem.buyer;
        emit ReadyForDownload(seqNo, txId, txItem.metaDataIdEncSeller, users);
    }

    function confirmDataTruth(string seqNo, uint256 txId, bool truth) external {
        //validate
        TransactionItem storage txItem = mapTransaction[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.buyer == msg.sender, "Invalid buyer");

        DataInfoPublished storage data = mapPublishedData[txItem.publishId];
        require(data.used, "Publish data does not exist");

        require(txItem.state == TransactionState.ReadyForDownload);
        if (!data.supportVerify) {
            if(truth) {
                payToSeller(txItem, data);
            }

            closeTransaction(txItem, seqNo, txId);
        } else {
            if (truth) {
                payToSeller(txItem, data);
                closeTransaction(txItem, seqNo, txId);
            } else {
                // arbitrate.
            }
        }
    }

    function arbitrate(string seqNo,uint txId,bool judge) external {
        TransactionItem storage txItem = mapTransaction[txId];
        require(txItem.used, "Transaction does not exist");

        mapArbitrate[txId][msg.sender] = ArbitrateResult(judge);
        uint256 c = mapCount[txId].length;
        mapCount[txId][c] = judge;

        txItem.state = TransactionState.Arbitrating;

        if (mapCount[txId].length == arbitratorNum) {
            uint8 truth;
            for (uint256 i = 0;i < arbitratorNum;i++) {
                if (mapCount[txId][i]) {
                    truth++;
                }
            }
            delete mapCount[txId];
            if (truth >= (arbitratorNum+1)/2) {
                DataInfoPublished storage data = mapPublishedData[txItem.publishId];
                payToSeller(txItem, data);
            }else {
                // reward arbitrators.
            }
            address[] memory users = new address[](1);
            users[0] = txItem.buyer;
            emit Arbitrate(seqNo, txId, users);
        }
    }

    function closeTransaction(TransactionItem storage txItem, string seqNo, uint256 txId) internal {
        txItem.state = TransactionState.Closed;

        address[] memory users = new address[](1);
        users[0] = address(0x00);
        emit TransactionClose(seqNo, txId, users);
    }

    function payToSeller(TransactionItem storage txItem, DataInfoPublished storage data) internal {
        if (txItem.buyerDeposit >= data.price) {
            txItem.buyerDeposit -= data.price;

            if (!token.transfer(data.seller, data.price)) {
                txItem.buyerDeposit += data.price;
                require(false, "Failed to pay to seller");
            }
        } else {
            require(false, "Low deposit value");
        }
    }

    function setVerifierDepositToken(uint256 deposit) external {
        require(owner == msg.sender, "The value only can be set by owner");

        verifierDepositToken = deposit;
    }

    function setVerifierNum(uint8 num) external {
        require(owner == msg.sender, "The value only can be set by owner");

        verifierNum = num;
    }

    function creditsToVerifier(string seqNo, uint256 txId, address to, uint8 credit) external {
        //validate
        require(to != 0x00, "Verifier address is zero");
        require(credit >= creditLow && credit <= creditHigh, "Valid credit scope is 0 <= credit <= 5");

        TransactionItem storage txItem = mapTransaction[txId];
        require(txItem.used, "Transaction does not exist");

        Verifier storage verifier = getVerifier(to);
        require(verifier.addr != 0x00, "Verifier does not exist");

        DataInfoPublished storage data = mapPublishedData[txItem.publishId];
        require(data.used, "Publish data does not exist");
        require(data.supportVerify, "The transaction do not support verification");

        bool valid;
        uint256 index;
        (valid, index) = verifierValid(verifier, txItem.verifiers);
        require(valid, "Invalid verifier");
        require(!txItem.creditGived[index], "The verifier's credit in this transaction has been submited");

        verifier.credits = (uint8)((verifier.credits * verifier.creditTimes + credit)/(++verifier.creditTimes));
        txItem.creditGived[index] = true;

        //disable verifier and forfeiture deposit while credit <= creditThreshold
        if (verifier.credits <= creditThreshold) {
            verifier.enable = false;
            verifier.deposit = 0;
            validVerifierCount--;
            require(validVerifierCount >= 1, "Invalid verifier count");

            address[] memory users = new address[](1);
            users[0] = address(0x00);
            emit VerifierDisable(seqNo, verifier.addr, users);
        }
    }

    function getVerifier(address v) internal view returns (Verifier storage){
        for (uint256 i = 0; i < verifiers.length; i++) {
            if (verifiers[i].addr == v) {
                return verifiers[i];
            }
        }

        return verifiers[0];
    }
}