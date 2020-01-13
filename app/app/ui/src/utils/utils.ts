import connects from "./connect";

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

        connects.addCallbackFunc("onPublish", presetFunc.onPublish);
        connects.addCallbackFunc("onVerifiersChosen", presetFunc.onVerifiersChosen);
        connects.addCallbackFunc("onAdvancePurchase", presetFunc.onAdvancePurchase);
        connects.addCallbackFunc("onConfirmPurchase", presetFunc.onConfirmPurchase);
        connects.addCallbackFunc("onReEncrypt", presetFunc.onReEncrypt);
        connects.addCallbackFunc("onTransactionClose", presetFunc.onTransactionClose);
        connects.addCallbackFunc("onRegisterVerifier", presetFunc.onRegisterVerifier);
        connects.addCallbackFunc("onVoteResult", presetFunc.onVoteResult);
        connects.addCallbackFunc("onVerifierDisable", presetFunc.onVerifierDisable);
        connects.addCallbackFunc("onArbitrationBegin", presetFunc.onArbitrationBegin);
        connects.addCallbackFunc("onArbitrationResult", presetFunc.onArbitrationResult);
    },

    reacquireData: function (mode: any, str?: any) {
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

            case "evtdl":
                  initFunc.initEvtDL(str);
                  break;
              default:
                  console.log("Invalid mode! (in data init)");
        }
    },

    setStateString: function (str: string) :number {
        let index = 0;
        for (let i = 0; i < utils.stateEnum.length; i++) {
            if (str === utils.stateEnum[i]) {
                index = i;
                break;
            }
        }
        return index;
    },

    functionDisabled: function (funcNum:number, stateStr: string) {
      return !utils.state[funcNum][utils.setStateString(stateStr)];
    },

    setSupportVerify: function (sv: any) {
        let result = "";
        if (sv) {
            result = "支持验证";
        } else {
            result = "不支持验证";
        }

        return result;
    },

    setNeedVerify: function (nv: any) {
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

    onPublish: function (payload:any, _this:any) {
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


    onVerifiersChosen: function (payload:any, _this:any) {
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


    onAdvancePurchase: function (payload:any, _this:any) {
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
            ArbitrateResult: payload.ArbitrateResult,
            SupportVerify:false,
            Verifier1Response:"",
            Verifier2Response:"",
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


    onConfirmPurchase: function (payload:any, _this:any) {
        console.log("购买数据事件回调：", payload);
        _this.$notify({
            title: "购买数据事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已确认购买。",
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.forEach(function (item: { TransactionId: any; State: string; }) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item: { TransactionId: any; State: string; }) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 3:
                _this.$store.state.transactionverifier.forEach(function (item: { TransactionId: any; }, index: number, arr: any[]) {
                    if (item.TransactionId === payload.TransactionId) {
                        // delete item

                      arr[index] = arr[0];
                        arr.shift();
                    }
                });
                break;
        }
    },

    onReEncrypt: function (payload:any, _this:any) {
        console.log("再加密数据事件回调：", payload);
        _this.$notify({
            title: "再加密数据事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已重新加密。",
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.forEach(function (item: { TransactionId: any; State: string; }) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item: { TransactionId: any; State: string; }) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
        }
    },

    onTransactionClose: function (payload:any, _this:any) {
        console.log("交易关闭事件回调：", payload);
        _this.$notify({
            title: "交易关闭事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已关闭。",
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.forEach(function (item: { TransactionId: any; State: string; }, index: any, arr: any) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item: { TransactionId: any; State: string; }, index: any, arr: any) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
            case 3:
                _this.$store.state.transactionverifier.forEach(function (item: { TransactionId: any; State: string; }, index: any, arr: any) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.State = utils.stateEnum[parseInt(payload.State)]
                    }
                });
                break;
        }
    },

    onRegisterVerifier: function (payload:any, _this:any) {
        console.log("注册成为验证者事件回调。");
        _this.$notify({
            title: "注册成为验证者事件回调：",
            message: "你已成功注册成为验证者！ :)",
            position: "top-left"
        });
    },

    onVoteResult: function (payload:any, _this:any) {
        console.log("验证者验证事件回调：", payload);
        _this.$notify({
            title: "验证者验证事件回调：",
            message: "收到新的验证者回复，交易ID：" + payload.TransactionId,
            position: "top-left"
        });
        if (parseInt(payload.Identify) === 2) {
            _this.$store.state.transactionbuy.forEach(function (item: { TransactionId: any; State: string; Verifier1Response: any; Verifier2Response: any; }) {
                if (item.TransactionId === payload.TransactionId) {
                    item.State = utils.stateEnum[parseInt(payload.State)]
                    if (payload.Verifier1Response !== "") {
                        item.Verifier1Response = payload.Verifier1Response
                    }
                    if (payload.Verifier2Response !== "") {
                        item.Verifier2Response = payload.Verifier2Response
                    }
                }
            });
        }
    },

    onVerifierDisable: function (payload:any, _this:any) {
        console.log("取消验证者验证资格事件回调：", payload);
        _this.$notify({
            title: "取消验证者验证资格事件回调：",
            message: "验证者： " + payload.Address + "被取消验证资格。",
            position: "top-left"
        });
    },

    onArbitrationBegin: function (payload:any, _this:any) {
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

    onArbitrationResult: function (payload:any, _this:any) {
        console.log("仲裁结果事件回调：", payload);
        _this.$notify({
            title: "仲裁结果事件回调：",
            message: "交易(ID: " + payload.TransactionId + ")已完成仲裁，仲裁结果为：" + payload.ArbitrateResult + " 。",
            position: "top-left"
        });
        switch (parseInt(payload.Identify)) {
            case 1:
                _this.$store.state.transactionsell.forEach(function (item: { TransactionId: any; ArbitrateResult: any; }) {
                    if (item.TransactionId === payload.TransactionId) {
                        item.ArbitrateResult = payload.ArbitrateResult
                    }
                });
                break;
            case 2:
                _this.$store.state.transactionbuy.forEach(function (item: { TransactionId: any; ArbitrateResult: any; }) {
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
        connects.send({Name: "getDataList", Payload: ""}, function (payload:any, _this:any) {
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

        }, function (payload:any, _this:any) {
            console.log("获取数据列表失败：", payload);
            _this.$alert(payload, "获取数据列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initTxS: function () {

        connects.send({Name: "getTxSell", Payload: ""}, function (payload:any, _this:any) {
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

        }, function (payload:any, _this:any) {
            console.log("获取当前用户为卖方的交易列表失败：", payload);
            _this.$alert(payload, "获取当前用户为卖方的交易列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initTxB: function () {

        connects.send({Name: "getTxBuy", Payload: ""}, function (payload:any, _this:any) {
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

        }, function (payload:any, _this:any) {
            console.log("获取当前用户为买方的交易列表失败：", payload);
            _this.$alert(payload, "获取当前用户为买方的交易列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initTxV: function () {

        connects.send({Name: "getTxVerify", Payload: ""}, function (payload:any, _this:any) {
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

        }, function (payload:any, _this:any) {
            console.log("获取当前用户为验证者的交易列表失败：", payload);
            _this.$alert(payload, "获取当前用户为验证者的交易列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initTxA: function () {

        connects.send({Name: "getTxArbitrate", Payload: ""}, function (payload:any, _this:any) {
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

        }, function (payload:any, _this:any) {
            console.log("获取当前用户为仲裁者的交易列表失败：", payload);
            _this.$alert(payload, "获取当前用户为仲裁者的交易列表失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    },

    initEvtDL: function (str: any) {

      connects.send({Name: "getEvtList", Payload: {eventBodys: str}}, function (payload:any, _this:any) {
      _this.$store.state.datalist = [];
      console.log("getEvtList:",payload);
      for (let i = 0; i < payload.length; i++) {
        _this.$store.state.datalist.push({
          ID:payload[i].Id,
          NotifyTo: payload[i].NotifyTo,
          EventName: payload[i].EventName,
          Keys: payload[i].Keys,
          EventBodys: payload[i].EventBodys,
          EventStatus: payload[i].EventStatus,
          CreatedTime: payload[i].CreatedTime,
        })
      }


      }, function (payload:any, _this:any) {
      console.log("获取数据列表失败：", payload);
      _this.$alert(payload, "获取数据列表失败！", {
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
