<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <div v-if="isReload">
        <el-row>
            <el-col :span="24"><div class="top">Dapp</div></el-col>
        </el-row>
        <el-row>
            <el-col :span="8">
                <div class="left">
                    <div class="left-explain">选择账户：</div>
                    <el-select class="left-account" v-model="account" placeholder="账户" clearable allow-create filterable>
                        <el-option v-for="acc in this.$store.state.accounts" :key="acc.address" :value="acc.address" :label="acc.address"></el-option>
                    </el-select>
                    <div class="left-button-margin">
                        <el-button class="left-button" @click="right('登录')">登录</el-button>
                        <el-button class="left-button" @click="right('新建')">创建新用户</el-button>
                    </div>
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
import {connect} from "../utils/connect.js";
export default {
    name: "login.vue",
    data() {
        return {
            account: "",
            password: "",
            isReload: false,
            showControl1: false,
            showControl2: false,
            buttonControl: true,
            describe: ""
        }
    },
    methods: {
        right: function (description) {
            this.showControl1 = true;this.showControl2 = false;
            switch (description) {
                case "登录": this.buttonControl = true;break;
                case "新建": this.buttonControl = false;break;
            }
            this.describe = description + ":";
        },
        hide: function () {this.showControl1 = false; this.showControl2 = false; this.password = "";},
        submit_login: function () {
            let acc = this.account;
            this.account = "";
            let pwd = this.password;
            this.password = "";
            connect.send({Name: "loginVerify", Payload: {address: acc, password: pwd}}, function (payload, _this) {
                _this.$router.push({ name: "home", params: {acc: acc}});
            }, function (payload, _this) {
                console.log("登录验证失败：", payload);
                _this.$alert(payload, "用户名或密码错误！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        submit_new: function () {
            let pwd = this.password;
            this.password = "";
            let _login = this;
            connect.send({Name: "createNewAccount", Payload: {password: pwd}}, function (payload, _this) {
                _login.account = payload;
                _login.showControl1 = false;
                _login.showControl2 = true;
            }, function (payload, _this) {
                console.log("创建新账户失败：", payload);
                _this.$alert(payload, "创建新账户失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        submit_keystore: function () {
            this.password = "";
            this.$router.push({ name: "home", params: {acc: this.account}});
        }
    },
    created() {
        window.sessionStorage.clear();
        if (this.$store.state.account !== "") {
            return window.location.reload();
        }
        this.isReload = true;
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
.left-button-margin {
    margin-left: 17%
}
.left-button {
    width: 65%;
    padding: 5px 20px;
    margin: 5px 10px;
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
