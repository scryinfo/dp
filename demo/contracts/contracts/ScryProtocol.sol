pragma solidity ^0.4.24;

contract ScryProtocol {
    enum TransactionState {Begin, Created, Voted, Purchasing, ReadyForDownload, Payed, Arbitrating, Closed}

    struct DataInfoPublished {
        bytes metaDataIdEncSeller;
        bytes32[] proofDataIds;
        uint numberOfProof;
        string despDataId;
        address seller;
        bool supportVerify;
        bool used;
    } 

    struct TransactionItem {
        TransactionState state;
        address buyer;
        address seller;
        string publishId;
        bytes meteDataIdEncBuyer;
        bytes metaDataIdEncSeller;
        uint buyerDeposit;
        bool used;
    }

    mapping (string => DataInfoPublished) mapPublishedData;
    mapping (uint => TransactionItem) mapTransaction;

    event Publish(string publishId, string despDataId, bool boardcast, address[] users);
    event TransactionCreate(uint256 transactionId, string publishId, bytes32 chosenProofIds, bool supportVerify, bool boardcast, address[] users);
    event Purchase(uint256 transactionId, string publishId, bytes metaDataIdEncSeller, bool boardcast, address[] users);
    event ReadyForDownload(uint256 transactionId, bytes metaDataIdEncBuyer, bool boardcast, address[] users);
    event Close(uint256 transactionId, bool boardcast, address[] users);

    uint256 transactionSeq = 0;
    uint256 encryptedIdLen = 32;

    constructor () public {
    }

    function publishDataInfo(string publishId, bytes metaDataIdEncSeller, bytes32[] proofDataIds, string despDataId, bool supportVerify) public {
        mapPublishedData[publishId] = DataInfoPublished(metaDataIdEncSeller, proofDataIds, proofDataIds.length, despDataId, msg.sender, supportVerify, true);
        
        address[] memory users = new address[](1);
        users[0] = address(0x00);
        emit Publish(publishId, despDataId, true, users);
    }

    function isPublishedDataExisted(string publishId) internal view returns (bool) {
        DataInfoPublished memory data = mapPublishedData[publishId];
        return data.used;
    }

    function prepareToBuy(string publishId) external {
        //published data info
        DataInfoPublished memory data = mapPublishedData[publishId];
        require(data.used == true, "Can not get published data by publish id");

        //create transaction
        uint txId = getTransactionId();
        bytes memory metaDataIdEncBuyer = new bytes(encryptedIdLen);
        metaDataIdEncBuyer = "";
        mapTransaction[txId] = TransactionItem(TransactionState.Created, msg.sender, data.seller, publishId, metaDataIdEncBuyer, data.metaDataIdEncSeller, 0, true);

        //do not support verification
        if (!data.supportVerify) {
            //choose proof data randomly for buyer 
            uint index = getRandomNumber(data.numberOfProof);
            require(index < data.numberOfProof, "Data index need to be valid");

            bytes32 proofId = data.proofDataIds[index];

            address[] memory users = new address[](1);
            users[0] = msg.sender;

            //TransactionCreat event
            emit TransactionCreate(txId, publishId, proofId, data.supportVerify, false, users);
        }
    }

    function getTransactionId() internal returns(uint) {
        return transactionSeq++;
    }

    function getRandomNumber(uint mod) internal view returns (uint) {
        return uint(keccak256(now, msg.sender)) % mod;
    }

    function buyData(uint256 txId) external {
        //validate
        TransactionItem memory txItem = mapTransaction[txId];
        require(txItem.used == true, "Can not get transaction by transaction id");
        require(txItem.buyer == msg.sender, "Invalid buyer");
        
        //set transaction state
        txItem.state = TransactionState.Purchasing;

        //event
        address[] memory users = new address[](1);
        users[0] = msg.sender;
        emit Purchase(txId, txItem.publishId, txItem.metaDataIdEncSeller, false, users);
    }

    function submitMetaDataIdEncWithBuyer(uint256 txId, bytes encryptedMetaDataId) external {
        //validate
        TransactionItem memory txItem = mapTransaction[txId];
        require(txItem.used == true, "Can not get transaction by transaction id");
        require(txItem.seller == msg.sender, "Invalid seller");

        txItem.meteDataIdEncBuyer = encryptedMetaDataId;
        txItem.state = TransactionState.ReadyForDownload;

        //ReadyForDownload event
        address[] memory users = new address[](1);
        users[0] = txItem.buyer;
        emit ReadyForDownload(txId, txItem.metaDataIdEncSeller, false, users);   
    }

    function confirmDataTruth(uint256 txId, bool truth) external {
        //validate
        TransactionItem memory txItem = mapTransaction[txId];
        require(txItem.buyer != 0, "Can not get transaction by transaction id");
        require(txItem.buyer == msg.sender, "Invalid buyer");

        DataInfoPublished memory data = mapPublishedData[txItem.publishId];
        require(txItem.used == true, "Can not get data by txItem.publishId");
        

        if (!data.supportVerify) {
            if(truth) {

            } else {

            }
            
            txItem.state = TransactionState.Closed;

            //Close event
            address[] memory users = new address[](1);
            users[0] = address(0x00);
            emit Close(txId, true, users);            
        }
    }

}