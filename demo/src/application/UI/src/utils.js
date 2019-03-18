import {dl_db, tx_db} from "./DBoptions.js"
let utils = {
    listen: function (_this) {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "welcome": console.log(message.payload)
                    break
                case "sendMessage":
                    _this.$notify({
                        title: "Notify: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "initDL": dl_db.init(_this)
                    break
                case "initTx":tx_db.init(_this)
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
                        Seller: _this.$store.state.account,
                        SupportVerify: message.payload.SupportVerify,
                        MetaDataExtension: message.payload.MetaDataExtension,
                        ProofDataExtensions: message.payload.ProofDataExtensions,
                        PublishID: message.payload.PublishID
                    })
                    dl_db.init(_this)
                    break
                case "onApprove":
                    console.log("Node: onApprove.callback. ", message.payload)
                    _this.$notify({
                        title: "onApprove.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "onTransactionCreate":
                    console.log("Node: onTransactionCreate.callback. ", message.payload)
                    _this.$notify({
                        title: "onTransactionCreate.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    let dataDetails = dl_db.read(message.payload.PublishId)
                    tx_db.write({
                        Title: dataDetails.Title,
                        Price: dataDetails.Price,
                        Keys: dataDetails.Keys,
                        Description: dataDetails.Description,
                        Buyer: _this.$store.state.account,
                        Seller: dataDetails.Seller,
                        State: message.payload.TxState,
                        SupportVerify: dataDetails.SupportVerify,
                        MetaDataExtension: dataDetails.MetaDataExtension,
                        ProofDataExtensions: dataDetails.ProofDataExtensions,
                        MetaDataIDEncWithSeller: "",          // the two encrypt ID will initialize later,
                        MetaDataIDEncWithBuyer: "",           // use (tx_db.read + tx_db.write) methods.
                        Verifier1Response: "Score, Comment",  // here three items will use a new update function,
                        Verifier2Response: "Score, Comment",  // but both vote and arbitrate are not finished,
                        ArbitrateResult: false,               // so just still it and wait for function implement.
                        PublishID: dataDetails.PublishID,
                        TransactionID: message.payload.TransactionID
                    })
                    tx_db.init(_this)
                    break
                case "onPurchase":
                    console.log("Node: onPurchase.callback. ", message.payload)
                    _this.$notify({
                        title: "onPurchase.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    let txDetailsOnPurchase = tx_db.read(message.payload.TransactionID)
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
                        TransactionID: txDetailsOnPurchase.TransactionID
                    })
                    tx_db.init(_this)
                    break
                case "onReadyForDownload":
                    console.log("Node: onReadyForDownload.callback. ", message.payload)
                    _this.$notify({
                        title: "onReadyForDownload.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    let txDetailsOnRFD = tx_db.read(message.payload.TransactionID)
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
                    })
                    tx_db.init(_this)
                    break
                case "onClose":
                    console.log("Node: onClose.callback. ", message.payload)
                    _this.$notify({
                        title: "onClose.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
            }
        })
    }
}

export { utils }
