pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/lib/transaction.sol";
import "../contracts/lib/common.sol";

contract TestTransaction {
    common.DataSet ds;

    function testPublishDataInfo() {
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
    }
}
