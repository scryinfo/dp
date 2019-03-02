let lfg = {
    listen:function () {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    console.log(message.payload)
                    return {payload: "payload"}
                    break
                case "about2":
                    console.log(message.payload)
                    return {payload: "payload"}
                    break
                case "welcome":
                    console.log(message.payload)
                    break
                case "sdkInit":
                    console.log(message.name+": "+message.payload)
                    break
            }
        })
    }
}

export {lfg}
