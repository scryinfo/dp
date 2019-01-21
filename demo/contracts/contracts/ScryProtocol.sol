pragma solidity ^0.4.24;

import "./ScryToken.sol";

contract ScryProtocol {
    enum TransactionState {Begin, Created, Voted, Buying, ReadyForDownload, Arbitrating, Payed, Closed}
    enum ErrorCode {
        OK,
        DuplicatePublishId,
        InvalidParameter,
        DataNotExist,
        NoEnoughBalance,
        FailedTransferFromCaller,
        TransactionNotExist,
        InvalidSender,
        FailedTransferToSeller,
        InvalidState
    }

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
        string publishId;
        bytes meteDataIdEncBuyer;
        bytes metaDataIdEncSeller;
        uint256 buyerDeposit;
        bool used;
    }

    mapping (string => DataInfoPublished) mapPublishedData;
    mapping (uint => TransactionItem) mapTransaction;

    event DataPublish(string seqNo, string publishId, uint256 price, string despDataId, address[] users);
    event TransactionCreate(string seqNo, uint256 transactionId, string publishId, bytes32 chosenProofIds, bool supportVerify, address[] users);
    event Buy(string seqNo, uint256 transactionId, string publishId, TransactionState state, bytes metaDataIdEncSeller, address[] users);
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
        address[] memory users = new address[](1);
        users[0] = msg.sender;

        //published data info
        DataInfoPublished memory data = mapPublishedData[publishId];
        require(data.used, "Publish data does not exist");
        require((token.balanceOf(msg.sender)) >= data.price, "No enough balance");
        require(token.transferFrom(msg.sender, address(this), data.price), "Failed to transfer token from caller");

        //create transaction
        uint txId = getTransactionId();
        bytes memory metaDataIdEncBuyer = new bytes(encryptedIdLen);
        mapTransaction[txId] = TransactionItem(TransactionState.Created, msg.sender, data.seller, 
                                                publishId, metaDataIdEncBuyer, data.metaDataIdEncSeller, data.price, true);

        if (!data.supportVerify) {
            //choose proof data randomly for buyer
            uint index = getRandomNumber(data.numberOfProof) % data.numberOfProof;
            bytes32 proofId = data.proofDataIds[index];

            //TransactionCreat event
            emit TransactionCreate(seqNo, txId, publishId, proofId, data.supportVerify, users);
        }
    }

    function getTransactionId() internal returns(uint) {
        return transactionSeq++;
    }

    function getRandomNumber(uint mod) internal view returns (uint) {
        return uint(keccak256(now, msg.sender)) % mod;
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
            users = new address[](1);
            users[0] = msg.sender;
        }

        txItem.state = TransactionState.Buying;

        emit Buy(seqNo, txId, txItem.publishId, mapTransaction[txId].state, txItem.metaDataIdEncSeller, users);
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

        if (!data.supportVerify) {
            require(txItem.state == TransactionState.ReadyForDownload);

            if(truth) {
                //pay to seller from contract
                if (txItem.buyerDeposit >= data.price) {
                    txItem.buyerDeposit -= data.price;
                    
                    if (!token.transfer(data.seller, data.price)) {
                        require(false, "Failed to pay to seller");
                    }
                } else {
                    require(false, "Low deposit value");
                }
            }
            
            txItem.state = TransactionState.Closed;

            //TransactionClose event
            users[0] = address(0x00);
            emit TransactionClose(seqNo, txId, users);
        }
    }
}