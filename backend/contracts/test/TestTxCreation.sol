pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/lib/transaction.sol";
import "../contracts/lib/common.sol";

contract TestTxCreation {
    common.DataSet ds;

    function makePublishData() internal {
        string memory publishId = "pid";
        uint256 price = 1;
        bytes memory metaDataIdEncSeller = "";
        bytes32[] memory proofDataIds = new bytes32[](2);
        proofDataIds[0] = "";
        proofDataIds[1] = "";
        string memory desc = "";
        bool supportVerify = true;
        address seller = 0x01234;

        ds.pubData.map[publishId] = common.DataInfoPublished(
            price,
            metaDataIdEncSeller,
            proofDataIds,
            proofDataIds.length,
            desc,
            seller,
            supportVerify,
            true
        );
    }

    function testCreateTransactionWithVerify() public {
        delete ds;

        makePublishData();

        string memory publishId = "pid";
        address[] memory verifiers = new address[](2);
        verifiers[0] = 0x01;
        verifiers[1] = 0x02;

        address[] memory arbitrators = new address[](2);
        arbitrators[0] = 0x03;
        arbitrators[1] = 0x04;

        string memory seqNo = "seq no";
        uint256 fee = 10000;

        ds.conf.transactionSeq = 10;
        ds.conf.verifierNum = 2;
        ds.conf.verifierBonus = 1000;
        ds.conf.arbitratorNum = 1;
        ds.conf.arbitratorBonus = 1000;
        transaction.createTxWithVerify(
            ds,
            verifiers,
            arbitrators,
            seqNo,
            publishId,
            fee
        );

        Assert.equal(ds.conf.transactionSeq, 11, "transaction sequence number should be correct");

        address seller = 0x01234;
        common.TransactionItem storage txData = ds.txData.map[10];
        Assert.equal(txData.used, true, "transaction data should be not empty");
        Assert.equal(txData.buyer, msg.sender, "buyer should be correct");
        Assert.equal(txData.seller, seller, "seller should be correct");
    }

}
