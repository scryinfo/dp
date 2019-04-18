<template>
    <div>
        <el-row>
            <el-col :span="24"><div class="top">App demo</div></el-col>
        </el-row>
        <el-row>
            <el-col :span="8">
                <div class="left">
                    <div class="left-explain">选择账户：</div>
                    <el-select class="left-account" v-model="account" placeholder="账户"
                        clearable allow-create filterable>
                        <el-option v-for="acc in this.$store.state.accounts" :key="acc.address"
                                   :value="acc.address" :label="acc.address"></el-option>
                    </el-select>
                    <el-button class="left-button" @click="right('登录')">登录</el-button>
                    <el-button class="left-button" @click="right('新建')">创建新用户</el-button>
                </div>
            </el-col>
            <el-col :span="16">
                <div class="right" id="show" v-if="showControl1">
                    <div class="right-show">{{describe}}
                        <el-input class="right-pwd" v-model="password" placeholder="密码"
                                  clearable show-password></el-input>
                    </div>
                    <el-button class="right-button" @click="hide">返回</el-button>
                    <el-button class="right-button" @click="submit_login" v-if="buttonControl">确认</el-button>
                    <el-button class="right-button" @click="submit_new" v-if="!buttonControl">确认</el-button>
                </div>
                <div class="right" id="show_new" v-if="showControl2">
                    你的新账户已创建完成： <br/>
                    {{account}}<br/>
                    如果需要在其他设备使用，请注意保存。<br/><br/><hr/><br/>
                    确定使用该账户登录吗？
                    <div class="right-pwd">
                        <el-button class="right-button" @click="hide">取消</el-button>
                        <el-button class="right-button" @click="submit_keystore">确认</el-button>
                    </div>
                </div>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import {db_init, acc_db} from "../DBoptions.js"
import {utils} from "../utils.js"
export default {
    name: "Login",
    data() {
        return {
            account: "",
            password: "",
            showControl1: false,
            showControl2: false,
            buttonControl: true,
            describe: ""
        }
    },
    methods: {
        right: function (description) {
            this.showControl1 = true;this.showControl2 = false
            switch (description) {
                case "登录": this.buttonControl = true;break
                case "新建": this.buttonControl = false;break
            }
            this.describe = description + ":"
        },
        hide: function () { this.showControl1 = false; this.showControl2 = false; this.password = "" },
        submit_login: function () {
            let pwd = this.password
            this.password = ""
            let _this = this
            astilectron.sendMessage({Name: "login.verify", Payload: {account: this.account,
                        password: pwd}}, function (message) {
                if (message.name !== "error") {
                    _this.$router.push({ name: "home", params: {acc: _this.account}})
                } else {
                    console.log("登录验证失败：", message)
                    _this.$alert(message.payload, "用户名或密码错误！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        },
        submit_new: function () {
            let _this = this
            astilectron.sendMessage({Name: "create.new.account", Payload: {password: this.password}}, function (message) {
                if (message.name !== "error") {
                    acc_db.write({
                        address: message.payload,
                        fromBlock: 1,
                        isVerifier: false
                    })
                    _this.account = message.payload
                    _this.showControl1 = false;_this.showControl2 = true
                }else {
                    console.log("创建新账户失败：", message)
                    _this.$alert(message.payload, "创建新账户失败！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        },
        submit_keystore: function () {
            this.password = ""
            this.$router.push({ name: "home", params: {acc: this.account}})
        }
    },
    created() {
        this.password = "";this.describe = "";this.account = ""
        let _this = this
        db_init.utilsDBInit()
        document.addEventListener("astilectron-ready", function() {
            utils.listen(_this)
        })
    }
}
</script>

<style scoped>
.top {
    background-color: grey;
    font-size: 20px;
    color: white;
    height: 50px;
    text-align: center;
    line-height: 50px;
}
.left {
    background-color: lightgray;
    height: calc(100vh - 50px);
}
.left-explain {
    text-align: left;
    padding: 100px 15% 0 15%;
}
.left-account {
    width: 70%;
    margin: 30px 15% 70px 15%;
}
.left-button {
    width: 65%;
    padding: 5px 20px;
    margin: 5px 17%;
}
.right {
    padding: 100px 20%;
    height: calc(100vh - 250px);
}
.right-show {
    width: 100%;
    text-align: left;
}
.right-pwd {
    width: 100%;
    margin: 30px 0 70px 0;
}
.right-button {
    width: 40%;
    margin: 0 2.5%;
}
</style>
