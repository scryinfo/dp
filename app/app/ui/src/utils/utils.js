import {connect} from "./connect";

let utils = {
    stateEnum: ["Begin", "Created", "Voted", "Buying", "ReadyForDownload", "Closed"],

    state: [
        [false, false, false,  true, false, false], // 0 seller: re-encrypt
        [false,  true,  true,  true, false, false], // 1 buyer: cancel
        [false,  true,  true, false, false, false], // 2 buyer: confirmPurchase
        [false, false, false, false,  true, false], // 3 buyer: decrypt/confirmData
        [false, false,  true,  true,  true,  true]  // 4 buyer: grade
    ],

    init: function () {
        connect.addCallbackFunc("onPublish", presetFunc.onPublish);
        connect.addCallbackFunc("onVerifiersChosen", presetFunc.onVerifiersChosen);
        connect.addCallbackFunc("onAdvancePurchase", presetFunc.onAdvancePurchase);
        connect.addCallbackFunc("onConfirmPurchase", presetFunc.onConfirmPurchase);
        connect.addCallbackFunc("onReEncrypt", presetFunc.onReEncrypt);
        connect.addCallbackFunc("onTransactionClose", presetFunc.onTransactionClose);
        connect.addCallbackFunc("onRegisterVerifier", presetFunc.onRegisterVerifier);
        connect.addCallbackFunc("onVoteResult", presetFunc.onVoteResult);
        connect.addCallbackFunc("onVerifierDisable", presetFunc.onVerifierDisable);
        connect.addCallbackFunc("onArbitrationBegin", presetFunc.onArbitrationBegin);
        connect.addCallbackFunc("onArbitrationResult", presetFunc.onArbitrationResult);
    },

    reacquireData: function (mode) {
        switch (mode) {
            case "dl":
                initFunc.initDL();
                break;
            case "txs":
                initFunc.initTxS();
                break;
            case "txb":
                initFunc.initTxB();
                break;
            case "txv":
                initFunc.initTxV();
                break;
            case "txa":
                initFunc.initTxA();
                break;
            case "all":
                initFunc.initAll();
                break;
            default:
                console.log("Invalid mode! (in data init)");
        }
    },

    setStateString: function (str) {
        let index;
        for (let i = 0; i < utils.stateEnum.length; i++) {
            if (str === utils.stateEnum[i]) {
                index = i;
                break;
            }
        }
        return index;
    },

    functionDisabled: function (funcNum, stateStr) {
        return !utils.state[funcNum][utils.setStateString(stateStr)];
    },

    setSupportVerify: function (sv) {
        let result = "";
        if (sv) {
            result = "支持验证";
        } else {
            result = "不支持验证";
        }

        return result;
    },

    setNeedVerify: function (nv) {
        let result = "";
        if (nv) {
            result = "已启用验证";
        } else {
            result = "未启用验证";
        }

        return result;
    }
};

let presetFunc = {
    onPublish: function (payload, _this) {
        console.log("发布事件回调：", payload);
        _this.$notify({
            title: "发布事件回调：",
            message: "数据ID：" + payload.PublishId,
            position: "top-left"
        });
        _this.$store.state.datalist.push({
            Title: payload.Title,
            Price: payload.Price,
            Keys: payload.Keys,
            Description: payload.Description,
            Seller: payload.Seller,
            SupportVerify: payload.SupportVerify,
            SVDisplay: utils.setSupportVerify(payload.SupportVerify),
            PublishId: payload.PublishId
        });
    },

    onVerifiersChosen: function (payload, _this) {
        console.log("选择验证者事件回调：", payload);
        _this.$notify({
            title: "选择验证者事件回调：",
            message: "你已被选中成为 ID: " + payload.TransactionId + " 交易 的验证者。",
            position: "top-left"
        });
        _this.$store.state.transactionverifier.push({
            PublishId: payload.PublishId,
            TransactionId: payload.TransactionId,
            Title: payload.Title,
            Price: payload.Price,
            Keys: payload.Keys,
            Description: payload.Description
        });
    },

    onAdvancePurchase: function (payload, _this) {
        console.log("创建交易事件回调：", payload);
        _this.$notify({
            title: "创建交易事件回调：",
            message: "创建交易(ID：" + payload.TransactionId + ")成功",
            position: "top-left"
        });
        let param = {
            PublishId: payload.PublishId,
            TransactionId: payload.TransactionId,
            Title: payload.Title,
            Price: payload.Price,
            Keys: payload.Keys,
            Description: payload.Description,
            State: utils.stateEnum[parseInt(payload.State)],
            SVDisplay: utils.setSupportVerify(payload.SupportVerify),
            NVDisplay: utils.setNeedVerify(payload.StartVerify),
            ArbitrateResult: payload.ArbitrateResult
        };
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.push(param);
                break;
            case 2:
                param.SupportVerify = payload.SupportVerify;
                param.Verifier1Response = payload.Verifier1Response;
                param.Verifier2Response = payload.Verifier2Response;
                _this.$store.state.transactionbuy.push(param);
                break;
            default:
                console.log("Invalid identify! (on advance purchase)");
                break;
        }
    },

    onConfirmPurchase: function (payload, _this) {
        console.log("购买数据事件回调：", payload);
        _this.$notify({
            title: "购买数据事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已确认购买。",
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.forEach(function (item) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 3:
                _this.$store.state.transactionverifier.forEach(function (item, index, arr) {
                    if (item.TransactionId === payload.TransactionId) {
                        // delete item
                        arr[index] = arr[0];
                        arr.shift();
                    }
                });
                break;
        }
    },

    onReEncrypt: function (payload, _this) {
        console.log("再加密数据事件回调：", payload);
        _this.$notify({
            title: "再加密数据事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已重新加密。",
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.forEach(function (item) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
        }
    },

    onTransactionClose: function (payload, _this) {
        console.log("交易关闭事件回调：", payload);
        _this.$notify({
            title: "交易关闭事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已关闭。",
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.forEach(function (item, index, arr) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item, index, arr) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 3:
                _this.$store.state.transactionverifier.forEach(function (item, index, arr) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
        }
    },

    onRegisterVerifier: function (payload, _this) {
        console.log("注册成为验证者事件回调。");
        _this.$notify({
            title: "注册成为验证者事件回调：",
            message: "你已成功注册成为验证者！ :)",
            position: "top-left"
        });
    },

    onVoteResult: function (payload, _this) {
        console.log("验证者验证事件回调：", payload);
        _this.$notify({
            title: "验证者验证事件回调：",
            message: "收到新的验证者回复，交易ID：" + payload.TransactionId,
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 3:
                _this.$store.state.transactionverifier.forEach(function (item) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
        }
    },

    onVerifierDisable: function (payload, _this) {
        console.log("取消验证者验证资格事件回调：", payload);
        _this.$notify({
            title: "取消验证者验证资格事件回调：",
            message: "验证者： " + payload.Address + "被取消验证资格。",
            position: "top-left"
        });
    },

    onArbitrationBegin: function (payload, _this) {
        console.log("仲裁开始事件回调：", payload);
        _this.$notify({
            title: "仲裁开始事件回调：",
            message: "你已被选中成为 ID: " + payload.TransactionId + " 交易 的仲裁者。",
            position: "top-left"
        });
        _this.$store.state.transactionarbitrator.push({
            PublishId: payload.PublishId,
            TransactionId: payload.TransactionId,
            Title: payload.Title,
            Price: payload.Price,
            Keys: payload.Keys,
            Description: payload.Description
        });
    },

    onArbitrationResult: function (payload, _this) {
        console.log("仲裁结果事件回调：", payload);
        _this.$notify({
            title: "仲裁结果事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已完成仲裁，仲裁结果为：" + payload.ArbitrateResult + " 。",
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.forEach(function (item) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.ArbitrateResult = payload.ArbitrateResult
                    }
                });
                break;
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.ArbitrateResult = payload.ArbitrateResult
                    }
                });
                break;
        }
    }
};

let initFunc = {
    initDL: function () {
        connect.send({Name: "getDataList", Payload: ""}, function (payload, _this) {
            _this.$store.state.datalist = [];
            for (let i = 0; i < payload.length; i++) {
                _this.$store.state.datalist.push({
                    Title: payload[i].Title,
                    Price: payload[i].Price,
                    Keys: payload[i].Keys,
                    Description: payload[i].Description,
                    Seller: payload[i].Seller,
                    SupportVerify: payload[i].SupportVerify,
                    PublishId: payload[i].PublishId,
                    SVDisplay: utils.setSupportVerify(payload[i].SupportVerify)
                })
            }
        }, function (payload, _this) {
            console.log("获取数据列表失败：", payload);
            _this.$alert(payload, "获取数据列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initTxS: function () {
        connect.send({Name: "getTxSell", Payload: ""}, function (payload, _this) {
            _this.$store.state.transactionsell = [];
            for (let i = 0; i < payload.length; i++) {
                _this.$store.state.transactionsell.push({
                    PublishId: payload[i].PublishId,
                    TransactionId: payload[i].TransactionId,
                    State: utils.stateEnum[parseInt(payload[i].State)],
                    Title: payload[i].Title,
                    Price: payload[i].Price,
                    Keys: payload[i].Keys,
                    Description: payload[i].Description,
                    ArbitrateResult: false,
                    SVDisplay: utils.setSupportVerify(payload[i].SupportVerify),
                    NVDisplay: utils.setNeedVerify(payload[i].StartVerify)
                })
            }
        }, function (payload, _this) {
            console.log("获取当前用户为卖方的交易列表失败：", payload);
            _this.$alert(payload, "获取当前用户为卖方的交易列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initTxB: function () {
        connect.send({Name: "getTxBuy", Payload: ""}, function (payload, _this) {
            _this.$store.state.transactionbuy = [];
            for (let i = 0; i < payload.length; i++) {
                _this.$store.state.transactionbuy.push({
                    PublishId: payload[i].PublishId,
                    TransactionId: payload[i].TransactionId,
                    State: utils.stateEnum[parseInt(payload[i].State)],
                    Title: payload[i].Title,
                    Price: payload[i].Price,
                    Keys: payload[i].Keys,
                    Description: payload[i].Description,
                    Verifier1Response: payload[i].Verifier1Response,
                    Verifier2Response: payload[i].Verifier2Response,
                    SupportVerify: payload[i].SupportVerify,
                    ArbitrateResult: payload[i].Description,
                    SVDisplay: utils.setSupportVerify(payload[i].SupportVerify),
                    NVDisplay: utils.setNeedVerify(payload[i].StartVerify),
                })
            }
        }, function (payload, _this) {
            console.log("获取当前用户为买方的交易列表失败：", payload);
            _this.$alert(payload, "获取当前用户为买方的交易列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initTxV: function () {
        connect.send({Name: "getTxVerify", Payload: ""}, function (payload, _this) {
            _this.$store.state.transactionverifier = [];
            if (payload.length > 0) {
                for (let i = 0; i < payload.length; i++) {
                    _this.$store.state.transactionverifier.push({
                        PublishId: payload[i].PublishId,
                        TransactionId: payload[i].TransactionId,
                        Title: payload[i].Title,
                        Price: payload[i].Price,
                        Keys: payload[i].Keys,
                        Description: payload[i].Description,
                    })
                }
            }
        }, function (payload, _this) {
            console.log("获取当前用户为验证者的交易列表失败：", payload);
            _this.$alert(payload, "获取当前用户为验证者的交易列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initTxA: function () {
        connect.send({Name: "getTxArbitrate", Payload: ""}, function (payload, _this) {
            _this.$store.state.transactionarbitrator = [];
            if (payload.length > 0) {
                for (let i = 0; i < payload.length; i++) {
                    _this.$store.state.transactionarbitrator.push({
                        PublishId: payload[i].PublishId,
                        TransactionId: payload[i].TransactionId,
                        Title: payload[i].Title,
                        Price: payload[i].Price,
                        Keys: payload[i].Keys,
                        Description: payload[i].Description,
                    })
                }
            }
        }, function (payload, _this) {
            console.log("获取当前用户为仲裁者的交易列表失败：", payload);
            _this.$alert(payload, "获取当前用户为仲裁者的交易列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initAll: function () {
        initFunc.initDL();
        initFunc.initTxS();
        initFunc.initTxB();
        initFunc.initTxV();
        initFunc.initTxA();
    }
};

export {utils}