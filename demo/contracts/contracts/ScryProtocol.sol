pragma solidity ^0.4.24;

contract ScryProtocol {

    struct DataInfoPublished {
        bytes metaDataIdEncSeller;
        string proofDataId;
        string despDataId;
        address sellerAddr;
        bool supportVerify;  
    } 

    mapping (string => DataInfoPublished) mapPublishedData;

    event Published(string publishId, string users, string despDataId);

    constructor () public {
    }

    function publishDataInfo(string publishId, bytes metaDataIdEncSeller, string proofDataId, string despDataId, bool supportVerify) {
        mapPublishedData[publishId] = DataInfoPublished(metaDataIdEncSeller, proofDataId, despDataId, msg.sender, supportVerify);
        
        //'*' means everyone can get despDataId
        emit Published(publishId, "*", despDataId);
    }

    function isPublishedDataExisted(string publishId) returns (bool) {
        DataInfoPublished data = mapPublishedData[publishId];
        if (data.sellerAddr == 0) {
            return false;
        }
        
        return true;
    }

    function prepareToBuy(string publishId) {
        
    }

    function buyData(string txId) {
        
    }

    function SubmitMetaDataIdEncWithBuyer(string txId, bytes encryptedMetaDataId) {
        
    }

    function ConfirmDataTruth(string txId, bool truth) {
        
    }

}