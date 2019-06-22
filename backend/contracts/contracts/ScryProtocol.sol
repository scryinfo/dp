pragma solidity ^0.4.24;

import "./lib/common.sol";
import "./lib/transaction.sol";
import "./lib/verification.sol";
import "./ScryToken.sol";

contract ScryProtocol {
    common.PublishedData private publishedData;
    common.TransactionData private txData;

    common.VoteData private voteData;
    common.ArbitratorData private arbitratorData;

    common.Verifiers private verifiers;
    common.Configuration private conf;

    address owner         = 0x0;
    ERC20   token;

    constructor (address _token) public {
        require(_token != 0x0);

        owner = msg.sender;
        token = ERC20(_token);

        //the first element used for empty usage
        verifiers.list[verifiers.list.length++] = common.Verifier(0x00, 0, 0, 0, false);

        conf = common.Configuration(2, 10000, 300, 1, 0, 500, 0, 5, 2, 0, 32); // simple arbitrate params for test
    }

    function registerAsVerifier(string seqNo) external {
        verification.register(verifiers, conf, seqNo, token);
    }

    function vote(string seqNo, uint txId, bool judge, string comments) external {
        verification.vote(verifiers, txData, voteData, conf, seqNo, txId, judge, comments, token);
    }

    function creditsToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) external {
        verification.creditsToVerifier(verifiers, publishedData, txData, conf, seqNo, txId, verifierIndex, credit);
    }

    function arbitrate(string seqNo, uint txId, bool judge) external {
        verification.arbitrate(txData, arbitratorData, conf, txId, judge, token);
        if (verification.arbitrateFinished(arbitratorData, conf, txId)) {
            transaction.arbitrateResult(publishedData, txData, arbitratorData, conf, seqNo, txId, token);
        }
    }

    function publishDataInfo(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller,
        bytes32[] proofDataIds, string descDataId, bool supportVerify) public {
        transaction.publishDataInfo(
            publishedData,
            seqNo,
            publishId,
            price,
            metaDataIdEncSeller,
            proofDataIds,
            descDataId,
            supportVerify
        );
    }

    function createTransaction(string seqNo, string publishId, bool startVerify) external {
        //get verifiers if verification needed
        address[] memory verifiersChosen;
        address[] memory arbitratorsChosen;
        if ( transaction.needVerification(publishedData, publishId, startVerify) ) {
            verifiersChosen = new address[](conf.verifierNum);
            verifiersChosen = verification.chooseVerifiers(verifiers, conf);

            arbitratorsChosen = new address[](conf.arbitratorNum);
            arbitratorsChosen = verification.chooseArbitrators(verifiers, conf, verifiersChosen);
        }

        //create tx
        transaction.createTransaction(
            publishedData,
            txData,
            conf,
            verifiersChosen,
            arbitratorsChosen,
            seqNo,
            publishId,
            startVerify,
            token
        );
    }

    function buyData(string seqNo, uint256 txId) external {
        transaction.buy(
            publishedData,
            txData,
            seqNo,
            txId
        );
    }

    function cancelTransaction(string seqNo, uint256 txId) external {
        transaction.cancelTransaction(
            txData,
            seqNo,
            txId,
            token
        );
    }

    function submitMetaDataIdEncWithBuyer(string seqNo, uint256 txId, bytes encryptedMetaDataId) external {
        transaction.submitMetaDataIdEncByBuyer(
            txData,
            seqNo,
            txId,
            encryptedMetaDataId
        );
    }

    function confirmDataTruth(string seqNo, uint256 txId, bool truth) external {
        transaction.confirmDataTruth(
            publishedData,
            txData,
            conf,
            seqNo,
            txId,
            truth,
            token
        );
    }
}
