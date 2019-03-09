import {dl_db, mt_db} from "./DBoptions.js"
let utils = {
    listen: function (_this) {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "welcome": console.log(message.payload); break
                case "sdkInit": console.log(message.name + ": " + message.payload); break
                case "sendMessage":
                    _this.$notify({
                        title: "Notify: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "initDL": dl_db.init(_this); break
                case "initMT": mt_db.init(_this); break
                case "onPublish":
                    console.log("Node: onPublish.callback. ", message.payload)
                    _this.$notify({
                        title: "onPublish.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "onApprove":
                    console.log("Node: onApprove.callback. ", message.payload)
                    _this.$notify({
                        title: "onApprove.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "onTransactionCreat":
                    console.log("Node: onTransactionCreat.callback. ", message.payload)
                    _this.$notify({
                        title: "onTransactionCreat.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    // go send the whole callback.event to js now, here will adjust later. core param is tID.
                    mt_db.write({
                        Title: "test title",
                        Price: 0,
                        Seller: "0x0000",
                        Buyer: "0x0001",
                        State: 0,
                        Verifier1Response: "3,v1r",
                        Verifier2Response: "3,v2r",
                        Verifier3Response: "3,v3r",
                        ArbitrateResult: false
                    }, message.payload.Data.transactionId)
                    break
                case "onPurchase":
                    console.log("Node: onPurchase.callback. ", message.payload)
                    _this.$notify({
                        title: "onPurchase.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "onReadyForDownload":
                    console.log("Node: onReadyForDownload.callback. ", message.payload)
                    _this.$notify({
                        title: "onReadyForDownload.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
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
