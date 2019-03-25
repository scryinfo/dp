<template>
    <div>
        <el-row class="row">
            <el-col :span="24"><div class="top">My Astilectron demo</div></el-col>
            <el-col :span="8">
                <div class="left">
                    <div class="left-explain">Select Account:</div>
                    <el-select class="left-account" v-model="account" placeholder="select account"
                        clearable allow-create filterable>
                        <el-option v-for="acc in this.$store.state.accounts" :key="acc.address"
                                   :value="acc.address" :label="acc.address"></el-option>
                    </el-select>
                    <el-button class="left-button" @click="right('Login')">Login</el-button>
                    <el-button class="left-button" @click="right('New')">Create New Account</el-button>
                </div>
            </el-col>
            <el-col :span="16">
                <div class="right" id="show" v-if="showControl1">
                    <div class="right-show">{{describe}}
                        <el-input class="right-pwd" v-model="password" placeholder="password"
                                  clearable show-password></el-input>
                    </div>
                    <el-button class="right-button" @click="hide">Back</el-button>
                    <el-button class="right-button" @click="submit_login" v-if="buttonControl">Submit</el-button>
                    <el-button class="right-button" @click="submit_new" v-if="!buttonControl">Submit</el-button>
                </div>
                <div class="right" id="show_new" v-if="showControl2">
                    Your account is created : &nbsp;{{account}}<br/>
                    account information will saves at local :&nbsp;&nbsp;&nbsp;indexDB<br/>
                    please remember it.<br/><hr/><br/>Do you want login with this account?
                    <div class="right-pwd">
                        <el-button class="right-button" @click="hide">No</el-button>
                        <el-button class="right-button" @click="submit_keystore">Yes</el-button>
                    </div>
                </div>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import {dl_db, tx_db, acc_db} from "../DBoptions.js"
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
                case "Login": this.buttonControl = true;break
                case "New": this.buttonControl = false;break
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
                    console.log("Node: login.verify failed. ", message)
                    _this.$alert(message.payload, "Error: account or password is wrong.", {
                        confirmButtonText: "I've got it.",
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
                    _this.account = message.payload
                    _this.showControl1 = false;_this.showControl2 = true
                }else {
                    console.log("Node: create.newAcc failed. ", message)
                    _this.$alert(message.payload, "Error: create new account failed.", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        },
        submit_keystore: function () {
            let pwd = this.password
            this.password = ""
            let _this = this
            astilectron.sendMessage({Name: "save.keystore", Payload: {account: this.$store.state.account,
                    password: pwd}}, function (message) {
                if (message.name !== "error") {
                    acc_db.write({
                        address: _this.account,
                        fromBlock: 1
                    })
                    _this.$router.push({ name: "home", params: {acc: _this.account}})
                } else {
                    console.log("Node: save.keystore failed. ", message)
                    _this.$alert(message.payload, "Error: save account information failed.", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        }
    },
    created() {
        this.password = "";this.describe = "";this.account = ""
        let _this = this
        dl_db.init(this)
        tx_db.init(this)
        acc_db.init(this)
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
    width: 65%;
    padding: 5px 20px;
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
    width: 40%;
    margin: 0 2.5%;
}
</style>
