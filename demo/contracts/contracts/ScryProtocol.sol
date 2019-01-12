pragma solidity ^0.4.24;

import "./ScryToken.sol";

contract ScryProtocol {
    enum TransactionState {Begin, Created, Voted, Buying, ReadyForDownload, Payed, Arbitrating, Closed}
    enum ErrorCode {
        OK,
        DuplicatePublishId,
        InvalidParameter,
        DataNotExist,
        NoEnoughBalance,
        FailedTransferFromCaller
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

    event DataPublish(bool successed, string seqNo, string publishId, uint256 price, string despDataId, address[] users, uint256 errorCode);
    event TransactionCreate(bool successed, string seqNo, uint256 transactionId, string publishId, bytes32 chosenProofIds, bool supportVerify, address[] users, uint256 errorCode);
    event Buy(bool successed, string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, address[] users, uint256 errorCode);
    event ReadyForDownload(bool successed, string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, address[] users, uint256 errorCode);
    event TransactionClose(bool successed, string seqNo, uint256 transactionId, address[] users, uint256 errorCode);

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

        if (publishId == "" || 
            seqNo == ""     ||
            proofDataIds.length == 0 ||
            despDataId == "") {
            emit DataPublish(false, seqNo, publishId, price, despDataId, true, users, ErrorCode.InvalidParameter);
            revert    
        }

        DataInfoPublished data = mapPublishedData[publishId]
        if (data.used) {
            emit DataPublish(false, seqNo, publishId, price, despDataId, true, users, ErrorCode.DuplicatePublishId);
            revert
        }

        data = DataInfoPublished(price, metaDataIdEncSeller, proofDataIds, proofDataIds.length, despDataId, msg.sender, supportVerify, true);
        
        emit DataPublish(true, seqNo, publishId, price, despDataId, true, users, "");
    }

    function isPublishedDataExisted(string publishId) internal view returns (bool) {
        DataInfoPublished memory data = mapPublishedData[publishId];
        return data.used;
    }

    function CreateTransaction(string publishId) external {
        address[] memory users = new address[](1);
        users[0] = msg.sender;
        uint256 errCode = ErrorCode.OK;

        //published data info
        DataInfoPublished memory data = mapPublishedData[publishId];

        if(!data.used) {
            errCode = ErrorCode.DataNotExist;
        } else if(token.balanceOf(msg.sender) < data.price) {
            errCode = ErrorCode.NoEnoughBalance;
        } else if (!token.transferFrom(msg.sender, address(this), data.price)) {
            errCode = ErrorCode.FailedTransferFromCaller;
        }

        if (errCode != ErrorCode.OK) {
            emit TransactionCreate(false, 0, publishId, 0, data.supportVerify, false, users, errCode);
            revert;
        }

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
            emit TransactionCreate(true, txId, publishId, proofId, data.supportVerify, false, users, errCode);
        }
    }

    function getTransactionId() internal returns(uint) {
        return transactionSeq++;
    }

    function getRandomNumber(uint mod) internal view returns (uint) {
        return uint(keccak256(now, msg.sender)) % mod;
    }

    function buyData(uint256 txId) external {
        bool error = false;
        string errMsg;
        address[] memory users = new address[](1);
        users[0] = msg.sender;

        //validate
        TransactionItem memory txItem = mapTransaction[txId];

        if (setError(!txItem.used, errMsg, "Can not get transaction by transaction id") || 
                              setError(txItem.buyer != msg.sender, errMsg, "Invalid buyer")) {
            emit Buy(false, txId, txItem.publishId, txItem.metaDataIdEncSeller, false, users, errMsg);
            revert;
        }

        txItem.state = TransactionState.Buying;
        emit Buy(true, txId, txItem.publishId, txItem.metaDataIdEncSeller, false, users, "");
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
                //pay to seller from contract
                require(token.balanceOf(owner) >= txItem.buyerDeposit && txItem.buyerDeposit >= data.price, "banlance must be enough");

                txItem.buyerDeposit -= data.price;
                if (!token.transfer(data.seller, data.price)) {
                    txItem.buyerDeposit += data.price;
                    require(false, "failed to transfer tokens to seller");
                }
            }
            
            txItem.state = TransactionState.Closed;

            //Close event
            address[] memory users = new address[](1);
            users[0] = address(0x00);
            emit Close(txId, true, users);            
        }
    }

    function getErrorCode(bool condition, uint256 errCode) internal view returns uint256 {
        if (condition) {
            return errCode;
        }

        return ErrorCode.OK;
    }

}