pragma solidity ^0.4.24;

contract ScryProtocol {

    struct DataInfoPublished {
        bytes metaDataIdEncSeller;
        string proofDataId;
        string despDataId;
        address seller;
        bool supportVerify;  
    } 

    struct Transaction {
        enum state {Begin, Created, Voted, Purchasing, ReadyForDownload, Payed, Arbitrating, Closed};
        address buyer;
        address seller;
        string publishId;
        bytes meteDataIdEncBuyer;
        uint buyerDeposit;
    }

    uint public transactionSeq = 0;

    mapping (string => DataInfoPublished) mapPublishedData;

    mapping (int => Transaction) mapTransaction;

    event Published(string publishId, string users, string despDataId);

    constructor () public {
    }

    function PublishDataInfo(string publishId, bytes metaDataIdEncSeller, string proofDataId, string despDataId, bool supportVerify) public {
        mapPublishedData[publishId] = DataInfoPublished(metaDataIdEncSeller, proofDataId, despDataId, msg.sender, supportVerify);
        
        //'*' means everyone can subscribe Published event
        emit Published(publishId, "*", despDataId);
    }

    function IsPublishedDataExisted(string publishId) private returns (bool) {
        DataInfoPublished data = mapPublishedData[publishId];
        if (data == 0 || data.seller == 0) {
            return false;
        }
        
        return true;
    }

    function PrepareToBuy(string publishId) public {
        //published data info
        DataInfoPublished data = mapPublishedData[publishId]

        //maybe publishId not valid: return gas and 
        require(data != 0, "Can not get published data by publish id.")

        //create transaction
        mapTransaction[GetTransactionId()] = Transaction(Created, msg.sender, data.seller, publishId, 0, 0)

        //choose proof data randomly for buyer


        //
        
    }

    function GetTransactionId() private returns(uint) {
        return transactionSeq++;
    }

    function BuyData(string txId) {
        
    }

    function SubmitMetaDataIdEncWithBuyer(string txId, bytes encryptedMetaDataId) {
        
    }

    function ConfirmDataTruth(string txId, bool truth) {
        
    }

}