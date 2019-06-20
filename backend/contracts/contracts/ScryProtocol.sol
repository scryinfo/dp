pragma solidity ^0.4.24;

import "./lib/common.sol";
import "./lib/transaction.sol";
import "./lib/verification.sol";

import "./ScryToken.sol";

contract ScryProtocol {
    using verification for *;

    common.PublishedData private publishedData;
    common.TransactionData private txData;
    common.VoteData voteData;

    common.Verifiers verifiers;
    common.Configuration conf;

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
        bytes32[] proofDataIds, string descDataId, bool supportVerify) external {

    }

    function createTransaction(string seqNo, string publishId, bool startVerify) external {

    }

    function buyData(string seqNo, uint256 txId) external {

    }

    function cancelTransaction(string seqNo, uint256 txId) external {

    }

    function submitMetaDataIdEncWithBuyer(string seqNo, uint256 txId, bytes encryptedMetaDataId) external {
    }
}
