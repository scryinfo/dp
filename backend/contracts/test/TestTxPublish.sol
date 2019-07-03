pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/lib/transaction.sol";
import "../contracts/lib/common.sol";

contract TestTxPublish {
    common.DataSet ds;

    function testPublishDataInfo() public {
        delete ds;

        string memory seqNo = "seqNo";
        string memory publishId = "publishId";
        uint256 price = 1000;
        bytes memory metaDataIdEncSeller = "meta data id";
        bytes32[] memory proofDataIds = new bytes32[](2);
        proofDataIds[0] = "proof1";
        proofDataIds[1] = "proof2";
        string memory desc = "data description";
        bool supportVerify = true;

        common.DataInfoPublished storage data = ds.pubData.map[publishId];
        Assert.equal(data.used, false, "data should be initialized correctly");

        transaction.publishDataInfo(ds, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, desc, supportVerify);

        Assert.equal(data.seller, tx.origin, "data should be correct");
        Assert.equal(data.used, true, "data should be correct");
        Assert.equal(data.price, price, "data should be correct");

        bool r = address(this).call(abi.encodePacked(this.publishTwiceWithSamePublishId.selector));
        Assert.isFalse(r, "a same publish id should not publish twice");
    }

    function publishTwiceWithSamePublishId() public {
        string memory seqNo = "";
        string memory publishId = "publishId";
        uint256 price = 1;
        bytes memory metaDataIdEncSeller = "";
        bytes32[] memory proofDataIds = new bytes32[](2);
        proofDataIds[0] = "";
        proofDataIds[1] = "";
        string memory desc = "";
        bool supportVerify = false;

        transaction.publishDataInfo(ds, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, desc, supportVerify);
    }


}
