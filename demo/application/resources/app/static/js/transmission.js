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
    listen:function () {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    dt.about(message.payload);
                    return {payload: "payload"};
                    break;
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