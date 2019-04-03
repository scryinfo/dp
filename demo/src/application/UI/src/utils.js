import {dl_db, tx_db, acc_db} from "./DBoptions.js"
let utils = {
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
                    console.log("Reset command received. ")
                    break
                case "initDL":
                    dl_db.init(_this)
                    console.log("dl_db init.")
                    break
                case "initTx":
                    tx_db.init(_this)
                    console.log("tx_db init.")
                    break
                case "onPublish":
                    console.log("Node: onPublish.callback. ", message.payload)
                    _this.$notify({
                        title: "onPublish.callback: ",
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
                            fromBlock: message.payload.Block + 1,
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onApprove":
                    console.log("Node: onApprove.callback. ", message.payload)
                    _this.$notify({
                        title: "onApprove.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: message.payload.Block + 1,
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onProofFilesExtensions":
                    dl_db.read(message.payload, function (dlInstance) {
                        astilectron.sendMessage({ Name:"extensions", Payload: {extensions: dlInstance.ProofDataExtensions}})
                    })
                    break
                case "onTransactionCreate":
                    console.log("Node: onTransactionCreate.callback. ", message.payload)
                    _this.$notify({
                        title: "onTransactionCreate.callback: ",
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
                            fromBlock: message.payload.Block + 1,
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onPurchase":
                    console.log("Node: onPurchase.callback. ", message.payload)
                    _this.$notify({
                        title: "onPurchase.callback: ",
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
                            fromBlock: message.payload.Block + 1,
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onReadyForDownload":
                    console.log("Node: onReadyForDownload.callback. ", message.payload)
                    _this.$notify({
                        title: "onReadyForDownload.callback: ",
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
                            fromBlock: message.payload.Block + 1,
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onClose":
                    console.log("Node: onClose.callback. ", message.payload)
                    _this.$notify({
                        title: "onClose.callback: ",
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
                            fromBlock: message.payload.Block + 1,
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onRegisterVerifier":
                    console.log("Node: onRegisterVerifier.callback. ", message.payload)
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: message.payload.Block + 1,
                            isVerifier: true
                        })
                    })
                    break
                case "onVote":
                    console.log("Node: onVote.callback. ", message.payload)
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
                                Verifier1: txDetailsOnV.Verifier1,
                                Verifier2: txDetailsOnV.Verifier2,
                                Verifier1Response: message.payload.VerifierResponse, // -
                                Verifier2Response: txDetailsOnV.Verifier2Response,
                                ArbitrateResult: txDetailsOnV.ArbitrateResult,
                                PublishID: txDetailsOnV.PublishID,
                                TransactionID: txDetailsOnV.TransactionID
                            }, function () {
                                tx_db.init(_this)
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
                                Verifier1: txDetailsOnV.Verifier1,
                                Verifier2: txDetailsOnV.Verifier2,
                                Verifier1Response: txDetailsOnV.Verifier1Response,
                                Verifier2Response: message.payload.VerifierResponse, // -
                                ArbitrateResult: txDetailsOnV.ArbitrateResult,
                                PublishID: txDetailsOnV.PublishID,
                                TransactionID: txDetailsOnV.TransactionID
                            }, function () {
                                tx_db.init(_this)
                            })
                        }
                    })
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: message.payload.Block + 1,
                            isVerifier: accInstance.isVerifier
                        })
                    })
                    break
                case "onVerifierDisable":
                    console.log("Node: onVerifierDisable.callback. ", message.payload)
                    acc_db.read(_this.$store.state.account, function (accInstance) {
                        acc_db.write({
                            address: accInstance.address,
                            fromBlock: message.payload.Block + 1,
                            isVerifier: false
                        })
                    })
                    break
            }
        })
    }
}

export { utils }
