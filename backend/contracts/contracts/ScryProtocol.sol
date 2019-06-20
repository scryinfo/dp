pragma solidity ^0.4.24;

import "./lib/common.sol";
import "./lib/transaction.sol";
import "./lib/verification.sol";
import "./ScryToken.sol";

contract ScryProtocol {
    using verification for *;
    using common for *;
    using transaction for *;

    common.PublishedData private publishedData;
    common.TransactionData private txData;
    common.VoteData private voteData;

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

        conf = common.Configuration(2, 300, 10000, 0, 5, 2, 0, 32);
    }

    function registerAsVerifier(string seqNo) external {
        verifiers.register(conf, seqNo, token);
    }

    function vote(string seqNo, uint txId, bool judge, string comments) external {
        verifiers.vote(txData, voteData, seqNo, txId, judge, comments, token);
    }

    function creditsToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) external {
        verifiers.creditsToVerifier(publishedData, txData, conf, seqNo, txId, verifierIndex, credit);
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
        if ( transaction.needVerification(publishedData, publishId, startVerify) ) {
            verifiersChosen = new address[](conf.verifierNum);
            verifiersChosen = verification.chooseVerifiers(verifiers, conf);
        }

        //create tx
        transaction.createTransaction(
            publishedData,
            txData,
            conf,
            verifiersChosen,
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
            seqNo,
            txId,
            truth,
            token
        );
    }
}
