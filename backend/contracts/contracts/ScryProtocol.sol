pragma solidity ^0.4.24;

import {common} from "./lib/common.sol";
import {transaction} from "./lib/transaction.sol";
import "./lib/verification.sol";

import "./ScryToken.sol";

contract ScryProtocol {
    using verification for *;

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

        // todo: conf init
    }

    function registerAsVerifier(string seqNo) external {

    }

    function vote(string seqNo, uint txId, bool judge, string comments) external {

    }

    function creditsToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) external {

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
        bool needVerify = transaction.needVerification(publishedData, publishId, startVerify);
        if (needVerify) {
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

    }

    function cancelTransaction(string seqNo, uint256 txId) external {

    }

    function submitMetaDataIdEncWithBuyer(string seqNo, uint256 txId, bytes encryptedMetaDataId) external {
    }
}
