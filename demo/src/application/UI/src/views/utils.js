let utils = {
    listen: function () {
        astilectron.onMessage(function(message) {
            switch (message.name !== "error") {
                case "welcome": console.log(message.payload); break
                case "sdkInit": console.log(message.name + ": " + message.payload); break
                case "sendMessage":
                    this.$notify({
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
