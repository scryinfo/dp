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
        string publishId;
        bytes meteDataIdEncBuyer;
        bytes metaDataIdEncSeller;
        uint256 buyerDeposit;
        bool used;
    }

    struct Verifier {
        address addr;
        uint256 deposit;
        int256 credits;
        bool enable;
    }

    Verifier[] verifiers;
    uint8 validVerifierCount = 0;
    uint8 verifierNum = 2;
    uint256 verifierDepositToken = 1000;

    struct Vote {
        bool judge;
        string comments;
        bool used;
    }

    mapping (string => DataInfoPublished) mapPublishedData;
    mapping (uint256 => TransactionItem) mapTransaction;
    mapping (uint256 => mapping(address => Vote)) mapVote;

    event RegisterVerifier(string seqNo, address[] users);
    event DataPublish(string seqNo, string publishId, uint256 price, string despDataId, address[] users);
    event TransactionCreate(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bool supportVerify, address[] users);
    event Vote(string seqNo, uint256 transactionId, bool judge, string comments, address[] users);
    event Buy(string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, address[] users);
    event ReadyForDownload(string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, address[] users);
    event TransactionClose(string seqNo, uint256 transactionId, address[] users);

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
    }

    function registerAsVerifier(string seqNo) {
        //deposit
        if (verifierDepositToken > 0) {
            require(token.balanceOf(msg.sender) >= verifierDepositToken, "No enough balance");        
            require(token.transferFrom(msg.sender, address(this), verifierDepositToken), "Failed to transfer token from caller");
        }
        
        verifiers[verifiers.length++] = Verifier(msg.sender, verifierDepositToken, true);
        validVerifierCount++;

        address[] memory users = new address[](1);
        users[0] = msg.sender;
        emit RegisterVerifier(seqNo, users);
    }

    function publishDataInfo(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller,
                                     bytes32[] proofDataIds, string despDataId, bool supportVerify) external {
        address[] memory users = new address[](1);
        users[0] = address(0x00);

        DataInfoPublished memory data = mapPublishedData[publishId];
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
        require(token.transferFrom(msg.sender, address(this), data.price), "Failed to transfer token from caller");

        //create transaction
        uint txId = getTransactionId();
        bytes memory metaDataIdEncBuyer = new bytes(encryptedIdLen);

        address[] memory chosenVerifiers;
        if (data.supportVerify) {
            //choose verifiers randomly
            chosenVerifiers = chooseVerifiers(verifierNum);
            emit TransactionCreate(seqNo, txId, publishId, data.proofDataIds, data.supportVerify, chooseVerifiers);
        }

        address[] memory users = new address[](1);
        users[0] = msg.sender;
        emit TransactionCreate(seqNo, txId, publishId, data.proofDataIds, data.supportVerify, users);

        mapTransaction[txId] = TransactionItem(TransactionState.Created, msg.sender, data.seller, chosenVerifiers,
                                                publishId, metaDataIdEncBuyer, data.metaDataIdEncSeller, data.price, true);
    }

    function chooseVerifiers(uint8 num) internal returns ([]address) {
        require(num < validVerifierCount, "No enough valid verifiers");
        address[] memory chosenVerifiers = new address[](num);

        for (int8 i = 0; i++; i <= num) {
            uint index = getRandomNumber(verifiers.length) % verifiers.length;
            Verifier v = verifiers[index];

            //loop if invalid verifier was chosen until get valid verifier
            address vp = v;
            while (!verifierValid(v, chosenVerifiers)) {
                v = verifier[(++index) % verifier.length].addr;
                require(v != vp, "Disordered verifiers")
            }

            chosenVerifiers[i] = v;
        }

        return chosenVerifiers;
    }

    function verifierValid(Verifier v, address[] arrayVerifier) internal view returns (bool) {
        return (v.enable && !verifierExist(v.addr, arrayVerifier));
    }

    function verifierExist(address v, address[] arrayVerifier) internal view returns (bool) {
        for (uint8 i = 0; i < arrayVerifier.length; i++) {
            if (arrayVerifier[i] == v) {
                return true;
            }
        }

        return false;
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
        require(verifiers[msg.sender].enable, "Invalid verifier")

        mapVote[txId][msg.sender] = Vote(judgement, comments);

        txItem.state = TransactionState.Voted;

        address[] memory users;
        users[0] = txItem.buyer;
        emit Vote(seqNo, txId, judge, comments, users);
    }

    function buyData(string seqNo, uint256 txId) external {
        address[] memory users;

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

        users = new address[](1);
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
        address[] memory users = new address[](1);

        //validate
        TransactionItem storage txItem = mapTransaction[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.buyer == msg.sender, "Invalid buyer");

        DataInfoPublished memory data = mapPublishedData[txItem.publishId];
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
                
            }
        }
    }

    function closeTransaction(TransactionItem txItem, string seqNo, uint256 txId) internal {
        txItem.state = TransactionState.Closed;

        users[0] = address(0x00);
        emit TransactionClose(seqNo, txId, users);
    }

    function payToSeller(TransactionItem txItem, DataInfoPublished data) internal {
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
        require(owner == msg.sender, "The value only can be set by owner")
        
        verifierDepositToken = deposit;
    }

    function setVerifierNum(uint8 num) {
        require(owner == msg.sender, "The value only can be set by owner")
        
        verifierNum = num;
    }
}