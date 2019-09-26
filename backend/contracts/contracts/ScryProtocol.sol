pragma solidity ^0.4.24;

import "./lib/common.sol";
import "./lib/transaction.sol";
import "./lib/verification.sol";
import "./ScryToken.sol";

contract ScryProtocol {
    common.DataSet private dataSet;

    address owner;
    ERC20   token;

    event DataPublish(string seqNo, string publishId, uint256 price, string despDataId, bool supportVerify, address[] users);
    event TransactionCreate(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bool needVerify, uint8 state, address[] users);
    event Buy(string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, uint8 state, uint8 index, address[] users);
    event TransactionClose(string seqNo, uint256 transactionId, uint8 state, uint8 index, address[] users);
    event VerifiersChosen(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, uint8 state, address[] users);
    event ReadyForDownload(string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, uint8 state, uint8 index, address[] users);
    event ArbitrationBegin(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bytes metaDataIdEncArbitrator, address[] users);
    event ArbitrationResult(string seqNo, uint256 transactionId, bool judge, uint8 identify, address[] users);
    event RegisterVerifier(string seqNo, address[] users);
    event Vote(string seqNo, uint256 transactionId, bool judge, string comments, uint8 state, uint8 index, address[] users);
    event VerifierDisable(string seqNo, address verifier, address[] users);

    constructor (address _token) public {
        require(_token != 0x0);

        owner = msg.sender;
        token = ERC20(_token);

        //the first element used for empty usage
        dataSet.verifiers.list[dataSet.verifiers.list.length++] = common.Verifier(0x00, 0, 0, 0, false);

        dataSet.conf = common.Configuration(2, 10000, 300,   1, 0, 500,   0, 5, 2,   0, 32); // simple arbitrate params for test
    }

    function registerAsVerifier(string seqNo) external {
        verification.register(dataSet, seqNo, token);
    }

    function vote(string seqNo, uint txId, bool judge, string comments) external {
        verification.vote(dataSet, seqNo, txId, judge, comments, token);
    }

    function creditsToVerifier(string seqNo, uint256 txId, uint8 verifierIndex, uint8 credit) external {
        verification.creditsToVerifier(dataSet, seqNo, txId, verifierIndex, credit);
    }

    function arbitrate(string seqNo, uint txId, bool judge) external {
        verification.arbitrate(dataSet, txId, judge, token);
        if (verification.arbitrateFinished(dataSet, txId)) {
            transaction.arbitrateResult(dataSet, seqNo, txId, token);
        }
    }

    function publishDataInfo(string seqNo, string publishId, uint256 price, bytes metaDataIdEncSeller,
        bytes32[] proofDataIds, string descDataId, bool supportVerify) public {
        transaction.publishDataInfo(
            dataSet,
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
        if ( transaction.needVerification(dataSet, publishId, startVerify) ) {
            address seller = dataSet.pubData.map[publishId].seller;

            verifiersChosen = new address[](dataSet.conf.verifierNum);
            verifiersChosen = verification.chooseVerifiers(dataSet, seller);

            arbitratorsChosen = new address[](dataSet.conf.arbitratorNum);
            arbitratorsChosen = verification.chooseArbitrators(dataSet, verifiersChosen, seller);
        }

        //create tx
        transaction.createTransaction(
            dataSet,
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
            dataSet,
            seqNo,
            txId
        );
    }

    function cancelTransaction(string seqNo, uint256 txId) external {
        transaction.cancelTransaction(
            dataSet,
            seqNo,
            txId,
            token
        );
    }

    function reEncryptMetaDataIdBySeller(string seqNo, uint256 txId, bytes encryptedMetaDataId, bytes encryptedMetaDataIds) external {
        transaction.reEncryptMetaDataIdBySeller(
            dataSet,
            seqNo,
            txId,
            encryptedMetaDataId,
            encryptedMetaDataIds
        );
    }

    function confirmDataTruth(string seqNo, uint256 txId, bool truth) external {
        transaction.confirmDataTruth(
            dataSet,
            seqNo,
            txId,
            truth,
            token
        );
    }

    function getBuyer(uint256 txId) external view returns (address) {
        return transaction.getBuyer(dataSet, txId);
    }

    function getArbitrators(uint256 txId) external view returns (address[]) {
        return transaction.getArbitrators(dataSet, txId);
    }
}
