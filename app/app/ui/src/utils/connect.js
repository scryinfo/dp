// Scry Info.  All rights reserved.
// license that can be found in the license file.

let connect = {
    ws: WebSocket,
    count: 0,
    MAX: 1000,
    t: setTimeout(function() {connect.reconnect();}, 200),

    ipfs: require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'}),

    map: {},
    msgMutex: true,
    msgParams: [],

    crypto: require('crypto'),

    WSConnect: function (_this) {
        // url: 'http://127.0.0.1:9822/#/'
        let port = window.location.href.split(":")[2].split("/")[0];

        connect.ws = new WebSocket("ws://127.0.0.1:"+ port + "/ws", "ws");
        connect.ws.onopen = function (evt) {
            console.log("connection onopen. ", evt);
            initAccs();
        };
        connect.ws.onmessage = function (evt) {
            console.log("received   : ", evt.data);
            let obj = JSON.parse(evt.data);
            connect.msgHandle(obj, _this);
        };
        connect.ws.onclose = function (evt) {
            console.log("connection onclose. ", evt);
            connect.ws.close();
            _this.$confirm("websocket连接已断开，请点击按钮重新连接。", "连接断开！", {
                confirmButtonText: "重新连接",
                cancelButtonText: "取消",
                type: "error"
            }).then(() => {
                connect.reconnect();
            }).catch(() => {
                _this.$message({
                    type:"error",
                    message:"websocket连接已断开。",
                    duration: 0,
                    showClose: true
                });
            });
            _this.$router.push("/");
        };
        connect.ws.onerror = function (evt) {
            console.log("connection onerror. ", evt);
            connect.ws.close();
        };
    },

    msgHandle: async function (obj, _this) {
        if (connect.msgMutex) {
            connect.msgMutex = false;
            await connect.map[obj.Name](obj.Payload, _this);
            await timeout(250);
            connect.msgMutex = true;
            if (connect.msgParams.length > 0) {
                connect.msgHandle(connect.msgParams.shift(), _this);
            }
        } else {
            connect.msgParams.push(obj);
        }
    },

    reconnect: function () {
        connect.count++;
        console.log("reconnection...【" + connect.count + "】");
        // 1: has connected with server
        if (connect.count >= connect.MAX || connect.ws.readyState === 1) {
            clearTimeout(connect.t);
        } else {
            // 3: has closed connection with server
            if (connect.ws.readyState === 3) {
                connect.WSConnect();
            }
            // 0: trying connect to server, 2: closing connection with server
            connect.t = setTimeout(function() {connect.reconnect();}, 200);
        }
    },

    send: function (obj, cbs, cbf) {
        if (!connect.ws) { return; }
        if (!!cbs) { connect.addCallbackFunc(obj.Name + ".callback", cbs); }
        if (!!cbf) { connect.addCallbackFunc(obj.Name + ".callback.error", cbf); }

        if (!!obj.Payload && !!obj.Payload.password) { // password hash
            obj.Payload.password = connect.calcHash(obj.Payload.password)
        }

        console.log("before send: ", JSON.stringify(obj));

        connect.ws.send(JSON.stringify(obj));
    },

    addCallbackFunc: function (name, func) {
        connect.map[name] = func;
    },

    cleanFuncMap: function () {
        connect.map = {};
    },

    calcHash: function (text) {
        return connect.crypto.createHash('sha256').update(text).digest('hex');
    }
};

function timeout(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function initAccs() {
    connect.send({Name: "getAccountsList", Payload: ""}, function (payload, _this) {
        _this.$store.state.accounts = [];
        for (let i = 0; i < payload.length; i++) {
            _this.$store.state.accounts.push({
                address: payload[i].Address
            })
        }
    }, function (payload, _this) {
        console.log("获取历史用户列表失败：", payload);
        _this.$alert(payload, "获取历史用户列表失败！", {
            confirmButtonText: "关闭",
            showClose: false,
            type: "error"
        });
    });
}

export { connect };
