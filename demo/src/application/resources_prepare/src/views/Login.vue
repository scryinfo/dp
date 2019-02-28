<template>
    <div>
        <el-row class="row">
            <el-col :span="24"><div class="top">My Astilectron demo</div></el-col>
            <el-col :span="8">
                <div class="left">
                    <div class="left-explain">select account:</div>
                    <el-select class="left-account" v-model="account" placeholder="select account">
                        <el-option v-for="acc in this.$store.state.accounts" :key="acc.address"
                                   :value="acc.address" :label="acc.address"></el-option>
                    </el-select>
                    <div><button class="left-button" @click="right('Login')">Login</button></div>
                    <div><button class="left-button" @click="right('New')">Create New Account</button></div>
                </div>
            </el-col>
            <el-col :span="16">
                <div class="right" id="show" v-if="showControl1">
                    <div class="right-show">{{describe}}
                        <el-input class="right-pwd" v-model="password"
                                  placeholder="password" type="password" :clearable="true"></el-input>
                    </div>
                    <div><button class="right-button" @click="hide">Back</button>
                    <button class="right-button" @click="submit_login" v-if="buttonControl">Submit</button>
                    <button class="right-button" @click="submit_new" v-if="!buttonControl">Submit</button></div>
                </div>
                <div class="right" id="show_new" v-if="showControl2">
                    <div>Your account is created : &nbsp;{{account}}<br/>
                        account information will saves at local : &nbsp;keystore<br/>
                        please keep it properly.<br/><hr/>Do you want to save and login with this account?</div>
                    <div class="right-pwd"><button class="right-button" @click="hide">No</button>
                        <button class="right-button" @click="submit_keystore">Yes</button></div>
                </div>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import {lfg} from '../listenFromGo'
import {dl_db, mt_db, acc_db, options} from "../options"
export default {
    name: "login",
    data() {
        return {
            account: "",
            password: "",
            showControl1: false,
            showControl2: false,
            buttonControl: true,
            describe: "",
        }
    },
    methods: {
        right: function (description) {
            this.showControl1 = true;this.showControl2 = false
            switch (description) {
                case "Login": this.buttonControl = true;break
                case "New": this.buttonControl = false;break
            }
            this.describe = description + ":"
        },
        hide: function () {this.showControl1 = false;this.showControl2 = false;this.password = ""},
        submit_login: function () {
            let pwd = this.password
            this.password = ""
            let _this = this
            astilectron.sendMessage({Name: "login.verify", Payload: {account: this.account,
                        password: pwd}}, function (message) {
                if (message.payload) {
                    _this.$router.push({ name: 'home', params: {acc: _this.account}})
                } else {
                    alert("account or password is wrong.")
                }
            })
        },
        submit_new: function () {
            let _this = this
            astilectron.sendMessage({Name: "create.new.account", Payload: {address: this.password}}, function (message) {
                acc_db.write(message.payload)
                _this.account = message.payload
                _this.showControl1 = false;_this.showControl2 = true
            })
        },
        submit_keystore: function () {
            let pwd = this.password
            this.password = ""
            let _this = this
            astilectron.sendMessage({Name: "save.keystore", Payload: {account: this.$store.state.account,
                    password: pwd}}, function (message) {
                if (message.payload) {
                    _this.$router.push({ name: 'home', params: {acc: _this.account}})
                } else {
                    alert("save account information failed.")
                }
            })
        }
    },
    created() {
        this.password = "";this.describe = "";this.account = ""
        let _this = this
        dl_db.init(this)
        mt_db.init(this)
        acc_db.init(this)
        document.addEventListener('astilectron-ready', function() {
            lfg.listen()
            options.init(_this)
        })
    }
}
</script>

<style>
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
    height: 500px;
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
    border-radius: 4px;
    width: 65%;
    margin: 5px 17%;
}
.right {
    height: 300px;
    padding: 100px 20%;
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
    border-radius: 4px;
    width: 40%;
    margin: 0 2.5%;
}
</style>
