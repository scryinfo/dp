// Scry Info.  All rights reserved.
// license that can be found in the license file.

let connect = {
    ws: WebSocket,
    ipfs: require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'}),
    map: {},
    msgMutex: true,
    msgWait: 0,
    msgParams: [],
    WSConnect: function (_this) {
        // url: 'http://127.0.0.1:9822/#/'
        let port = window.location.href.split(":")[2].split("/")[0];
        connect.ws = new WebSocket("ws://127.0.0.1:"+ port + "/ws", "ws");
        connect.ws.onopen = function (evt) {
            console.log("connection onopen. ", evt);
        };
        connect.ws.onmessage = function (evt) {
            console.log(evt.data);
            let obj = JSON.parse(evt.data);
            connect.msgHandle(obj, _this);
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
    msgHandle: async function (obj, _this) {
        if (connect.msgMutex) {
            connect.msgMutex = false;
            await connect.map[obj.Name](obj.Payload, _this);
            connect.msgMutex = true;
            if (connect.msgWait > 0) {
                connect.msgHandle(connect.msgParams.shift(), _this);
                connect.msgWait--;
            }
        } else {
            connect.msgParams.push(obj);
            connect.msgWait++;
        }
    },
    send: function (obj, cbs, cbf) {
        if (!connect.ws) { return; }
        if (!!cbs) { connect.addCallbackFunc(obj.Name + ".callback", cbs); }
        if (!!cbf) { connect.addCallbackFunc(obj.Name + ".callback.error", cbf); }

        connect.ws.send(JSON.stringify(obj));
    },
    addCallbackFunc: function (name, func) {
        connect.map[name] = func;
    },
    cleanFuncMap: function () {
        connect.map = {};
    }
};

export { connect };
