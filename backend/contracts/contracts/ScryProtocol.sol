pragma solidity ^0.4.24;

import "./lib/common.sol";
import "./lib/transaction.sol";
import "./lib/verification.sol";

import "./ScryToken.sol";

contract ScryProtocol {
    using verification for common.Verifiers;
    
    common.PublishedData private publishedData;
    common.TransactionData private txData;
    common.VoteData voteData;

    common.Verifiers verifiers;
    common.Configuration conf;

    address owner         = 0x0;
    ERC20   token;

    event RegisterVerifier(string seqNo, address[] users);
    event VerifiersChosen(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, common.TransactionState state, address[] users);
    event Vote(string seqNo, uint256 transactionId, bool judge, string comments, common.TransactionState state, uint8 index, address[] users);
    event VerifierDisable(string seqNo, address verifier, address[] users);


    constructor (address _token) public {
        require(_token != 0x0);

        owner = msg.sender;
        token = ERC20(_token);

        // todo: conf init
    }

    function registerAsVerifier(string seqNo) external {
        bool exist;
        (exist, ) = verifiers.inVerifier(msg.sender);
        require(!exist, "Address already registered");

        //deposit
        if (conf.verifierDepositToken > 0) {
            require(token.balanceOf(msg.sender) >= conf.verifierDepositToken, "No enough balance");
            require(token.transferFrom(msg.sender, address(this), conf.verifierDepositToken), "Pay deposit failed");
        }

        verifiers.register(conf, msg.sender);

        address[] memory users = new address[](1);
        users[0] = msg.sender;
        emit RegisterVerifier(seqNo, users);
    }

    // simulate call chooseVerifiers interface
    function chooseVerifiersInterfaceCallSimulation() internal {
        address[] memory verifiersChosen = new address[](conf.verifierNum);
        verifiersChosen = verifiers.chooseVerifiers(conf);
        emit VerifiersChosen("some params");
    }

    function vote(string seqNo, uint txId, bool judge, string comments) external {
        common.TransactionItem storage txItem = transactionData.map[txId];

        require(txItem.used, "Transaction does not exist");
        require(txItem.state == TransactionState.Created || txItem.state == TransactionState.Voted, "Invalid transaction state");

        bool valid;
        uint8 index;
        (valid, index) = verifiers.validVerifiers(txItem.verifiers, msg.sender);
        require(valid, "Invalid verifier");

        if (!voteData.map[txId][msg.sender].used) {
            // interface2
        }

        // interface3

        address[] memory users = new address[](1);
        users[0] = txItem.buyer;
        emit Vote(seqNo, txId, judge, comments, txItem.state, index+1, users);
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
