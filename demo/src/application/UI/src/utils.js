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
            }
        })
    }
}

export { utils }
