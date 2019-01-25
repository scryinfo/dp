let dt = {
    init:function () {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        document.addEventListener('astilectron-ready', function() {
            // -
            dt.listen();
        });
    },
    smtgo:function () {
        astilectron.sendMessage({Name : "hello",Payload : "message from js"},function (message) {
            console.log("received " + message.payload);
        });
    },
    listen:function () {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                case "about2":
                    dt.about(message.payload);
                    return {payload: "payload"};
                    break;
                case "welcome":
                    asticode.notifier.info(message.payload);
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