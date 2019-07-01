import {acc_db, dl_db, txBuyer_db, txSeller_db, txVerifier_db} from "./DBoptions";
import {connect} from "./connect";

let utils = {
    voteMutex: true,
    voteWait: 0,
    voteParams: [],
    init: function () {
        connect.addCallbackFunc("onPublish", presetFunc.onPublish);
        connect.addCallbackFunc("onProofFilesExtensions", presetFunc.onProofFilesExtensions);
        connect.addCallbackFunc("onVerifiersChosen", presetFunc.onVerifiersChosen);
        connect.addCallbackFunc("onTransactionCreate", presetFunc.onTransactionCreate);
        connect.addCallbackFunc("onPurchase", presetFunc.onPurchase);
        connect.addCallbackFunc("onReadyForDownload", presetFunc.onReadyForDownload);
        connect.addCallbackFunc("onClose", presetFunc.onClose);
        connect.addCallbackFunc("onRegisterVerifier", presetFunc.onRegisterVerifier);
        connect.addCallbackFunc("onVote", presetFunc.onVote);
        connect.addCallbackFunc("onVerifierDisable", presetFunc.onVerifierDisable);
        connect.addCallbackFunc("onArbitrationBegin", presetFunc.onArbitrationBegin);
        connect.addCallbackFunc("onArbitrationResult", presetFunc.onArbitrationResult);
    },
    setDefaultBalance: function (_this) {
        _this.$store.state.balance[0] = { Balance: "-", Time: "-"};
        _this.$store.state.balance[1] = { Balance: "-", Time: "-"};
    },
    vote: function (_this, payload) {
        console.log("验证者验证事件回调：", payload);
        _this.$notify({
            title: "验证者验证事件回调：",
            message: payload,
            position: "top-left"
        });
        if (payload.VerifierIndex === "0") {
            txVerifier_db.read(payload.TransactionID, function (txDetailsOnV) {
                txVerifier_db.write({
                    Title: txDetailsOnV.Title,
                    Price: txDetailsOnV.Price,
                    Keys: txDetailsOnV.Keys,
                    Description: txDetailsOnV.Description,
                    Buyer: txDetailsOnV.Buyer,
                    Seller: txDetailsOnV.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnV.SupportVerify,
                    StartVerify: txDetailsOnV.StartVerify,
                    MetaDataExtension: txDetailsOnV.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnV.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnV.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnV.MetaDataIDEncWithBuyer,
                    MetaDataIDEncWithArbitrator: txDetailsOnV.MetaDataIDEncWithArbitrator,
                    Verifier1Response: payload.VerifierResponse, // -
                    Verifier2Response: txDetailsOnV.Verifier2Response,
                    ArbitrateResult: txDetailsOnV.ArbitrateResult,
                    PublishID: txDetailsOnV.PublishID,
                    TransactionID: txDetailsOnV.TransactionID
                }, function () {
                    txVerifier_db.init(_this);
                    utils.voteMutex = true;
                    if (utils.voteWait > 0) {
                        utils.voteWait--;
                        utils.vote(_this, utils.voteParams.shift());
                    }
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        }
        if (payload.VerifierIndex === "1") {
            txBuyer_db.read(payload.TransactionID, function (txDetailsOnV) {
                txBuyer_db.write({
                    Title: txDetailsOnV.Title,
                    Price: txDetailsOnV.Price,
                    Keys: txDetailsOnV.Keys,
                    Description: txDetailsOnV.Description,
                    Buyer: txDetailsOnV.Buyer,
                    Seller: txDetailsOnV.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnV.SupportVerify,
                    StartVerify: txDetailsOnV.StartVerify,
                    MetaDataExtension: txDetailsOnV.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnV.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnV.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnV.MetaDataIDEncWithBuyer,
                    MetaDataIDEncWithArbitrator: txDetailsOnV.MetaDataIDEncWithArbitrator,
                    Verifier1Response: payload.VerifierResponse, // -
                    Verifier2Response: txDetailsOnV.Verifier2Response,
                    ArbitrateResult: txDetailsOnV.ArbitrateResult,
                    PublishID: txDetailsOnV.PublishID,
                    TransactionID: txDetailsOnV.TransactionID
                }, function () {
                    txBuyer_db.init(_this);
                    utils.voteMutex = true;
                    if (utils.voteWait > 0) {
                        utils.voteWait--;
                        utils.vote(_this, utils.voteParams.shift());
                    }
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        }
        if (payload.VerifierIndex === "2") {
            txBuyer_db.read(payload.TransactionID, function (txDetailsOnV) {
                txBuyer_db.write({
                    Title: txDetailsOnV.Title,
                    Price: txDetailsOnV.Price,
                    Keys: txDetailsOnV.Keys,
                    Description: txDetailsOnV.Description,
                    Buyer: txDetailsOnV.Buyer,
                    Seller: txDetailsOnV.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnV.SupportVerify,
                    StartVerify: txDetailsOnV.StartVerify,
                    MetaDataExtension: txDetailsOnV.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnV.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnV.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnV.MetaDataIDEncWithBuyer,
                    MetaDataIDEncWithArbitrator: txDetailsOnV.MetaDataIDEncWithArbitrator,
                    Verifier1Response: txDetailsOnV.Verifier1Response,
                    Verifier2Response: payload.VerifierResponse, // -
                    ArbitrateResult: txDetailsOnV.ArbitrateResult,
                    PublishID: txDetailsOnV.PublishID,
                    TransactionID: txDetailsOnV.TransactionID
                }, function () {
                    txBuyer_db.init(_this);
                    utils.voteMutex = true;
                    if (utils.voteWait > 0) {
                        utils.voteWait--;
                        utils.vote(_this, utils.voteParams.shift());
                    }
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        }
    }
};

let presetFunc = {
    onPublish: function (payload, _this) {
        console.log("发布事件回调：", payload);
        _this.$notify({
            title: "发布事件回调：",
            message: payload,
            position: "top-left"
        });
        dl_db.write({
            Title: payload.Title,
            Price: parseInt(payload.Price),
            Keys: payload.Keys,
            Description: payload.Description,
            Seller: payload.Seller,
            SupportVerify: payload.SupportVerify,
            MetaDataExtension: payload.MetaDataExtension,
            ProofDataExtensions: payload.ProofDataExtensions,
            PublishID: payload.PublishID    // keyPath
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
            connect.send({ Name:"extensions", Payload: {extensions: dlInstance.ProofDataExtensions}}, function () {},
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
            message: payload,
            position: "top-left"
        });
        dl_db.read(payload.PublishID, function (dataDetails) {
            txVerifier_db.write({
                Title: dataDetails.Title,
                Price: dataDetails.Price,
                Keys: dataDetails.Keys,
                Description: dataDetails.Description,
                Buyer: "",
                Seller: dataDetails.Seller,
                State: payload.TxState, // -
                SupportVerify: dataDetails.SupportVerify,
                StartVerify: true, // - !
                MetaDataExtension: dataDetails.MetaDataExtension,
                ProofDataExtensions: dataDetails.ProofDataExtensions,
                MetaDataIDEncWithSeller: "",
                MetaDataIDEncWithBuyer: "",
                Verifier1Response: "",
                Verifier2Response: "",
                PublishID: dataDetails.PublishID,
                TransactionID: payload.TransactionID    // keyPath
            }, function () {
                txVerifier_db.init(_this);
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
    onTransactionCreate: function (payload, _this) {
        console.log("创建交易事件回调：", payload);
        _this.$notify({
            title: "创建交易事件回调：",
            message: payload,
            position: "top-left"
        });
        dl_db.read(payload.PublishID, function (dataDetails) {
            txBuyer_db.write({
                Title: dataDetails.Title,
                Price: dataDetails.Price,
                Keys: dataDetails.Keys,
                Description: dataDetails.Description,
                Buyer: payload.Buyer, // -
                Seller: dataDetails.Seller,
                State: payload.TxState, // -
                SupportVerify: dataDetails.SupportVerify,
                StartVerify: payload.StartVerify, // -
                MetaDataExtension: dataDetails.MetaDataExtension,
                ProofDataExtensions: dataDetails.ProofDataExtensions,
                MetaDataIDEncWithSeller: "",
                MetaDataIDEncWithBuyer: "",
                Verifier1Response: "",
                Verifier2Response: "",
                PublishID: dataDetails.PublishID,
                TransactionID: payload.TransactionID    // keyPath
            }, function () {
                txBuyer_db.init(_this);
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
    onPurchase: function (payload, _this) {
        console.log("购买数据事件回调：", payload);
        _this.$notify({
            title: "购买数据事件回调：",
            message: payload,
            position: "top-left"
        });
        if (payload.UserIndex === "0") {
            dl_db.read(payload.PublishID, function (dataDetails) {
                txSeller_db.write({
                    Title: dataDetails.Title,
                    Price: dataDetails.Price,
                    Keys: dataDetails.Keys,
                    Description: dataDetails.Description,
                    Buyer: "",
                    Seller: dataDetails.Seller,
                    State: payload.TxState, // -
                    SupportVerify: dataDetails.SupportVerify,
                    StartVerify: false,
                    MetaDataExtension: dataDetails.MetaDataExtension,
                    ProofDataExtensions: dataDetails.ProofDataExtensions,
                    MetaDataIDEncWithSeller: payload.MetaDataIdEncWithSeller, // -
                    MetaDataIDEncWithBuyer: "",
                    Verifier1Response: "",
                    Verifier2Response: "",
                    PublishID: dataDetails.PublishID,
                    TransactionID: payload.TransactionID // keyPath
                },function () {
                    txSeller_db.init(_this);
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        } else if (payload.UserIndex === "1") {
            txBuyer_db.read(payload.TransactionID, function (txDetailsOnPurchase) {
                txBuyer_db.write({
                    Title: txDetailsOnPurchase.Title,
                    Price: txDetailsOnPurchase.Price,
                    Keys: txDetailsOnPurchase.Keys,
                    Description: txDetailsOnPurchase.Description,
                    Buyer: txDetailsOnPurchase.Buyer,
                    Seller: txDetailsOnPurchase.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnPurchase.SupportVerify,
                    StartVerify: txDetailsOnPurchase.StartVerify,
                    MetaDataExtension: txDetailsOnPurchase.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnPurchase.ProofDataExtensions,
                    MetaDataIDEncWithSeller: payload.MetaDataIdEncWithSeller, // -
                    MetaDataIDEncWithBuyer: txDetailsOnPurchase.MetaDataIDEncWithBuyer,
                    Verifier1Response: txDetailsOnPurchase.Verifier1Response,
                    Verifier2Response: txDetailsOnPurchase.Verifier2Response,
                    PublishID: txDetailsOnPurchase.PublishID,
                    TransactionID: txDetailsOnPurchase.TransactionID // keyPath
                },function () {
                    txBuyer_db.init(_this);
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        } else {
            txVerifier_db.read(payload.TransactionID, function (txDetailsOnPurchase) {
                txVerifier_db.write({
                    Title: txDetailsOnPurchase.Title,
                    Price: txDetailsOnPurchase.Price,
                    Keys: txDetailsOnPurchase.Keys,
                    Description: txDetailsOnPurchase.Description,
                    Buyer: txDetailsOnPurchase.Buyer,
                    Seller: txDetailsOnPurchase.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnPurchase.SupportVerify,
                    StartVerify: txDetailsOnPurchase.StartVerify,
                    MetaDataExtension: txDetailsOnPurchase.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnPurchase.ProofDataExtensions,
                    MetaDataIDEncWithSeller: payload.MetaDataIdEncWithSeller, // -
                    MetaDataIDEncWithBuyer: txDetailsOnPurchase.MetaDataIDEncWithBuyer,
                    Verifier1Response: txDetailsOnPurchase.Verifier1Response,
                    Verifier2Response: txDetailsOnPurchase.Verifier2Response,
                    PublishID: txDetailsOnPurchase.PublishID,
                    TransactionID: txDetailsOnPurchase.TransactionID // keyPath
                },function () {
                    txVerifier_db.init(_this);
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        }
    },
    onReadyForDownload: function (payload, _this) {
        console.log("再加密数据事件回调：", payload);
        _this.$notify({
            title: "再加密数据事件回调：",
            message: payload,
            position: "top-left"
        });
        if (payload.UserIndex === "0") {
            txSeller_db.read(payload.TransactionID, function (txDetailsOnRFD) {
                txSeller_db.write({
                    Title: txDetailsOnRFD.Title,
                    Price: txDetailsOnRFD.Price,
                    Keys: txDetailsOnRFD.Keys,
                    Description: txDetailsOnRFD.Description,
                    Buyer: txDetailsOnRFD.Buyer,
                    Seller: txDetailsOnRFD.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnRFD.SupportVerify,
                    StartVerify: txDetailsOnRFD.StartVerify,
                    MetaDataExtension: txDetailsOnRFD.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnRFD.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnRFD.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: payload.MetaDataIdEncWithBuyer, // -
                    Verifier1Response: txDetailsOnRFD.Verifier1Response,
                    Verifier2Response: txDetailsOnRFD.Verifier2Response,
                    PublishID: txDetailsOnRFD.PublishID,
                    TransactionID: txDetailsOnRFD.TransactionID
                }, function () {
                    txSeller_db.init(_this);
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        } else if (payload.UserIndex === "1") {
            txBuyer_db.read(payload.TransactionID, function (txDetailsOnRFD) {
                txBuyer_db.write({
                    Title: txDetailsOnRFD.Title,
                    Price: txDetailsOnRFD.Price,
                    Keys: txDetailsOnRFD.Keys,
                    Description: txDetailsOnRFD.Description,
                    Buyer: txDetailsOnRFD.Buyer,
                    Seller: txDetailsOnRFD.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnRFD.SupportVerify,
                    StartVerify: txDetailsOnRFD.StartVerify,
                    MetaDataExtension: txDetailsOnRFD.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnRFD.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnRFD.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: payload.MetaDataIdEncWithBuyer, // -
                    Verifier1Response: txDetailsOnRFD.Verifier1Response,
                    Verifier2Response: txDetailsOnRFD.Verifier2Response,
                    PublishID: txDetailsOnRFD.PublishID,
                    TransactionID: txDetailsOnRFD.TransactionID
                }, function () {
                    txBuyer_db.init(_this);
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        }
    },
    onClose: function (payload, _this) {
        console.log("交易关闭事件回调：", payload);
        _this.$notify({
            title: "交易关闭事件回调：",
            message: payload,
            position: "top-left"
        });
        if (payload.UserIndex === "0") {
            txSeller_db.read(payload.TransactionID, function (txDetailsOnC) {
                txSeller_db.write({
                    Title: txDetailsOnC.Title,
                    Price: txDetailsOnC.Price,
                    Keys: txDetailsOnC.Keys,
                    Description: txDetailsOnC.Description,
                    Buyer: txDetailsOnC.Buyer,
                    Seller: txDetailsOnC.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnC.SupportVerify,
                    StartVerify: txDetailsOnC.StartVerify,
                    MetaDataExtension: txDetailsOnC.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnC.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnC.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnC.MetaDataIDEncWithBuyer,
                    Verifier1Response: txDetailsOnC.Verifier1Response,
                    Verifier2Response: txDetailsOnC.Verifier2Response,
                    PublishID: txDetailsOnC.PublishID,
                    TransactionID: txDetailsOnC.TransactionID
                }, function () {
                    txSeller_db.init(_this);
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        } else if (payload.UserIndex === "1") {
            txBuyer_db.read(payload.TransactionID, function (txDetailsOnC) {
                txBuyer_db.write({
                    Title: txDetailsOnC.Title,
                    Price: txDetailsOnC.Price,
                    Keys: txDetailsOnC.Keys,
                    Description: txDetailsOnC.Description,
                    Buyer: txDetailsOnC.Buyer,
                    Seller: txDetailsOnC.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnC.SupportVerify,
                    StartVerify: txDetailsOnC.StartVerify,
                    MetaDataExtension: txDetailsOnC.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnC.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnC.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnC.MetaDataIDEncWithBuyer,
                    Verifier1Response: txDetailsOnC.Verifier1Response,
                    Verifier2Response: txDetailsOnC.Verifier2Response,
                    PublishID: txDetailsOnC.PublishID,
                    TransactionID: txDetailsOnC.TransactionID
                }, function () {
                    txBuyer_db.init(_this);
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        } else {
            txVerifier_db.read(payload.TransactionID, function (txDetailsOnC) {
                txVerifier_db.write({
                    Title: txDetailsOnC.Title,
                    Price: txDetailsOnC.Price,
                    Keys: txDetailsOnC.Keys,
                    Description: txDetailsOnC.Description,
                    Buyer: txDetailsOnC.Buyer,
                    Seller: txDetailsOnC.Seller,
                    State: payload.TxState, // -
                    SupportVerify: txDetailsOnC.SupportVerify,
                    StartVerify: txDetailsOnC.StartVerify,
                    MetaDataExtension: txDetailsOnC.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnC.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnC.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnC.MetaDataIDEncWithBuyer,
                    Verifier1Response: txDetailsOnC.Verifier1Response,
                    Verifier2Response: txDetailsOnC.Verifier2Response,
                    PublishID: txDetailsOnC.PublishID,
                    TransactionID: txDetailsOnC.TransactionID
                }, function () {
                    txVerifier_db.init(_this);
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            nickname: accInstance.nickname,
                            fromBlock: Math.max(accInstance.fromBlock, payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        });
                    });
                });
            });
        }
    },
    onRegisterVerifier: function (payload, _this) {
        console.log("注册成为验证者事件回调：", payload);
        _this.$notify({
            title: "注册成为验证者事件回调：",
            message: payload,
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
    onVote: function (payload, _this) {
        if (utils.voteMutex === true) {
            utils.voteMutex = false;
            utils.vote(_this, payload);
        } else {
            utils.voteWait++;
            utils.voteParams.push(payload);
        }
    },
    onVerifierDisable: function (payload, _this) {
        console.log("取消验证者验证资格事件回调：", payload);
        _this.$notify({
            title: "取消验证者验证资格事件回调：",
            message: payload,
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

    },
    onArbitrationResult: function (payload, _this) {

    }
};

export {utils}