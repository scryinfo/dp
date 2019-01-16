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
        FailedTransferFromCaller,
        TransactionNotExist,
        InvalidSender,
        FailedTransferToSeller
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

    event DataPublish(bool successed, string seqNo, string publishId, uint256 price, string despDataId, address[] users, ErrorCode errorCode);
    event TransactionCreate(bool successed, string seqNo, uint256 transactionId, string publishId, bytes32 chosenProofIds, bool supportVerify, address[] users, ErrorCode errorCode);
    event Buy(bool successed, string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, address[] users, ErrorCode errorCode);
    event ReadyForDownload(bool successed, string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, address[] users, ErrorCode errorCode);
    event TransactionClose(bool successed, string seqNo, uint256 transactionId, address[] users, ErrorCode errorCode);

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
        if (data.used) {
            emit DataPublish(false, seqNo, publishId, price, despDataId, users, ErrorCode.DuplicatePublishId);
            revert();
        }

        data = DataInfoPublished(price, metaDataIdEncSeller, proofDataIds, proofDataIds.length, despDataId, msg.sender, supportVerify, true);
        
        emit DataPublish(true, seqNo, publishId, price, despDataId, users, ErrorCode.OK);
    }

    function isPublishedDataExisted(string publishId) internal view returns (bool) {
        DataInfoPublished memory data = mapPublishedData[publishId];
        return data.used;
    }

    function CreateTransaction(string seqNo, string publishId) external {
        address[] memory users = new address[](1);
        users[0] = msg.sender;
        ErrorCode errCode = ErrorCode.OK;

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
            emit TransactionCreate(false, seqNo, 0, publishId, 0, data.supportVerify, users, errCode);
            revert();
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
            emit TransactionCreate(true, seqNo, txId, publishId, proofId, data.supportVerify, users, errCode);
        }
    }

    function getTransactionId() internal returns(uint) {
        return transactionSeq++;
    }

    function getRandomNumber(uint mod) internal view returns (uint) {
        return uint(keccak256(now, msg.sender)) % mod;
    }

    function buyData(string seqNo, uint256 txId) external {
        address[] memory users = new address[](1);
        users[0] = msg.sender;
        ErrorCode errCode = ErrorCode.OK;

        //validate
        TransactionItem memory txItem = mapTransaction[txId];

        if (!txItem.used) {
            errCode = ErrorCode.TransactionNotExist;
        } else if (txItem.buyer != msg.sender) {
            errCode = ErrorCode.InvalidSender;
        } 

        if (errCode != ErrorCode.OK) {
            emit Buy(false, seqNo, txId, txItem.publishId, txItem.metaDataIdEncSeller, users, errCode);
            revert();
        }

        txItem.state = TransactionState.Buying;
        users[0] = txItem.seller;
        emit Buy(true, seqNo, txId, txItem.publishId, txItem.metaDataIdEncSeller, users, errCode);
    }

    function submitMetaDataIdEncWithBuyer(string seqNo, uint256 txId, bytes encryptedMetaDataId) external {
        address[] memory users = new address[](1);
        users[0] = msg.sender;
        ErrorCode errCode = ErrorCode.OK;

        //validate
        TransactionItem memory txItem = mapTransaction[txId];
        if (!txItem.used) {
            errCode = ErrorCode.TransactionNotExist;
        } else if (txItem.seller != msg.sender) {
            errCode = ErrorCode.InvalidSender;
        } 

        if (errCode != ErrorCode.OK) {
            emit ReadyForDownload(false, seqNo, txId, txItem.metaDataIdEncSeller, users, errCode);   
            revert();
        }

        txItem.meteDataIdEncBuyer = encryptedMetaDataId;
        txItem.state = TransactionState.ReadyForDownload;

        //ReadyForDownload event
        users[0] = txItem.buyer;
        emit ReadyForDownload(false, seqNo, txId, txItem.metaDataIdEncSeller, users, errCode);   
    }

    function confirmDataTruth(string seqNo, uint256 txId, bool truth) external {
        address[] memory users = new address[](1);
        users[0] = msg.sender;
        ErrorCode errCode = ErrorCode.OK;

        //validate
        TransactionItem memory txItem = mapTransaction[txId];
        if (!txItem.used) {
            errCode = ErrorCode.TransactionNotExist;
        } else if (txItem.buyer != msg.sender) {
            errCode = ErrorCode.InvalidSender;
        } 

        if (errCode != ErrorCode.OK) {
            emit TransactionClose(false, seqNo, txId, users, errCode);
            revert();
        }

        DataInfoPublished memory data = mapPublishedData[txItem.publishId];
        if (!data.used) {
            emit TransactionClose(false, seqNo, txId, users, ErrorCode.DataNotExist);
            revert();
        }

        if (!data.supportVerify) {
            if(truth) {
                //pay to seller from contract
                txItem.buyerDeposit -= data.price;
                if (!token.transfer(data.seller, data.price)) {
                    txItem.buyerDeposit += data.price;
                    emit TransactionClose(false, seqNo, txId, users, ErrorCode.FailedTransferToSeller);
                    revert();
                }
            }
            
            txItem.state = TransactionState.Closed;

            //TransactionClose event
            users[0] = address(0x00);
            emit TransactionClose(true, seqNo, txId, users, ErrorCode.OK);
        }
    }
}