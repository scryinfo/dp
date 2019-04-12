import {dl_db, tx_db, acc_db} from "./DBoptions.js"
let utils = {
    voteMutex: true,
    voteWait: 0,
    voteParams: [],
    listen: function (_this) {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "welcome":
                    _this.$notify({
                        title: "Notify: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "resetChain":
                    dl_db.reset()
                    tx_db.reset()
                    acc_db.reset()
                    console.log("重置完成")
                    break
                case "initDL":
                    dl_db.init(_this)
                    console.log("数据列表初始化完成")
                    break
                case "initTx":
                    tx_db.init(_this)
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
                case "onApprove":
                    console.log("允许合约转账事件回调：", message.payload)
                    _this.$notify({
                        title: "允许合约转账事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "onProofFilesExtensions":
                    dl_db.read(message.payload, function (dlInstance) {
                        astilectron.sendMessage({ Name:"extensions", Payload: {extensions: dlInstance.ProofDataExtensions}})
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
                        tx_db.write({
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
                            MetaDataIDEncWithArbitrator: "",
                            Verifier1: message.payload.Verifier1,   // -
                            Verifier2: message.payload.Verifier2,   // -
                            Verifier1Response: "",
                            Verifier2Response: "",
                            ArbitrateResult: false,
                            PublishID: dataDetails.PublishID,
                            TransactionID: message.payload.TransactionID    // keyPath
                        }, function () {
                            tx_db.init(_this)
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
                    tx_db.read(message.payload.TransactionID, function (txDetailsOnPurchase) {
                        tx_db.write({
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
                            MetaDataIDEncWithArbitrator: txDetailsOnPurchase.MetaDataIDEncWithArbitrator,
                            Verifier1: txDetailsOnPurchase.Verifier1,
                            Verifier2: txDetailsOnPurchase.Verifier2,
                            Verifier1Response: txDetailsOnPurchase.Verifier1Response,
                            Verifier2Response: txDetailsOnPurchase.Verifier2Response,
                            ArbitrateResult: txDetailsOnPurchase.ArbitrateResult,
                            PublishID: txDetailsOnPurchase.PublishID,
                            TransactionID: txDetailsOnPurchase.TransactionID // keyPath
                        },function () {
                            tx_db.init(_this)
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
                case "onReadyForDownload":
                    console.log("再加密数据事件回调：", message.payload)
                    _this.$notify({
                        title: "再加密数据事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    tx_db.read(message.payload.TransactionID, function (txDetailsOnRFD) {
                        tx_db.write({
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
                            MetaDataIDEncWithArbitrator: message.payload.MetaDataIDEncWithArbitrators, // - 
                            Verifier1: txDetailsOnRFD.Verifier1,
                            Verifier2: txDetailsOnRFD.Verifier2,
                            Verifier1Response: txDetailsOnRFD.Verifier1Response,
                            Verifier2Response: txDetailsOnRFD.Verifier2Response,
                            ArbitrateResult: txDetailsOnRFD.ArbitrateResult,
                            PublishID: txDetailsOnRFD.PublishID,
                            TransactionID: txDetailsOnRFD.TransactionID
                        }, function () {
                            tx_db.init(_this)
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
                case "onClose":
                    console.log("交易关闭事件回调：", message.payload)
                    _this.$notify({
                        title: "交易关闭事件回调：",
                        message: message.payload,
                        position: "top-left"
                    })
                    tx_db.read(message.payload.TransactionID, function (txDetailsOnC) {
                        tx_db.write({
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
                            MetaDataIDEncWithArbitrator: txDetailsOnC.MetaDataIDEncWithArbitrator,
                            Verifier1: txDetailsOnC.Verifier1,
                            Verifier2: txDetailsOnC.Verifier2,
                            Verifier1Response: txDetailsOnC.Verifier1Response,
                            Verifier2Response: txDetailsOnC.Verifier2Response,
                            ArbitrateResult: txDetailsOnC.ArbitrateResult,
                            PublishID: txDetailsOnC.PublishID,
                            TransactionID: txDetailsOnC.TransactionID
                        }, function () {
                            tx_db.init(_this)
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
        tx_db.read(message.payload.TransactionID, function (txDetailsOnV) {
            if (message.payload.VerifierIndex === "0") {
                tx_db.write({
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
                    Verifier1: txDetailsOnV.Verifier1,
                    Verifier2: txDetailsOnV.Verifier2,
                    Verifier1Response: message.payload.VerifierResponse, // -
                    Verifier2Response: txDetailsOnV.Verifier2Response,
                    ArbitrateResult: txDetailsOnV.ArbitrateResult,
                    PublishID: txDetailsOnV.PublishID,
                    TransactionID: txDetailsOnV.TransactionID
                }, function () {
                    tx_db.init(_this)
                    utils.voteMutex = true
                    if (utils.voteWait > 0) {
                        utils.voteWait--
                        utils.vote(_this, utils.voteParams.shift())
                    }
                })
            }
            if (message.payload.VerifierIndex === "1") {
                tx_db.write({
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
                    MetaDataIDEncWithArbitrator:txDetailsOnV.MetaDataIDEncWithArbitrator,
                    Verifier1: txDetailsOnV.Verifier1,
                    Verifier2: txDetailsOnV.Verifier2,
                    Verifier1Response: txDetailsOnV.Verifier1Response,
                    Verifier2Response: message.payload.VerifierResponse, // -
                    ArbitrateResult: txDetailsOnV.ArbitrateResult,
                    PublishID: txDetailsOnV.PublishID,
                    TransactionID: txDetailsOnV.TransactionID
                }, function () {
                    tx_db.init(_this)
                    utils.voteMutex = true
                    if (utils.voteWait > 0) {
                        utils.voteWait--
                        utils.vote(_this, utils.voteParams.shift())
                    }
                })
            }
        })
        acc_db.read(_this.$store.state.account, function (accInstance) {
            acc_db.write({
                address: accInstance.address,
                fromBlock: Math.max(accInstance.fromBlock, message.payload.Block + 1),
                isVerifier: accInstance.isVerifier
            })
        })
    }
}

export { utils }
