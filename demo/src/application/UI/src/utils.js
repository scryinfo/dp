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
                case "onPublish":
                    console.log(message.payload)
                    _this.$notify({
                        title: "onPublish.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "onApprove":
                    console.log(message.payload)
                    _this.$notify({
                        title: "onApprove.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
                    break
                case "onTransactionCreat":
                    console.log(message.payload)
                    _this.$notify({
                        title: "onTransactionCreat.callback: ",
                        message: message.payload,
                        position: "top-left"
                    })
            }
        })
    }
}

export { utils }
