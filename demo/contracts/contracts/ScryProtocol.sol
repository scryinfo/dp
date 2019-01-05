pragma solidity ^0.4.24;

contract ScryProtocol {
    enum TransactionState {Begin, Created, Voted, Purchasing, ReadyForDownload, Payed, Arbitrating, Closed}

    struct DataInfoPublished {
        bytes metaDataIdEncSeller;
        string[] proofDataIds;
        uint numberOfProof;
        string despDataId;
        address seller;
        bool supportVerify;  
    } 

    struct Transaction {
        TransactionState state;
        address buyer;
        address seller;
        string publishId;
        bytes meteDataIdEncBuyer;
        uint buyerDeposit;
    }

    mapping (string => DataInfoPublished) mapPublishedData;
    mapping (uint => Transaction) mapTransaction;

    event Published(string publishId, string users, string despDataId);
    event TransactionCreated(uint transactionId, string publishId, string chosenProofIds, bool supportVerify, address[] users);

    uint public transactionSeq = 0;
    uint public encryptedIdLen = 32;

    constructor () public {
    }

    function publishDataInfo(string publishId, bytes metaDataIdEncSeller, string proofDataIdList, string despDataId, bool supportVerify) public {
        //split string to string array
        string[] memory proofDataIds;
        uint proofNum;

        (proofDataIds, proofNum) = splitString(proofDataIdList);
        mapPublishedData[publishId] = DataInfoPublished(metaDataIdEncSeller, proofDataIds, proofNum, despDataId, msg.sender, supportVerify);
        
        //'*' means everyone can subscribe Published event
        emit Published(publishId, "*", despDataId);
    }

    function isPublishedDataExisted(string publishId) private returns (bool) {
        DataInfoPublished memory data = mapPublishedData[publishId];
        if (data.seller == 0) {
            return false;
        }
        
        return true;
    }

    function prepareToBuy(string publishId) public {
        //published data info
        DataInfoPublished memory data = mapPublishedData[publishId];
        require(data.seller != 0, "Can not get published data by publish id.");

        //create transaction
        uint txId = getTransactionId();
        bytes memory metaDataIdEncBuyer = new bytes(encryptedIdLen);
        metaDataIdEncBuyer = "";
        mapTransaction[txId] = Transaction(TransactionState.Created, msg.sender, data.seller, publishId, metaDataIdEncBuyer, 0);

        //do not support verification
        if (!data.supportVerify) {
            //choose proof data randomly for buyer 
            uint index = getRandomNumber(data.numberOfProof);
            require(index < data.numberOfProof, "Data index need to be valid.");

            string memory proofId = data.proofDataIds[index];
            address[] memory users = new address[](1);
            users[0] = msg.sender;

            //TransactionCreated event
            emit TransactionCreated(txId, publishId, proofId, data.supportVerify, users);
        }
    }

    function getTransactionId() private returns(uint) {
        return transactionSeq++;
    }

    function splitString(string src) private returns (string[], uint) {
        //split string into string array 
        string[] memory result = new string[](1);
        result[0] = src;
        return (result, 1);
    }

    function getRandomNumber(uint mod) private returns (uint) {
        return uint(keccak256(now, msg.sender)) % mod;
    }

    function addressToString(address x) returns (string) {
        bytes memory b = new bytes(20);
        for (uint i = 0; i < 20; i++)
            b[i] = byte(uint8(uint(x) / (2**(8*(19 - i)))));
        return string(b);
    }

    function buyData(string txId) {
        
    }

    function submitMetaDataIdEncWithBuyer(string txId, bytes encryptedMetaDataId) {
        
    }

    function confirmDataTruth(string txId, bool truth) {
        
    }

}