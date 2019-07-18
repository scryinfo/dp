pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/lib/verification.sol";
import "../contracts/lib/common.sol";
import "../contracts/ScryToken.sol";

contract TestVfVote is ScryToken {
    common.DataSet ds;

    ERC20 token = this;

    string publishId = "testPidVote";
    bytes metaDataIdEncSeller = "";
    address seller = 0x01;
    address verifier1 = msg.sender;
    uint256 txId = 20;

    function testVote() public {
        prepareParamsForVote();

        bool judgeF = true;
        string memory commentsF = "comment first time. ";

        verification.vote(ds, "seqNo", txId, judgeF, commentsF, token);

        checkParamsForVote(judgeF, commentsF);

        bool r = address(this).call(abi.encodePacked(this.reVote.selector));
        Assert.isFalse(r, "vote twice should not be allowed");
    }

    function reVote() public {
        verification.vote(ds, "seqNo", txId, false, "comment second time. ", token);
    }

    function prepareParamsForVote() internal {
        transfer(address(this), 300);
        require(balanceOf(address(this)) == 300, "balance of test contract is wrong. ");

        makePublishData();
        makeTxData();
        makeVerifierData();

        ds.conf = common.Configuration(2, 10000, 300,   1, 0, 500,   0, 5, 2,   0, 32);
    }

    function checkParamsForVote(bool judge, string comments) internal view {
        require(balanceOf(verifier1) == 1000000000, "token balance of msg.sender is wrong");
        require(balanceOf(address(this)) == 0, "token balance of test contract is wrong. ");

        require(ds.voteData.map[txId][verifier1].judge == judge, "vote result - judge is wrong. ");
        require(stringCompare(ds.voteData.map[txId][verifier1].comments, comments), "vote result - comments is wrong. ");
        require(ds.voteData.map[txId][verifier1].used == true, "vote result - used is wrong. ");

        require(ds.txData.map[txId].state == common.TransactionState.Voted, "tx state is wrong. ");
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
        verifiers[1] = 0x2202;
        bool[] memory creditGiven = new bool[](2);
        creditGiven[0] = false;
        creditGiven[1] = false;
        address[] memory arbitrators = new address[](1);
        arbitrators[0] = 0x0101;
        bytes memory metaDataIdEncBuyer = "";
        bytes[] memory metaDataIdEncArbitrators = new bytes[](1);

        ds.txData.map[txId] = common.TransactionItem(
            common.TransactionState.Created,
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
        ds.verifiers.list.push(common.Verifier(0x2202, 10000, 3, 1, true));
        ds.verifiers.list.push(common.Verifier(0x0001, 10000, 0, 0, true));
        ds.verifiers.list.push(common.Verifier(0x0101, 10000, 5, 10, true));
        ds.verifiers.validVerifierCount = 4;
    }

    function stringCompare(string str1, string str2) internal pure returns(bool) {
        require(bytes(str1).length == bytes(str2).length);

        return keccak256(abi.encodePacked(str1)) == keccak256(abi.encodePacked(str2));
    }

}
