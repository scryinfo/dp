pragma solidity ^0.4.24;

import "./common.sol";
import "../ScryToken.sol";

library transaction {
    event DataPublish(string seqNo, string publishId, uint256 price, string despDataId, bool supportVerify, address[] users);
    event TransactionCreate(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bool needVerify, uint8 state, address[] users);
    event Buy(string seqNo, uint256 transactionId, string publishId, bytes metaDataIdEncSeller, uint8 state, uint8 index, address[] users);
    event TransactionClose(string seqNo, uint256 transactionId, uint8 state, uint8 index, address[] users);
    event VerifiersChosen(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, uint8 state, address[] users);
    event ReadyForDownload(string seqNo, uint256 transactionId, bytes metaDataIdEncBuyer, uint8 state, uint8 index, address[] users);
    event ArbitrationBegin(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bytes metaDataIdEncArbitrator, address[] users);
    event ArbitrationResult(string seqNo, uint256 transactionId, bool judge, uint8 identify, address[] users);

    function publishDataInfo(
        common.DataSet storage ds,
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

        common.DataInfoPublished storage data = ds.pubData.map[publishId];
        require(!data.used, "Duplicate publish id");

        ds.pubData.map[publishId] = common.DataInfoPublished(
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
        common.DataSet storage ds,
        string publishId,
        bool startVerify
    ) external view returns (bool) {
        common.DataInfoPublished storage pubItem = ds.pubData.map[publishId];

        return needVerification(pubItem, startVerify);
    }

    function needVerification(
        common.DataInfoPublished storage pubItem,
        bool startVerify
    ) internal view returns (bool) {
        require(pubItem.used, "Publish data does not exist");
        return pubItem.supportVerify && startVerify;
    }

    function createTransaction(
        common.DataSet storage ds,
        address[] verifiers,
        address[] arbitrators,
        string seqNo,
        string publishId,
        bool startVerify,
        ERC20 token
    ) external {
        common.DataInfoPublished storage data = ds.pubData.map[publishId];
        require(data.used, "Publish data does not exist");

        bool needVerify;
        uint256 fee;
        (needVerify, fee) = prepareToCreateTx(ds, publishId, verifiers.length, startVerify);

        buyerDeposit(token, fee);

        if (needVerify) {
            createTxWithVerify(ds, verifiers, arbitrators, seqNo, publishId, fee);
        } else {
            createTxWithoutVerify(ds, seqNo, publishId, fee);
        }

    }

    function prepareToCreateTx(
        common.DataSet storage ds,
        string publishId,
        uint256 verifiersLength,
        bool startVerify
    ) internal view returns (bool, uint256) {
        common.DataInfoPublished storage pubItem = ds.pubData.map[publishId];
        bool needVerify = needVerification(pubItem, startVerify);
        if (needVerify) {
            require(verifiersLength == ds.conf.verifierNum, "Invalid number of verifiers");
        }

        uint256 fee = getFee(ds, publishId, needVerify);

        return (needVerify, fee);
    }

    function createTxWithVerify(
        common.DataSet storage ds,
        address[] verifiers,
        address[] arbitrators,
        string seqNo,
        string publishId,
        uint256 fee
    ) internal {
        uint256 txId = getTransactionId(ds);
        common.DataInfoPublished storage pubItem = ds.pubData.map[publishId];
        bytes32[] storage proofIds = pubItem.proofDataIds;

        address[] memory users = new address[](1);
        for (uint8 i = 0; i < ds.conf.verifierNum; i++) {
            users[0] = verifiers[i];
            emit VerifiersChosen(seqNo, txId, publishId, proofIds, uint8(common.TransactionState.Created), users);
        }

        bool[] memory creditGiven = new bool[](ds.conf.verifierNum);

        bytes memory metaDataIdEncryptedData = new bytes(ds.conf.encryptedIdLen);
        bytes[] memory metaDataIdsEnc = new bytes[](ds.conf.arbitratorNum);
        ds.txData.map[txId] = common.TransactionItem(
            common.TransactionState.Created,
            msg.sender,
            pubItem.seller,
            verifiers,
            creditGiven,
            arbitrators,
            publishId,
            metaDataIdEncryptedData,
            pubItem.metaDataIdEncSeller,
            metaDataIdsEnc,
            fee,
            ds.conf.verifierBonus,
            ds.conf.arbitratorBonus,
            true,
            true
        );

        users[0] = msg.sender;
        emit TransactionCreate(seqNo, txId, publishId, proofIds, true, uint8(common.TransactionState.Created), users);
    }

    function createTxWithoutVerify(
        common.DataSet storage ds,
        string seqNo,
        string publishId,
        uint256 fee
    ) internal {
        uint256 txId = getTransactionId(ds);
        common.DataInfoPublished storage pubItem = ds.pubData.map[publishId];
        bytes32[] storage proofIds = pubItem.proofDataIds;

        address[] memory fills;
        address[] memory users = new address[](1);
        bool[] memory creditGiven;
        bytes memory metaDataIdEncryptedData = new bytes(ds.conf.encryptedIdLen);
        bytes[] memory magic = new bytes[](ds.conf.arbitratorNum);

        ds.txData.map[txId] = common.TransactionItem(
            common.TransactionState.Created,
            msg.sender,
            pubItem.seller,
            fills,
            creditGiven,
            fills,
            publishId,
            metaDataIdEncryptedData,
            pubItem.metaDataIdEncSeller,
            magic,
            fee,
            0,
            0,
            false,
            true
        );

        users[0] = msg.sender;
        emit TransactionCreate(seqNo, txId, publishId, proofIds, true, uint8(common.TransactionState.Created), users);
    }

    function buyerDeposit(
        ERC20 token,
        uint256 fee
    ) internal {
        require((token.balanceOf(msg.sender)) >= fee, "No enough balance");
        require(token.transferFrom(msg.sender, address(this), fee), "Buyer pay deposit failed");
    }

    function getFee(
        common.DataSet storage ds,
        string publishId,
        bool needVerify
    ) internal view returns (uint256) {
        uint256 fee = ds.pubData.map[publishId].price;
        if (needVerify) {
            fee += ds.conf.verifierBonus * ds.conf.verifierNum + ds.conf.arbitratorBonus * ds.conf.arbitratorNum;
        }

        return fee;
    }


    function getTransactionId(common.DataSet storage ds) internal returns(uint) {
        return ds.conf.transactionSeq++;
    }


    function buy(
        common.DataSet storage ds,
        string seqNo,
        uint256 txId
    ) external {
        common.TransactionItem storage txItem = ds.txData.map[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.buyer == msg.sender, "Invalid buyer");

        common.DataInfoPublished storage data = ds.pubData.map[txItem.publishId];
        require(data.used, "Publish data does not exist");

        //buyer can decide to buy even though no verifier response
        require(txItem.state == common.TransactionState.Created || txItem.state == common.TransactionState.Voted, "Invalid transaction state");
        txItem.state = common.TransactionState.Buying;

        address[] memory users = new address[](1);
        users[0] = txItem.seller;
        emit Buy(seqNo, txId, txItem.publishId, txItem.metaDataIdEncSeller, uint8(txItem.state), 0, users);

        users[0] = msg.sender;
        emit Buy(seqNo, txId, txItem.publishId, txItem.metaDataIdEncSeller, uint8(txItem.state), 1, users);

        for (uint8 i = 0; i < txItem.verifiers.length; i++) {
            users[0] = txItem.verifiers[i];
            emit Buy(seqNo, txId, txItem.publishId, txItem.metaDataIdEncSeller, uint8(txItem.state), 2, users);
        }
    }

    function cancelTransaction(
        common.DataSet storage ds,
        string seqNo,
        uint256 txId,
        ERC20 token
    ) external {
        common.TransactionItem storage txItem = ds.txData.map[txId];
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
        emit TransactionClose(seqNo, txId, uint8(txItem.state), 0, users);

        users[0] = txItem.buyer;
        emit TransactionClose(seqNo, txId, uint8(txItem.state), 1, users);

        for (uint8 i = 0; i < txItem.verifiers.length; i++) {
            users[0] = txItem.verifiers[i];
            emit TransactionClose(seqNo, txId, uint8(txItem.state), 2, users);
        }
    }

    function revertToBuyer(common.TransactionItem storage txItem, ERC20 token) internal {
        uint256 deposit = txItem.buyerDeposit;
        txItem.buyerDeposit = 0;

        if (!token.transfer(txItem.buyer, deposit)) {
            txItem.buyerDeposit = deposit;
            require(false, "Failed to revert buyer his token");
        }
    }

    function reEncryptMetaDataIdBySeller(
        common.DataSet storage ds,
        string seqNo,
        uint256 txId,
        bytes encryptedMetaDataIdWithBuyer,
        bytes encryptedMetaDataIdWithArbitrators
    ) external {
        common.TransactionItem storage txItem = ds.txData.map[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.seller == msg.sender, "Invalid seller");
        require(txItem.state == common.TransactionState.Buying, "Invalid transaction state");

        txItem.meteDataIdEncBuyer = encryptedMetaDataIdWithBuyer;

        deserializeAndSave(txItem, encryptedMetaDataIdWithArbitrators, ds.conf.arbitratorNum);

        txItem.state = common.TransactionState.ReadyForDownload;

        address[] memory users = new address[](1);
        users[0] = txItem.seller;
        emit ReadyForDownload(seqNo, txId, txItem.meteDataIdEncBuyer, uint8(txItem.state), 0, users);

        users[0] = txItem.buyer;
        emit ReadyForDownload(seqNo, txId, txItem.meteDataIdEncBuyer, uint8(txItem.state), 1, users);
    }

    function deserializeAndSave(common.TransactionItem storage txItem, bytes encIds, uint8 num) internal {
        require(encIds.length % num == 0, "Invalid Ids length");

        bytes[] memory ids = new bytes[](num);
        uint256 idLen = encIds.length / num;

        for (uint8 i = 0; i < num; i++) {
            bytes memory id = new bytes(idLen);
            for (uint256 count = 0; count < idLen; count++) {
                id[count] = encIds[i*idLen + count];
            }
            ids[i] = id;
        }

        txItem.metaDataIdEncArbitrators = ids;
    }

    function confirmDataTruth(
        common.DataSet storage ds,
        string seqNo,
        uint256 txId,
        bool truth,
        ERC20 token
    ) external {
        //validate
        common.TransactionItem storage txItem = ds.txData.map[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.buyer == msg.sender, "Invalid buyer");

        common.DataInfoPublished storage data = ds.pubData.map[txItem.publishId];
        require(data.used, "Publish data does not exist");

        require(txItem.state == common.TransactionState.ReadyForDownload, "Invalid transaction state");
        if (!txItem.needVerify) {
            if(truth) {
                payToSeller(txItem, data, token);
            }

            closeTransaction(txItem, seqNo, txId);
        } else {
            if (truth) {
                payToSeller(txItem, data, token);
                revertToBuyer(txItem, token);
                closeTransaction(txItem, seqNo, txId);
            } else {
                address[] memory users = new address[](1);
                for (uint8 i = 0; i < ds.conf.arbitratorNum; i++) {
                    users[0] = txItem.arbitrators[i];
                    emit ArbitrationBegin(seqNo, txId, txItem.publishId, data.proofDataIds, txItem.metaDataIdEncArbitrators[i], users);
                }
            }
        }
    }

    function arbitrateResult(common.DataSet storage ds, string seqNo, uint256 txId, ERC20 token) internal {
        uint8 truth;
        for (uint8 i = 0;i < ds.conf.arbitratorNum;i++) {
            if (ds.arbitratorData.map[txId][i].judge) {
                truth++;
            }
        }

        bool result;
        common.TransactionItem storage txItem = ds.txData.map[txId];
        if (truth >= (ds.conf.arbitratorNum+1)/2) {
            common.DataInfoPublished storage data = ds.pubData.map[txItem.publishId];
            payToSeller(txItem, data, token);
            result = true;
        }

        revertToBuyer(txItem, token);

        address[] memory users = new address[](1);
        users[0] = txItem.seller;
        emit ArbitrationResult(seqNo, txId, result, 0, users);

        users[0] = txItem.buyer;
        emit ArbitrationResult(seqNo, txId, result, 1, users);

        closeTransaction(txItem, seqNo, txId);
    }

    function payToSeller(
        common.TransactionItem storage txItem,
        common.DataInfoPublished storage data,
        ERC20 token
    ) internal {
        if (txItem.buyerDeposit >= data.price) {
            txItem.buyerDeposit -= data.price;

            if (!token.transfer(data.seller, data.price)) {
                txItem.buyerDeposit += data.price;
                require(false, "Failed to pay to seller");
            }
        } else {
            require(false, "No enough deposit for seller");
        }
    }

    function getBuyerAddrInDesignatedTx(common.DataSet storage ds, uint256 txId) internal view returns (address) {
        common.TransactionItem memory txItem = ds.txData.map[txId];
        require(msg.sender == txItem.seller, "Invalid caller");
        require(txItem.state == common.TransactionState.Buying, "Invalid transaction state");

        return txItem.buyer;
    }

    function getArbitratorsAddrsInDesignatedTx(common.DataSet storage ds, uint256 txId) internal view returns (address[]) {
        common.TransactionItem memory txItem = ds.txData.map[txId];
        require(msg.sender == txItem.seller, "Invalid caller");
        require(txItem.state == common.TransactionState.Buying, "Invalid transaction state");

        return txItem.arbitrators;
    }
}
