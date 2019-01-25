let dt = {
    init:function () {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        document.addEventListener('astilectron-ready', function() {
            // -
            dt.listen();
            document.getElementById("account").innerHTML = `<button onclick="dt.smtgo()">Account</button>`;
        });
    },
    smtgo:function () {
        astilectron.sendMessage({"name" : "hello",payload : "message from js"},function (message) {
            console.log("received " + message.payload);
        });
    },
    listen:function () {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    dt.about(message.payload);
                    return {payload: "payload"};
                    break;
            }
        });
    },
    about:function(html) {
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
};