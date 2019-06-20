pragma solidity ^0.4.24;

import "./common.sol";
import "../ScryToken.sol";

library transaction {
    event DataPublish(string seqNo, string publishId, uint256 price, string despDataId, bool supportVerify, address[] users);
    event TransactionCreate(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bool needVerify, common.TransactionState state, address[] users);
    event Buy(string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, common.TransactionState state, address buyer, uint8 index, address[] users);
    event TransactionClose(string seqNo, uint256 transactionId, common.TransactionState state, uint8 index, address[] users);
    event VerifiersChosen(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, common.TransactionState state, address[] users);
    event ReadyForDownload(string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, common.TransactionState state, uint8 index, address[] users);

    function publishDataInfo(
        common.PublishedData storage self,
        string seqNo,
        string publishId,
        uint256 price,
        bytes metaDataIdEncSeller,
        bytes32[] proofDataIds,
        string descDataId,
        bool supportVerify
    ) external {
        address[] memory users = new address[](1);
        users[0] = address(0x00);

        common.DataInfoPublished storage data = self.map[publishId];
        require(!data.used, "Duplicate publish id");

        self.map[publishId] = common.DataInfoPublished(
            price,
            metaDataIdEncSeller,
            proofDataIds,
            proofDataIds.length,
            descDataId,
            msg.sender,
            supportVerify,
            true
        );

        emit DataPublish(seqNo, publishId, price, descDataId, supportVerify, users);
    }

    function needVerification(
        common.PublishedData storage pubData,
        string publishId,
        bool startVerify
    ) external view returns (bool) {
        common.DataInfoPublished storage data = pubData.map[publishId];
        require(data.used, "Publish data does not exist");

        return needVerification(data, startVerify);
    }

    function needVerification(
        common.DataInfoPublished storage pubItem,
        bool startVerify
    ) public view returns (bool) {
        return pubItem.supportVerify && startVerify;
    }

    function createTransaction(
        common.PublishedData storage pubData,
        common.TransactionData storage txData,
        common.Configuration storage conf,
        address[] verifiers,
        string seqNo,
        string publishId,
        bool startVerify,
        ERC20 token
    ) public {
        common.DataInfoPublished storage data = pubData.map[publishId];
        require(data.used, "Publish data does not exist");
        require(verifiers.length == conf.verifierNum, "Invalid number of verifiers");

        uint256 fee = data.price;
        bool needVerify = needVerification(data, startVerify);
        if (needVerify) {
            fee += conf.verifierBonus * conf.verifierNum;
        }

        require((token.balanceOf(msg.sender)) >= fee, "No enough balance");
        require(token.transferFrom(msg.sender, address(this), fee), "Failed to transfer token from caller");

        uint txId = getTransactionId(conf);
        bytes memory metaDataIdEncryptedData = new bytes(conf.encryptedIdLen);
        bool[] memory creditGiven;
        address[] memory users = new address[](1);
        if (needVerify) {
            for (uint8 i = 0; i < conf.verifierNum; i++) {
                users[0] = verifiers[i];
                emit VerifiersChosen(seqNo, txId, publishId, data.proofDataIds, common.TransactionState.Created, users);
            }

            creditGiven = new bool[](conf.verifierNum);
        }

        txData.map[txId] = common.TransactionItem(
            common.TransactionState.Created,
            msg.sender,
            data.seller,
            verifiers,
            creditGiven,
            publishId,
            metaDataIdEncryptedData,
            data.metaDataIdEncSeller,
            fee,
            conf.verifierBonus,
            needVerify,
            true
        );

        users[0] = msg.sender;
        emit TransactionCreate(seqNo, txId, publishId, data.proofDataIds, needVerify, common.TransactionState.Created, users);
    }

    function getTransactionId(common.Configuration storage conf) internal returns(uint) {
        return conf.transactionSeq++;
    }


    function buy(
        common.PublishedData storage pubData,
        common.TransactionData storage txData,
        string seqNo,
        uint256 txId
    ) external {
        common.TransactionItem storage txItem = txData.map[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.buyer == msg.sender, "Invalid buyer");

        common.DataInfoPublished storage data = pubData.map[txItem.publishId];
        require(data.used, "Publish data does not exist");

        //buyer can decide to buy even though no verifier response
        require(txItem.state == common.TransactionState.Created || txItem.state == common.TransactionState.Voted, "Invalid transaction state");
        txItem.state = common.TransactionState.Buying;

        address[] memory users = new address[](1);
        users[0] = txItem.seller;
        emit Buy(seqNo, txId, txItem.publishId, txItem.metaDataIdEncSeller, txItem.state, txItem.buyer, 0, users);

        users[0] = msg.sender;
        emit Buy(seqNo, txId, txItem.publishId, txItem.metaDataIdEncSeller, txItem.state, txItem.buyer, 1, users);
    }

    function cancelTransaction(
        common.TransactionData storage txData,
        string seqNo,
        uint256 txId,
        ERC20 token
) external {
        common.TransactionItem storage txItem = txData.map[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.buyer == msg.sender, "Invalid cancel operator");

        require(txItem.state == common.TransactionState.Created || txItem.state == common.TransactionState.Voted ||
        txItem.state == common.TransactionState.Buying, "Invalid transaction state");

        revertToBuyer(txItem, token);
        closeTransaction(txItem, seqNo, txId);
    }

    function closeTransaction(
        common.TransactionItem storage txItem,
        string seqNo,
        uint256 txId
    ) internal {
        txItem.state = common.TransactionState.Closed;

        address[] memory users = new address[](1);
        users[0] = txItem.seller;
        emit TransactionClose(seqNo, txId, txItem.state, 0, users);

        users[0] = txItem.buyer;
        emit TransactionClose(seqNo, txId, txItem.state, 1, users);
    }

    function revertToBuyer(common.TransactionItem storage txItem, ERC20 token) internal {
        uint256 deposit = txItem.buyerDeposit;
        txItem.buyerDeposit = 0;

        if (!token.transfer(txItem.buyer, deposit)) {
            txItem.buyerDeposit = deposit;
            require(false, "Failed to revert to buyer his token");
        }
    }

    function submitMetaDataIdEncByBuyer(
        common.TransactionData storage txData,
        string seqNo,
        uint256 txId,
        bytes encryptedMetaDataId
    ) external {
        common.TransactionItem storage txItem = txData.map[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.seller == msg.sender, "Invalid seller");
        require(txItem.state == common.TransactionState.Buying, "Invalid transaction state");

        txItem.meteDataIdEncBuyer = encryptedMetaDataId;
        txItem.state = common.TransactionState.ReadyForDownload;

        address[] memory users = new address[](1);
        users[0] = txItem.seller;
        emit ReadyForDownload(seqNo, txId, txItem.meteDataIdEncBuyer, txItem.state, 0, users);

        users[0] = txItem.buyer;
        emit ReadyForDownload(seqNo, txId, txItem.meteDataIdEncBuyer, txItem.state, 1, users);
    }
}
