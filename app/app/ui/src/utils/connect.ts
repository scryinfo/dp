// import NodeJS from 'timers'
export class Connect {
    public ws: WebSocket | undefined;
    public count: number = 0;
    public MAX: number = 1000;
    public map: Map<string, any> = new Map<string, any>();
    public msgMutex: boolean = true;
    public msgParams: Array<any> = [];
    t : number = 0;

    WSConnect (_this?:any){

        // url: 'http://127.0.0.1:9822/#/'
        let port = 9822;

        this.ws = new WebSocket("ws://127.0.0.1:"+ port + "/ws", "ws");
        this.ws.onopen =  (evt:Event) => {
            console.log("connection onopen. ", evt);
            initAccs(this);
        };
        this.ws.onmessage = (evt:MessageEvent)=> {
            console.log("received   : ", evt.data);
            let obj = JSON.parse(evt.data);
            this.msgHandle(obj, _this);
        };
        this.ws.onclose =  (evt:CloseEvent) => {
            console.log("connection onclose. ", evt);
            // @ts-ignore
            this.ws.close();
            _this.$confirm("websocket连接已断开，请点击按钮重新连接。", "连接断开！", {
                confirmButtonText: "重新连接",
                cancelButtonText: "取消",
                type: "error"
            }).then(() => {
                // @ts-ignore
                this.reconnect(this.count,this.MAX,this.ws.readyState);
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
        this.ws.onerror =(evt:Event) => {
            console.log("connection onerror. ", evt);
            // @ts-ignore
            this.ws.close();
        };
    }

    addCallbackFunc (name:string, func:(payload:any, _this:any)=>void) {
        this.map.set(name,func);
    }

    cleanFuncMap () {
        this.map.clear();
    }


    send (obj:any, cbs:(payload:any, _this:any)=>void, cbf:(payload:any, _this:any)=>void) {
        if (!this.ws) { console.error(obj); return; }
        if (!!cbs) { this.addCallbackFunc(obj.Name + ".callback", cbs); }
        if (!!cbf) { this.addCallbackFunc(obj.Name + ".callback.error", cbf); }

        console.log("before send: ", JSON.stringify(obj));
        this.ws.send(JSON.stringify(obj));
    }

    reconnect(count: number, MAX: number, readyState: number) {
        count++;
        console.log("reconnection...【" + count + "】");
        // 1: has connected with server
        if (count >= MAX || readyState === 1) {
            clearTimeout(this.t);
        } else {
            // 3: has closed connection with server
            if (readyState === 3) {
                this.WSConnect();
            }
            // 0: trying connect to server, 2: closing connection with server
            this.t = window.setTimeout(() =>{this.reconnect(count, MAX, readyState);}, 200);
        }
    }

    async msgHandle (obj:any , _this:any)  {
        if (this.msgMutex) {
            this.msgMutex = false;
            console.log(obj);
            console.log(_this);
            await this.map.get(obj.Name)(obj.Payload,_this);
            await timeout(250);
            this.msgMutex = true;
            if (this.msgParams.length > 0) {
                let o = this.msgParams.shift();
                this.msgHandle(o, _this);
            }
        } else {
            this.msgParams.push(obj);
        }
    }

    constructor() {
    }
}


function timeout(ms:number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function initAccs (connect:Connect) {
    connect.send({Name: "getAccountsList", Payload: ""}, function (payload:any, _this:any) {
        _this.$store.state.accounts = [];
        for (let i = 0; i < payload.length; i++) {
            _this.$store.state.accounts.push({
                address: payload[i].Address
            })
        }
    }, function (payload: any, _this:any) {
        console.log(_this);
        console.log(payload);
        console.log("获取历史用户列表失败：", payload);
        _this.$alert(payload, "获取历史用户列表失败！", {
            confirmButtonText: "关闭",
            showClose: false,
            type: "error"
        });
    });
}

let connects = new Connect();

export default connects;

