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
                        PublishID: message.payload.PublishID
                    }, function () {
                        dl_db.init(_this)
                    })
                    acc_db.write({
                        address: _this.$store.state.account,
                        fromBlock: message.payload.Block + 1
                    })
                    break
                case "onApprove":
                    console.log("Node: onApprove.callback. ", message.payload)
                    _this.$notify({
                        title: "onApprove.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    acc_db.write({
                        address: _this.$store.state.account,
                        fromBlock: message.payload.Block + 1
                    })
                    break
                case "onTransactionCreatePrepare":
                    dl_db.read(message.payload, function (dlInstance) {
                        astilectron.sendMessage({ Name:"prepared", Payload: {extensions: dlInstance.ProofDataExtensions}})
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
                            MetaDataExtension: dataDetails.MetaDataExtension,
                            ProofDataExtensions: dataDetails.ProofDataExtensions,
                            MetaDataIDEncWithSeller: "",          // the two encrypt IDs will initialize later,
                            MetaDataIDEncWithBuyer: "",           // use (tx_db.read + tx_db.write) methods.
                            Verifier1Response: "Score, Comment",  // here three items will use a new method,
                            Verifier2Response: "Score, Comment",  // but both vote and arbitrate are not finished,
                            ArbitrateResult: false,               // so just still it and wait for function implement.
                            PublishID: dataDetails.PublishID,
                            TransactionID: message.payload.TransactionID    // keyPath
                        }, function () {
                            tx_db.init(_this)
                        })
                    })
                    acc_db.write({
                        address: _this.$store.state.account,
                        fromBlock: message.payload.Block + 1
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
                            MetaDataExtension: txDetailsOnPurchase.MetaDataExtension,
                            ProofDataExtensions: txDetailsOnPurchase.ProofDataExtensions,
                            MetaDataIDEncWithSeller: message.payload.MetaDataIdEncWithSeller, // -
                            MetaDataIDEncWithBuyer: txDetailsOnPurchase.MetaDataIDEncWithBuyer,
                            Verifier1Response: txDetailsOnPurchase.Verifier1Response,
                            Verifier2Response: txDetailsOnPurchase.Verifier2Response,
                            ArbitrateResult: txDetailsOnPurchase.ArbitrateResult,
                            PublishID: txDetailsOnPurchase.PublishID,
                            TransactionID: txDetailsOnPurchase.TransactionID // keyPath
                        },function () {
                            tx_db.init(_this)
                        })
                    })
                    acc_db.write({
                        address: _this.$store.state.account,
                        fromBlock: message.payload.Block + 1
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
                            MetaDataExtension: txDetailsOnRFD.MetaDataExtension,
                            ProofDataExtensions: txDetailsOnRFD.ProofDataExtensions,
                            MetaDataIDEncWithSeller: txDetailsOnRFD.MetaDataIDEncWithSeller,
                            MetaDataIDEncWithBuyer: message.payload.MetaDataIdEncWithBuyer, // -
                            Verifier1Response: txDetailsOnRFD.Verifier1Response,
                            Verifier2Response: txDetailsOnRFD.Verifier2Response,
                            ArbitrateResult: txDetailsOnRFD.ArbitrateResult,
                            PublishID: txDetailsOnRFD.PublishID,
                            TransactionID: txDetailsOnRFD.TransactionID
                        }, function () {
                            tx_db.init(_this)
                        })
                    })
                    acc_db.write({
                        address: _this.$store.state.account,
                        fromBlock: message.payload.Block + 1
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
                            MetaDataExtension: txDetailsOnC.MetaDataExtension,
                            ProofDataExtensions: txDetailsOnC.ProofDataExtensions,
                            MetaDataIDEncWithSeller: txDetailsOnC.MetaDataIDEncWithSeller,
                            MetaDataIDEncWithBuyer: txDetailsOnC.MetaDataIDEncWithBuyer,
                            Verifier1Response: txDetailsOnC.Verifier1Response,
                            Verifier2Response: txDetailsOnC.Verifier2Response,
                            ArbitrateResult: txDetailsOnC.ArbitrateResult,
                            PublishID: txDetailsOnC.PublishID,
                            TransactionID: txDetailsOnC.TransactionID
                        }, function () {
                            tx_db.init(_this)
                        })
                    })
                    acc_db.write({
                        address: _this.$store.state.account,
                        fromBlock: message.payload.Block + 1
                    })
                    break
            }
        })
    }
}

export { utils }
