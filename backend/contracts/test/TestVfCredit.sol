pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/lib/verification.sol";
import "../contracts/lib/common.sol";

contract TestVfCredit {
    common.DataSet ds;

    string publishId = "pid";
    bytes metaDataIdEncSeller = "";
    address seller = 0x01;
    address verifier1 = 0x2201;
    address verifier2 = 0x2202;
    address arbitrator1 = 0x0101;
    uint256 txId = 20;

    function testCreditsToVerifier() public {
        prepareParamsForCredit();

        verification.creditsToVerifier(ds, "seqNo", txId, 0, 3);
        verification.creditsToVerifier(ds, "seqNo", txId, 1, 1);

        checkParamsForCredit();

        bool r = address(this).call(abi.encodePacked(this.reCreditToVerifier.selector));
        Assert.isFalse(r, "credit twice should not be allowed");
    }

    function reCreditToVerifier() public {
        verification.creditsToVerifier(ds, "seqNo", txId, 0, 5);
    }

    function prepareParamsForCredit() public {
        makePublishData();
        makeTxData();
        makeVerifierData();

        ds.conf = common.Configuration(2, 10000, 300,   1, 0, 500,   0, 5, 2,   0, 32);
    }

    function checkParamsForCredit() public {
        common.Verifier memory v1 = verification.getVerifier(ds.verifiers, verifier1);
        common.Verifier memory v2 = verification.getVerifier(ds.verifiers, verifier2);
        
        Assert.equal(v1.deposit, 13000, "verifier1 deposit unexpected");
        Assert.equal(uint(v1.credits), uint(3), "verifier1 credits unexpected");
        Assert.equal(v1.creditTimes, 11, "verifier1 creditTimes unexpected");
        Assert.equal(v1.enable, true, "verifier1 enable unexpected");

        Assert.equal(v2.deposit, 0, "verifier2 deposit unexpected");
        Assert.equal(uint(v2.credits), uint(2), "verifier2 credits unexpected");
        Assert.equal(v2.creditTimes, 2, "verifier2 creditTimes unexpected");
        Assert.equal(v2.enable, false, "verifier2 enable unexpected");

        Assert.equal(uint(ds.verifiers.validVerifierCount), uint(3), "validVerifierCount unexpected");

        for (uint8 i = 0; i < ds.conf.verifierNum; i++) {
            Assert.equal(ds.txData.map[txId].creditGiven[i], true, "txItem.creditGiven unexpected");
        }
    }

    function makePublishData() internal {
        bytes32[] memory proofDataIds = new bytes32[](2);
        proofDataIds[0] = "";
        proofDataIds[1] = "";
        string memory desc = "";

        ds.pubData.map[publishId] = common.DataInfoPublished(
            1, // price
            metaDataIdEncSeller,
            proofDataIds,
            proofDataIds.length,
            desc,
            seller,
            true, // supportVerify
            true
        );
    }

    function makeTxData() internal {
        address buyer = 0x02;
        address[] memory verifiers = new address[](2);
        verifiers[0] = verifier1;
        verifiers[1] = verifier2;
        bool[] memory creditGiven = new bool[](2);
        creditGiven[0] = false;
        creditGiven[1] = false;
        address[] memory arbitrators = new address[](1);
        arbitrators[0] = arbitrator1;
        bytes memory metaDataIdEncBuyer = "";
        bytes[] memory metaDataIdEncArbitrators = new bytes[](1);

        ds.txData.map[txId] = common.TransactionItem(
            common.TransactionState.Buying,
            buyer,
            seller,
            verifiers,
            creditGiven,
            arbitrators,
            publishId,
            metaDataIdEncBuyer,
            metaDataIdEncSeller,
            metaDataIdEncArbitrators,
            3000, // buyer deposit
            300, // verifier bonus
            500, // arbitrator bonus
            true, // need verify
            true // used
        );
    }

    function makeVerifierData() internal {
        ds.verifiers.list.push(common.Verifier(0x00, 0, 0, 0, false));
        ds.verifiers.list.push(common.Verifier(verifier1, 10000, 3, 10, true));
        ds.verifiers.list.push(common.Verifier(verifier2, 10000, 3, 1, true));
        ds.verifiers.list.push(common.Verifier(0x0001, 10000, 0, 0, true));
        ds.verifiers.list.push(common.Verifier(arbitrator1, 10000, 5, 10, true));
        ds.verifiers.validVerifierCount = 4;
    }
}
