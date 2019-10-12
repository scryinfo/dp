import {acc_db, dl_db, tx_db} from "./DBoptions";
import {connect} from "./connect";

let utils = {
    stateEnum: ["Begin", "Created", "Voted", "Buying", "ReadyForDownload", "Closed"],
    state: [ // tx state -> func state, means func button can or can't click
        [false, false, false,  true, false, false], // 0 seller: re-encrypt
        [false,  true,  true,  true, false, false], // 1 buyer: cancel
        [false,  true,  true, false, false, false], // 2 buyer: confirmPurchase & verifier: vote
        [false, false, false, false,  true, false], // 3 buyer: decrypt/confirmData & arbitrator: decrypt/arbitrate
        [false, false,  true,  true,  true,  true]  // 4 buyer: grade
    ],
    init: function () {
        connect.addCallbackFunc("onPublish", presetFunc.onPublish);
        connect.addCallbackFunc("onProofFilesExtensions", presetFunc.onProofFilesExtensions);
        connect.addCallbackFunc("onVerifiersChosen", presetFunc.onVerifiersChosen);
        connect.addCallbackFunc("onAdvancePurchase", presetFunc.onAdvancePurchase);
        connect.addCallbackFunc("onConfirmPurchase", presetFunc.onConfirmPurchase);
        connect.addCallbackFunc("onReEncrypt", presetFunc.onReEncrypt);
        connect.addCallbackFunc("onFinishPurchase", presetFunc.onFinishPurchase);
        connect.addCallbackFunc("onRegisterVerifier", presetFunc.onRegisterVerifier);
        connect.addCallbackFunc("onVoteResult", presetFunc.onVoteResult);
        connect.addCallbackFunc("onVerifierDisable", presetFunc.onVerifierDisable);
        connect.addCallbackFunc("onArbitrationBegin", presetFunc.onArbitrationBegin);
        connect.addCallbackFunc("onArbitrationResult", presetFunc.onArbitrationResult);
    },
    setDefaultBalance: function (_this) {
        _this.$store.state.balance[0] = { Balance: "-", Time: "-"};
        _this.$store.state.balance[1] = { Balance: "-", Time: "-"};
    },
    timeout: function (ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    },
    setStateString: function (str) {
        let index;
        for (let i = 0; i < utils.stateEnum.length; i++) {
            if (str === utils.stateEnum[i]) {
                index = i;
                break;
            }
        }
        return index;
    },
    functionDisabled: function (funcNum, stateStr) {
        return !utils.state[funcNum][utils.setStateString(stateStr)];
    }
};

let presetFunc = {
    onPublish: function (payload, _this) {
        console.log("发布事件回调：", payload);
        _this.$notify({
            title: "发布事件回调：",
            message: "数据ID：" + payload.PublishId,
            position: "top-left"
        });
        dl_db.write({
            Title: payload.Title,
            Price: parseInt(payload.Price),
            Keys: payload.Keys,
            Description: payload.Description,
            Seller: payload.Seller,
            SupportVerify: payload.SupportVerify,
            MetaDataExtension: payload.extensions.metaDataExtension,
            ProofDataExtensions: payload.extensions.proofDataExtensions,
            PublishId: payload.PublishId    // keyPath
        }, function () {
            dl_db.init(_this);
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    },
    onProofFilesExtensions: function (payload, _this) {
        dl_db.read(payload, function (dlInstance) {
            connect.send({ Name:"extensions", Payload: {extensions: {proofDataExtensions: dlInstance.ProofDataExtensions}}}, function () {},
                function (payload, _this) {
                    _this.$alert(payload, "获取证明文件扩展名失败！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    });
                });
        });
    },
    onVerifiersChosen: function (payload, _this) {
        console.log("选择验证者事件回调：", payload);
        _this.$notify({
            title: "选择验证者事件回调：",
            message: "你已被选中成为 ID: " + payload.TransactionId + " 交易 的验证者。",
            position: "top-left"
        });
        dl_db.read(payload.PublishId, function (dataDetails) {
            tx_db.write({
                Title: dataDetails.Title,
                Price: dataDetails.Price,
                Keys: dataDetails.Keys,
                Description: dataDetails.Description,
                Buyer: "",
                Seller: dataDetails.Seller,
                State: payload.TransactionState, // -
                SupportVerify: dataDetails.SupportVerify,
                StartVerify: true, // - !
                MetaDataExtension: dataDetails.MetaDataExtension,
                ProofDataExtensions: dataDetails.ProofDataExtensions,
                MetaDataIdEncWithSeller: "",
                MetaDataIdEncWithBuyer: "",
                MetaDataIdEncWithArbitrator: "",
                Verifier1Response: "",
                Verifier2Response: "",
                ArbitrateResult: "",
                PublishId: dataDetails.PublishId,
                TransactionId: payload.TransactionId,    // keyPath
                Identify: 2
            }, function () {
                tx_db.initVerifier(_this);
            });
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    },
    onAdvancePurchase: function (payload, _this) {
        console.log("创建交易事件回调：", payload);
        _this.$notify({
            title: "创建交易事件回调：",
            message: "创建交易(ID：" + payload.TransactionId + ")成功",
            position: "top-left"
        });
        dl_db.read(payload.PublishId, function (dataDetails) {
            let param = {
                Title: dataDetails.Title,
                Price: dataDetails.Price,
                Keys: dataDetails.Keys,
                Description: dataDetails.Description,
                Buyer: payload.Buyer, // -
                Seller: dataDetails.Seller,
                State: payload.TransactionState, // -
                SupportVerify: dataDetails.SupportVerify,
                StartVerify: payload.StartVerify, // -
                MetaDataExtension: dataDetails.MetaDataExtension,
                ProofDataExtensions: dataDetails.ProofDataExtensions,
                MetaDataIdEncWithSeller: "",
                MetaDataIdEncWithBuyer: "",
                MetaDataIdEncWithArbitrator: "",
                Verifier1Response: "",
                Verifier2Response: "",
                ArbitrateResult: "",
                PublishId: dataDetails.PublishId,
                TransactionId: payload.TransactionId,    // keyPath
                Identify: 0
            };
            switch (_this.$store.state.account.toLowerCase()) {
                case dataDetails.Seller.toLowerCase():
                    tx_db.write(param, function () {
                        tx_db.initSeller(_this);
                    });break;
                default:
                    param.Identify = 1;
                    tx_db.write(param, function () {
                        tx_db.initBuyer(_this);
                    });break;
            }
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    },
    onConfirmPurchase: function (payload, _this) {
        console.log("购买数据事件回调：", payload);
        _this.$notify({
            title: "购买数据事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已确认购买。",
            position: "top-left"
        });
        tx_db.read(payload.TransactionId, function (txDetailsOnPurchase) {
            let param = {
                Title: txDetailsOnPurchase.Title,
                Price: txDetailsOnPurchase.Price,
                Keys: txDetailsOnPurchase.Keys,
                Description: txDetailsOnPurchase.Description,
                Buyer: txDetailsOnPurchase.Buyer,
                Seller: txDetailsOnPurchase.Seller,
                State: payload.TransactionState, // -
                SupportVerify: txDetailsOnPurchase.SupportVerify,
                StartVerify: txDetailsOnPurchase.StartVerify,
                MetaDataExtension: txDetailsOnPurchase.MetaDataExtension,
                ProofDataExtensions: txDetailsOnPurchase.ProofDataExtensions,
                MetaDataIdEncWithSeller: payload.EncryptedId.encryptedMetaDataId, // -
                MetaDataIdEncWithBuyer: txDetailsOnPurchase.MetaDataIdEncWithBuyer,
                MetaDataIdEncWithArbitrator: txDetailsOnPurchase.MetaDataIdEncWithArbitrator,
                Verifier1Response: txDetailsOnPurchase.Verifier1Response,
                Verifier2Response: txDetailsOnPurchase.Verifier2Response,
                ArbitrateResult: txDetailsOnPurchase.ArbitrateResult,
                PublishId: txDetailsOnPurchase.PublishId,
                TransactionId: txDetailsOnPurchase.TransactionId, // keyPath
                Identify: 0
            };
            switch (_this.$store.state.account.toLowerCase()) {
                case txDetailsOnPurchase.Seller.toLowerCase():
                    tx_db.write(param, function () {
                        tx_db.initSeller(_this);
                    });break;
                case txDetailsOnPurchase.Buyer.toLowerCase():
                    param.Identify = 1;
                    tx_db.write(param, function () {
                        tx_db.initBuyer(_this);
                    });break;
                default:
                    param.Identify = 2;
                    tx_db.write(param, function () {
                        tx_db.initVerifier(_this);
                    });break;
            }
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    },
    onReEncrypt: function (payload, _this) {
        console.log("再加密数据事件回调：", payload);
        _this.$notify({
            title: "再加密数据事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已重新加密。",
            position: "top-left"
        });
        tx_db.read(payload.TransactionId, function (txDetailsOnRFD) {
            let param = {
                Title: txDetailsOnRFD.Title,
                Price: txDetailsOnRFD.Price,
                Keys: txDetailsOnRFD.Keys,
                Description: txDetailsOnRFD.Description,
                Buyer: txDetailsOnRFD.Buyer,
                Seller: txDetailsOnRFD.Seller,
                State: payload.TransactionState, // -
                SupportVerify: txDetailsOnRFD.SupportVerify,
                StartVerify: txDetailsOnRFD.StartVerify,
                MetaDataExtension: txDetailsOnRFD.MetaDataExtension,
                ProofDataExtensions: txDetailsOnRFD.ProofDataExtensions,
                MetaDataIdEncWithSeller: txDetailsOnRFD.MetaDataIdEncWithSeller,
                MetaDataIdEncWithBuyer: payload.EncryptedId.encryptedMetaDataId, // -
                MetaDataIdEncWithArbitrator: txDetailsOnRFD.MetaDataIdEncWithArbitrator,
                Verifier1Response: txDetailsOnRFD.Verifier1Response,
                Verifier2Response: txDetailsOnRFD.Verifier2Response,
                ArbitrateResult: txDetailsOnRFD.ArbitrateResult,
                PublishId: txDetailsOnRFD.PublishId,
                TransactionId: txDetailsOnRFD.TransactionId,
                Identify: 0
            };
            switch (_this.$store.state.account) {
                case txDetailsOnRFD.Seller.toLowerCase():
                    tx_db.write(param, function () {
                        tx_db.initSeller(_this);
                    });break;
                default:
                    param.Identify = 1;
                    tx_db.write(param, function () {
                        tx_db.initBuyer(_this);
                    });break;
            }
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    },
    onFinishPurchase: function (payload, _this) {
        console.log("交易关闭事件回调：", payload);
        _this.$notify({
            title: "交易关闭事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已关闭。",
            position: "top-left"
        });
        tx_db.read(payload.TransactionId, function (txDetailsOnC) {
            let param = {
                Title: txDetailsOnC.Title,
                Price: txDetailsOnC.Price,
                Keys: txDetailsOnC.Keys,
                Description: txDetailsOnC.Description,
                Buyer: txDetailsOnC.Buyer,
                Seller: txDetailsOnC.Seller,
                State: payload.TransactionState, // -
                SupportVerify: txDetailsOnC.SupportVerify,
                StartVerify: txDetailsOnC.StartVerify,
                MetaDataExtension: txDetailsOnC.MetaDataExtension,
                ProofDataExtensions: txDetailsOnC.ProofDataExtensions,
                MetaDataIdEncWithSeller: txDetailsOnC.MetaDataIdEncWithSeller,
                MetaDataIdEncWithBuyer: txDetailsOnC.MetaDataIdEncWithBuyer,
                MetaDataIdEncWithArbitrator: txDetailsOnC.MetaDataIdEncWithArbitrator,
                Verifier1Response: txDetailsOnC.Verifier1Response,
                Verifier2Response: txDetailsOnC.Verifier2Response,
                ArbitrateResult: txDetailsOnC.ArbitrateResult,
                PublishId: txDetailsOnC.PublishId,
                TransactionId: txDetailsOnC.TransactionId,
                Identify: 0
            };
            switch (_this.$store.state.account) {
                case txDetailsOnC.Seller.toLowerCase():
                    tx_db.write(param, function () {
                        tx_db.initSeller(_this);
                    });break;
                case txDetailsOnC.Buyer.toLowerCase():
                    param.Identify = 1;
                    tx_db.write(param, function () {
                        tx_db.initBuyer(_this);
                    });break;
                default:
                    param.Identify = 2;
                    tx_db.write(param, function () {
                        tx_db.initVerifier(_this);
                    });break;
            }
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    },
    onRegisterVerifier: function (payload, _this) {
        console.log("注册成为验证者事件回调：", payload);
        _this.$notify({
            title: "注册成为验证者事件回调：",
            message: "你已成功注册成为验证者！ :)",
            position: "top-left"
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: true
            });
        });
    },
    onVoteResult: function (payload, _this) {
        console.log("验证者验证事件回调：", payload);
        _this.$notify({
            title: "验证者验证事件回调：",
            message: "收到新的验证者回复，交易ID：" + payload.TransactionId,
            position: "top-left"
        });
        tx_db.read(payload.TransactionId, function (txDetailsOnV) {
            let param = {
                Title: txDetailsOnV.Title,
                Price: txDetailsOnV.Price,
                Keys: txDetailsOnV.Keys,
                Description: txDetailsOnV.Description,
                Buyer: txDetailsOnV.Buyer,
                Seller: txDetailsOnV.Seller,
                State: payload.TransactionState, // -
                SupportVerify: txDetailsOnV.SupportVerify,
                StartVerify: txDetailsOnV.StartVerify,
                MetaDataExtension: txDetailsOnV.MetaDataExtension,
                ProofDataExtensions: txDetailsOnV.ProofDataExtensions,
                MetaDataIdEncWithSeller: txDetailsOnV.MetaDataIdEncWithSeller,
                MetaDataIdEncWithBuyer: txDetailsOnV.MetaDataIdEncWithBuyer,
                MetaDataIdEncWithArbitrator: txDetailsOnV.MetaDataIdEncWithArbitrator,
                Verifier1Response: txDetailsOnV.Verifier1Response, // -
                Verifier2Response: txDetailsOnV.Verifier2Response, // -
                ArbitrateResult: txDetailsOnV.ArbitrateResult,
                PublishId: txDetailsOnV.PublishId,
                TransactionId: txDetailsOnV.TransactionId,
                Identify: 1
            };
            switch (parseInt(payload.VerifyResult.VerifierIndex)) {
                case 0:
                    param.Identify = 2;
                    tx_db.write(param, function () {
                        tx_db.initVerifier(_this);
                    });break;
                case 1:
                    param.Verifier1Response = payload.VerifyResult.VerifierResponse;
                    tx_db.write(param, function () {
                        tx_db.initBuyer(_this);
                    });break;
                case 2:
                    param.Verifier2Response = payload.VerifyResult.VerifierResponse;
                    tx_db.write(param, function () {
                        tx_db.initBuyer(_this);
                    });break;
            }
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    },
    onVerifierDisable: function (payload, _this) {
        console.log("取消验证者验证资格事件回调：", payload);
        _this.$notify({
            title: "取消验证者验证资格事件回调：",
            message: "验证者： " + payload.VerifierDisabled.VerifierAddress + "被取消验证资格。",
            position: "top-left"
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: false
            });
        });
    },
    onArbitrationBegin: function (payload, _this) {
        console.log("仲裁开始事件回调：", payload);
        _this.$notify({
            title: "仲裁开始事件回调：",
            message: "你已被选中成为 ID: " + payload.TransactionId + " 交易 的仲裁者。",
            position: "top-left"
        });
        dl_db.read(payload.PublishId, function (dataDetails) {
            tx_db.write({
                Title: dataDetails.Title,
                Price: dataDetails.Price,
                Keys: dataDetails.Keys,
                Description: dataDetails.Description,
                Buyer: "",
                Seller: dataDetails.Seller,
                State: "ReadyForDownload", // - !
                SupportVerify: dataDetails.SupportVerify,
                StartVerify: true, // - !
                MetaDataExtension: dataDetails.MetaDataExtension,
                ProofDataExtensions: dataDetails.ProofDataExtensions,
                MetaDataIdEncWithSeller: "",
                MetaDataIdEncWithBuyer: "",
                MetaDataIdEncWithArbitrator: payload.EncryptedId.encryptedMetaDataId,
                Verifier1Response: "",
                Verifier2Response: "",
                ArbitrateResult: "",
                PublishId: dataDetails.PublishId,
                TransactionId: payload.TransactionId,    // keyPath
                Identify: 3
            }, function () {
                tx_db.initArbitrator(_this);
            });
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    },
    onArbitrationResult: function (payload, _this) {
        console.log("仲裁结果事件回调：", payload);
        _this.$notify({
            title: "仲裁结果事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已完成仲裁，仲裁结果为：" + payload.Arbitrate.ArbitrateResult + " 。",
            position: "top-left"
        });
        tx_db.read(payload.TransactionId, function (txDetailsOAR) {
            let param = {
                Title: txDetailsOAR.Title,
                Price: txDetailsOAR.Price,
                Keys: txDetailsOAR.Keys,
                Description: txDetailsOAR.Description,
                Buyer: txDetailsOAR.Buyer,
                Seller: txDetailsOAR.Seller,
                State: txDetailsOAR.State,
                SupportVerify: txDetailsOAR.SupportVerify,
                StartVerify: txDetailsOAR.StartVerify,
                MetaDataExtension: txDetailsOAR.MetaDataExtension,
                ProofDataExtensions: txDetailsOAR.ProofDataExtensions,
                MetaDataIdEncWithSeller: txDetailsOAR.MetaDataIdEncWithSeller,
                MetaDataIdEncWithBuyer: txDetailsOAR.MetaDataIdEncWithBuyer,
                MetaDataIdEncWithArbitrator: txDetailsOAR.MetaDataIdEncWithArbitrator,
                Verifier1Response: txDetailsOAR.Verifier1Response,
                Verifier2Response: txDetailsOAR.Verifier2Response,
                ArbitrateResult: payload.Arbitrate.ArbitrateResult, // -
                PublishId: txDetailsOAR.PublishId,
                TransactionId: txDetailsOAR.TransactionId, // keyPath
                Identify: 0
            };
            switch (_this.$store.state.account) {
                case txDetailsOAR.Seller.toLowerCase():
                    tx_db.write(param, function () {
                        tx_db.initSeller(_this);
                    });break;
                default:
                    param.Identify = 1;
                    tx_db.write(param, function () {
                        tx_db.initBuyer(_this);
                    });break;
            }
        });
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                nickname: accInstance.nickname,
                fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                isVerifier: accInstance.isVerifier
            });
        });
    }
};

export {utils}