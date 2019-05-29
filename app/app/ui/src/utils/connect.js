// Scry Info.  All rights reserved.
// license that can be found in the license file.

let connect = {
    ws: WebSocket,
    // todo: ipfs node config
    ipfs: require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'}),
    map: {},
    WSConnect: function (_this) {
        let port = window.location.href.split(":")[2].split("/")[0];
        connect.ws = new WebSocket("ws://127.0.0.1:"+ port + "/ws", "ws");
        connect.ws.onopen = function (evt) {
            console.log("connection onopen. ", evt);
        };
        connect.ws.onmessage = function (evt) {
            console.log(evt.data);
            let obj = JSON.parse(evt.data);
            connect.map[obj.Name](obj.Payload, _this);
        };
        connect.ws.onclose = function (evt) {
            console.log("connection onclose. ", evt);
            connect.ws.close();
        };
        connect.ws.onerror = function (evt) {
            console.log("connection onerror. ", evt);
            connect.ws.close();
        };
        window.onbeforeunload = function () {
            connect.ws.close();
        };
        connect.WSConnect = function () {};
    },
    send: function (obj, cbs, cbf) {
        if (!connect.ws) {
            return;
        }
        connect.addCallbackFunc(obj.Name + ".callback", cbs);
        connect.addCallbackFunc(obj.Name + ".callback.error", cbf);
        connect.ws.send(JSON.stringify(obj));
    },
    addCallbackFunc: function (name, func) {
        connect.map[name] = func;
    },
    cleanMap: function () {
        connect.map = {};
    }
};

export { connect };
