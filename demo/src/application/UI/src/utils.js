import {dl_db, acc_db, txBuyer_db, txSeller_db, txVerifier_db} from "./DBoptions.js"
let utils = {
    voteMutex: true,
    voteWait: 0,
    voteParams: [],
    listen: function (_this) {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "welcome":
                    _this.$notify({
                        title: "通知: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    acc_db.init(_this)
                    break
                case "resetChain":
                    dl_db.reset()
                    acc_db.reset()
                    txBuyer_db.reset()
                    txSeller_db.reset()
                    txVerifier_db.reset()
                    console.log("重置完成")
                    break
                case "initDL":
                    dl_db.init(_this)
                    console.log("数据列表初始化完成")
                    break
                case "initTx":
                    txBuyer_db.init(_this)
                    txSeller_db.init(_this)
                    txVerifier_db.init(_this)
                    console.log("交易列表初始化完成")
                    break
                case "onPublish":
                    console.log("发布事件回调：", message.payload)
                    _this.$notify({
                        title: "发布事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    dl_db.write({
                        Title: message.payload.Title,
                        Price: parseInt(message.payload.Price),
                        Keys: message.payload.Keys,
                        Description: message.payload.Description,
                        Seller: message.payload.Seller,
                        SupportVerify: message.payload.SupportVerify,
                        MetaDataExtension: message.payload.MetaDataExtension,
                        ProofDataExtensions: message.payload.ProofDataExtensions,
                        PublishID: message.payload.PublishID    // keyPath
                    }, function () {
                        dl_db.init(_this)
                    })
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onProofFilesExtensions":
                    dl_db.read(message.payload, function (dlInstance) {
                        astilectron.sendMessage({ Name:"extensions", Payload: {extensions: dlInstance.ProofDataExtensions}})
                    })
                    break
                case "onVerifiersChosen":
                    console.log("选择验证者事件回调：", message.payload)
                    _this.$notify({
                        title: "选择验证者事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    dl_db.read(message.payload.PublishID, function (dataDetails) {
                        txVerifier_db.write({
                            Title: dataDetails.Title,
                            Price: dataDetails.Price,
                            Keys: dataDetails.Keys,
                            Description: dataDetails.Description,
                            Buyer: "",
                            Seller: dataDetails.Seller,
                            State: message.payload.TxState, // -
                            SupportVerify: dataDetails.SupportVerify,
                            StartVerify: true, // - !
                            MetaDataExtension: dataDetails.MetaDataExtension,
                            ProofDataExtensions: dataDetails.ProofDataExtensions,
                            MetaDataIDEncWithSeller: "",
                            MetaDataIDEncWithBuyer: "",
                            Verifier1Response: "",
                            Verifier2Response: "",
                            PublishID: dataDetails.PublishID,
                            TransactionID: message.payload.TransactionID    // keyPath
                        }, function () {
                            txVerifier_db.init(_this)
                        })
                    })
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onTransactionCreate":
                    console.log("创建交易事件回调：", message.payload)
                    _this.$notify({
                        title: "创建交易事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    dl_db.read(message.payload.PublishID, function (dataDetails) {
                        txBuyer_db.write({
                            Title: dataDetails.Title,
                            Price: dataDetails.Price,
                            Keys: dataDetails.Keys,
                            Description: dataDetails.Description,
                            Buyer: message.payload.Buyer, // -
                            Seller: dataDetails.Seller,
                            State: message.payload.TxState, // -
                            SupportVerify: dataDetails.SupportVerify,
                            StartVerify: message.payload.StartVerify, // -
                            MetaDataExtension: dataDetails.MetaDataExtension,
                            ProofDataExtensions: dataDetails.ProofDataExtensions,
                            MetaDataIDEncWithSeller: "",
                            MetaDataIDEncWithBuyer: "",
                            Verifier1Response: "",
                            Verifier2Response: "",
                            PublishID: dataDetails.PublishID,
                            TransactionID: message.payload.TransactionID    // keyPath
                        }, function () {
                            txBuyer_db.init(_this)
                        })
                    })
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onPurchase":
                    console.log("购买数据事件回调：", message.payload)
                    _this.$notify({
                        title: "购买数据事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    if (message.payload.UserIndex === "0") {
                        dl_db.read(message.payload.PublishID, function (dataDetails) {
                            txSeller_db.write({
                                Title: dataDetails.Title,
                                Price: dataDetails.Price,
                                Keys: dataDetails.Keys,
                                Description: dataDetails.Description,
                                Buyer: "",
                                Seller: dataDetails.Seller,
                                State: message.payload.TxState, // -
                                SupportVerify: dataDetails.SupportVerify,
                                StartVerify: false,
                                MetaDataExtension: dataDetails.MetaDataExtension,
                                ProofDataExtensions: dataDetails.ProofDataExtensions,
                                MetaDataIDEncWithSeller: message.payload.MetaDataIdEncWithSeller, // -
                                MetaDataIDEncWithBuyer: "",
                                Verifier1Response: "",
                                Verifier2Response: "",
                                PublishID: dataDetails.PublishID,
                                TransactionID: message.payload.TransactionID // keyPath
                            },function () {
                                txSeller_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    } else if (message.payload.UserIndex === "1") {
                        txBuyer_db.read(message.payload.TransactionID, function (txDetailsOnPurchase) {
                            txBuyer_db.write({
                                Title: txDetailsOnPurchase.Title,
                                Price: txDetailsOnPurchase.Price,
                                Keys: txDetailsOnPurchase.Keys,
                                Description: txDetailsOnPurchase.Description,
                                Buyer: txDetailsOnPurchase.Buyer,
                                Seller: txDetailsOnPurchase.Seller,
                                State: message.payload.TxState, // -
                                SupportVerify: txDetailsOnPurchase.SupportVerify,
                                StartVerify: txDetailsOnPurchase.StartVerify,
                                MetaDataExtension: txDetailsOnPurchase.MetaDataExtension,
                                ProofDataExtensions: txDetailsOnPurchase.ProofDataExtensions,
                                MetaDataIDEncWithSeller: message.payload.MetaDataIdEncWithSeller, // -
                                MetaDataIDEncWithBuyer: txDetailsOnPurchase.MetaDataIDEncWithBuyer,
                                Verifier1Response: txDetailsOnPurchase.Verifier1Response,
                                Verifier2Response: txDetailsOnPurchase.Verifier2Response,
                                PublishID: txDetailsOnPurchase.PublishID,
                                TransactionID: txDetailsOnPurchase.TransactionID // keyPath
                            },function () {
                                txBuyer_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    } else {
                        txVerifier_db.read(message.payload.TransactionID, function (txDetailsOnPurchase) {
                            txVerifier_db.write({
                                Title: txDetailsOnPurchase.Title,
                                Price: txDetailsOnPurchase.Price,
                                Keys: txDetailsOnPurchase.Keys,
                                Description: txDetailsOnPurchase.Description,
                                Buyer: txDetailsOnPurchase.Buyer,
                                Seller: txDetailsOnPurchase.Seller,
                                State: message.payload.TxState, // -
                                SupportVerify: txDetailsOnPurchase.SupportVerify,
                                StartVerify: txDetailsOnPurchase.StartVerify,
                                MetaDataExtension: txDetailsOnPurchase.MetaDataExtension,
                                ProofDataExtensions: txDetailsOnPurchase.ProofDataExtensions,
                                MetaDataIDEncWithSeller: message.payload.MetaDataIdEncWithSeller, // -
                                MetaDataIDEncWithBuyer: txDetailsOnPurchase.MetaDataIDEncWithBuyer,
                                Verifier1Response: txDetailsOnPurchase.Verifier1Response,
                                Verifier2Response: txDetailsOnPurchase.Verifier2Response,
                                PublishID: txDetailsOnPurchase.PublishID,
                                TransactionID: txDetailsOnPurchase.TransactionID // keyPath
                            },function () {
                                txVerifier_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    }
                    break
                case "onReadyForDownload":
                    console.log("再加密数据事件回调：", message.payload)
                    _this.$notify({
                        title: "再加密数据事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    if (message.payload.UserIndex === "0") {
                        txSeller_db.read(message.payload.TransactionID, function (txDetailsOnRFD) {
                            txSeller_db.write({
                                Title: txDetailsOnRFD.Title,
                                Price: txDetailsOnRFD.Price,
                                Keys: txDetailsOnRFD.Keys,
                                Description: txDetailsOnRFD.Description,
                                Buyer: txDetailsOnRFD.Buyer,
                                Seller: txDetailsOnRFD.Seller,
                                State: message.payload.TxState, // -
                                SupportVerify: txDetailsOnRFD.SupportVerify,
                                StartVerify: txDetailsOnRFD.StartVerify,
                                MetaDataExtension: txDetailsOnRFD.MetaDataExtension,
                                ProofDataExtensions: txDetailsOnRFD.ProofDataExtensions,
                                MetaDataIDEncWithSeller: txDetailsOnRFD.MetaDataIDEncWithSeller,
                                MetaDataIDEncWithBuyer: message.payload.MetaDataIdEncWithBuyer, // -
                                Verifier1Response: txDetailsOnRFD.Verifier1Response,
                                Verifier2Response: txDetailsOnRFD.Verifier2Response,
                                PublishID: txDetailsOnRFD.PublishID,
                                TransactionID: txDetailsOnRFD.TransactionID
                            }, function () {
                                txSeller_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    } else if (message.payload.UserIndex === "1") {
                        txBuyer_db.read(message.payload.TransactionID, function (txDetailsOnRFD) {
                            txBuyer_db.write({
                                Title: txDetailsOnRFD.Title,
                                Price: txDetailsOnRFD.Price,
                                Keys: txDetailsOnRFD.Keys,
                                Description: txDetailsOnRFD.Description,
                                Buyer: txDetailsOnRFD.Buyer,
                                Seller: txDetailsOnRFD.Seller,
                                State: message.payload.TxState, // -
                                SupportVerify: txDetailsOnRFD.SupportVerify,
                                StartVerify: txDetailsOnRFD.StartVerify,
                                MetaDataExtension: txDetailsOnRFD.MetaDataExtension,
                                ProofDataExtensions: txDetailsOnRFD.ProofDataExtensions,
                                MetaDataIDEncWithSeller: txDetailsOnRFD.MetaDataIDEncWithSeller,
                                MetaDataIDEncWithBuyer: message.payload.MetaDataIdEncWithBuyer, // -
                                Verifier1Response: txDetailsOnRFD.Verifier1Response,
                                Verifier2Response: txDetailsOnRFD.Verifier2Response,
                                PublishID: txDetailsOnRFD.PublishID,
                                TransactionID: txDetailsOnRFD.TransactionID
                            }, function () {
                                txBuyer_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    } else {
                        txVerifier_db.read(message.payload.TransactionID, function (txDetailsOnRFD) {
                            txVerifier_db.write({
                                Title: txDetailsOnRFD.Title,
                                Price: txDetailsOnRFD.Price,
                                Keys: txDetailsOnRFD.Keys,
                                Description: txDetailsOnRFD.Description,
                                Buyer: txDetailsOnRFD.Buyer,
                                Seller: txDetailsOnRFD.Seller,
                                State: message.payload.TxState, // -
                                SupportVerify: txDetailsOnRFD.SupportVerify,
                                StartVerify: txDetailsOnRFD.StartVerify,
                                MetaDataExtension: txDetailsOnRFD.MetaDataExtension,
                                ProofDataExtensions: txDetailsOnRFD.ProofDataExtensions,
                                MetaDataIDEncWithSeller: txDetailsOnRFD.MetaDataIDEncWithSeller,
                                MetaDataIDEncWithBuyer: message.payload.MetaDataIdEncWithBuyer, // -
                                Verifier1Response: txDetailsOnRFD.Verifier1Response,
                                Verifier2Response: txDetailsOnRFD.Verifier2Response,
                                PublishID: txDetailsOnRFD.PublishID,
                                TransactionID: txDetailsOnRFD.TransactionID
                            }, function () {
                                txVerifier_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    }
                    break
                case "onClose":
                    console.log("交易关闭事件回调：", message.payload)
                    _this.$notify({
                        title: "交易关闭事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    if (message.payload.UserIndex === "0") {
                        txSeller_db.read(message.payload.TransactionID, function (txDetailsOnC) {
                            txSeller_db.write({
                                Title: txDetailsOnC.Title,
                                Price: txDetailsOnC.Price,
                                Keys: txDetailsOnC.Keys,
                                Description: txDetailsOnC.Description,
                                Buyer: txDetailsOnC.Buyer,
                                Seller: txDetailsOnC.Seller,
                                State: message.payload.TxState, // -
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
                                txSeller_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    } else if (message.payload.UserIndex === "1") {
                        txBuyer_db.read(message.payload.TransactionID, function (txDetailsOnC) {
                            txBuyer_db.write({
                                Title: txDetailsOnC.Title,
                                Price: txDetailsOnC.Price,
                                Keys: txDetailsOnC.Keys,
                                Description: txDetailsOnC.Description,
                                Buyer: txDetailsOnC.Buyer,
                                Seller: txDetailsOnC.Seller,
                                State: message.payload.TxState, // -
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
                                txBuyer_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    } else {
                        txVerifier_db.read(message.payload.TransactionID, function (txDetailsOnC) {
                            txVerifier_db.write({
                                Title: txDetailsOnC.Title,
                                Price: txDetailsOnC.Price,
                                Keys: txDetailsOnC.Keys,
                                Description: txDetailsOnC.Description,
                                Buyer: txDetailsOnC.Buyer,
                                Seller: txDetailsOnC.Seller,
                                State: message.payload.TxState, // -
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
                                txVerifier_db.init(_this)
                                acc_db.read(_this.$store.state.account, function (accInstance) {
                                    acc_db.write({
                                        address: accInstance.address,
                                        fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                                        isVerifier: accInstance.isVerifier
                                    })
                                })
                            })
                        })
                    }
                    break
                case "onRegisterVerifier":
                    console.log("注册成为验证者事件回调：", message.payload)
                    _this.$notify({
                        title: "注册成为验证者事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                            isVerifier: true
                        })
                    })
                    break
                case "onVote":
                    if (utils.voteMutex === true) {
                        utils.voteMutex = false
                        utils.vote(_this, message)
                    } else {
                        utils.voteWait++
                        utils.voteParams.push(message)
                    }
                    break
                case "onVerifierDisable":
                    console.log("取消验证者验证资格事件回调：", message.payload)
                    _this.$notify({
                        title: "取消验证者验证资格事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                            isVerifier: false
                        })
                    })
                    break
            }
        })
    },
    vote: function (_this, message) {
        console.log("验证者验证事件回调：", message.payload)
        _this.$notify({
            title: "验证者验证事件回调：",
            message: message.payload,
            position: "top-left"
        })
        if (message.payload.VerifierIndex === "0") {
            txVerifier_db.read(message.payload.TransactionID, function (txDetailsOnV) {
                txVerifier_db.write({
                    Title: txDetailsOnV.Title,
                    Price: txDetailsOnV.Price,
                    Keys: txDetailsOnV.Keys,
                    Description: txDetailsOnV.Description,
                    Buyer: txDetailsOnV.Buyer,
                    Seller: txDetailsOnV.Seller,
                    State: message.payload.TxState, // -
                    SupportVerify: txDetailsOnV.SupportVerify,
                    StartVerify: txDetailsOnV.StartVerify,
                    MetaDataExtension: txDetailsOnV.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnV.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnV.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnV.MetaDataIDEncWithBuyer,
                    MetaDataIDEncWithArbitrator: txDetailsOnV.MetaDataIDEncWithArbitrator,
                    Verifier1Response: message.payload.VerifierResponse, // -
                    Verifier2Response: txDetailsOnV.Verifier2Response,
                    ArbitrateResult: txDetailsOnV.ArbitrateResult,
                    PublishID: txDetailsOnV.PublishID,
                    TransactionID: txDetailsOnV.TransactionID
                }, function () {
                    txVerifier_db.init(_this)
                    utils.voteMutex = true
                    if (utils.voteWait > 0) {
                        utils.voteWait--
                        utils.vote(_this, utils.voteParams.shift())
                    }
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        })
                    })
                })
            })
        }
        if (message.payload.VerifierIndex === "1") {
            txBuyer_db.read(message.payload.TransactionID, function (txDetailsOnV) {
                txBuyer_db.write({
                    Title: txDetailsOnV.Title,
                    Price: txDetailsOnV.Price,
                    Keys: txDetailsOnV.Keys,
                    Description: txDetailsOnV.Description,
                    Buyer: txDetailsOnV.Buyer,
                    Seller: txDetailsOnV.Seller,
                    State: message.payload.TxState, // -
                    SupportVerify: txDetailsOnV.SupportVerify,
                    StartVerify: txDetailsOnV.StartVerify,
                    MetaDataExtension: txDetailsOnV.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnV.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnV.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnV.MetaDataIDEncWithBuyer,
                    MetaDataIDEncWithArbitrator: txDetailsOnV.MetaDataIDEncWithArbitrator,
                    Verifier1Response: message.payload.VerifierResponse, // -
                    Verifier2Response: txDetailsOnV.Verifier2Response,
                    ArbitrateResult: txDetailsOnV.ArbitrateResult,
                    PublishID: txDetailsOnV.PublishID,
                    TransactionID: txDetailsOnV.TransactionID
                }, function () {
                    txBuyer_db.init(_this)
                    utils.voteMutex = true
                    if (utils.voteWait > 0) {
                        utils.voteWait--
                        utils.vote(_this, utils.voteParams.shift())
                    }
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        })
                    })
                })
            })
        }
        if (message.payload.VerifierIndex === "2") {
            txBuyer_db.read(message.payload.TransactionID, function (txDetailsOnV) {
                txBuyer_db.write({
                    Title: txDetailsOnV.Title,
                    Price: txDetailsOnV.Price,
                    Keys: txDetailsOnV.Keys,
                    Description: txDetailsOnV.Description,
                    Buyer: txDetailsOnV.Buyer,
                    Seller: txDetailsOnV.Seller,
                    State: message.payload.TxState, // -
                    SupportVerify: txDetailsOnV.SupportVerify,
                    StartVerify: txDetailsOnV.StartVerify,
                    MetaDataExtension: txDetailsOnV.MetaDataExtension,
                    ProofDataExtensions: txDetailsOnV.ProofDataExtensions,
                    MetaDataIDEncWithSeller: txDetailsOnV.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: txDetailsOnV.MetaDataIDEncWithBuyer,
                    MetaDataIDEncWithArbitrator: txDetailsOnV.MetaDataIDEncWithArbitrator,
                    Verifier1Response: txDetailsOnV.Verifier1Response,
                    Verifier2Response: message.payload.VerifierResponse, // -
                    ArbitrateResult: txDetailsOnV.ArbitrateResult,
                    PublishID: txDetailsOnV.PublishID,
                    TransactionID: txDetailsOnV.TransactionID
                }, function () {
                    txBuyer_db.init(_this)
                    utils.voteMutex = true
                    if (utils.voteWait > 0) {
                        utils.voteWait--
                        utils.vote(_this, utils.voteParams.shift())
                    }
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                            isVerifier: accInstance.isVerifier
                        })
                    })
                })
            })
        }
    }
}

export { utils }
