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
        asticode.loader.show();
        astilectron.sendMessage({"name": "Account", payload: "administrator"}, function(message) {
            // Init
            asticode.loader.hide();

            if (message.name === "error") {
                // Process error
            }
            // document.getElementById("account").innerHTML = message.payload;
            dt.about(message.payload);
        });
    },
    listen:function () {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    dt.about(message.payload);
                    return {payload: "payload"};
                    break;
                case "check.out.menu":
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