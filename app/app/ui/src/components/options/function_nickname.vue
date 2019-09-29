<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-row>
            <el-col :span="24" class="section-item center nickname_address">
                用户地址：{{ this.$store.state.account }}
            </el-col>
        </el-row>
        <el-row class="section-item center nickname_nickname">
            <el-col :span="12">
                用户昵称：{{ this.$store.state.nickname }}
            </el-col>
            <el-col :span="9">
                <el-input v-model="nickName" placeholder="nickname, suggest different structure from address"></el-input>
            </el-col>
            <el-col :span="3" >
                <el-button size="mini" type="primary" class="section-item-right" @click="modifyNickname">修改昵称</el-button>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="24" class="section-item center">
                <el-button size="mini" type="primary" class="section-item-right" @click="DBBackup">用户信息备份</el-button>
                <el-button size="mini" type="primary" class="section-item-right" @click="DBRestore">用户信息还原</el-button>
            </el-col>
        </el-row>
    </section>
</template>

<script>
import {acc_db} from "../../utils/DBoptions";
import {connect} from "../../utils/connect";
export default {
    name: "nickname.vue",
    data () {
        return {
            nickName: "",
            accBackup: []
        }
    },
    methods: {
        modifyNickname: function () {
            let _nickname = this;
            if (_nickname.validNickname()) {
                _nickname.$alert("请勿使用标准地址类型（以'0x'开头的40位十六进制数）作为昵称。", "非法昵称！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
                return;
            }
            this.$store.state.nickname = this.nickName;
            acc_db.read(this.$store.state.account, function (accInstance) {
                acc_db.write({
                    address: accInstance.address,
                    nickname: _nickname.nickName,
                    fromBlock: accInstance.fromBlock,
                    isVerifier: accInstance.isVerifier
                });
            })
        },
        validNickname: function () {
            let reg = /^0x(\d|[a-f]){40}$/i; // match standard 40 digits Hexadecimal number
            return reg.test(this.nickName);
        },
        DBBackup: function () {
            let _this = this;
            acc_db.readAll(function (accs) {
                for (let i = 0; i < accs.length; i++) {
                    _this.accBackup.push({
                        "address": accs[i].address,
                        "nickname": accs[i].nickname
                    });
                }
                connect.send({Name:"accountsBackup",Payload:{"accounts": _this.accBackup}}, function (payload, _this) {
                    console.log("用户信息备份成功", payload);
                    _this.accBackup = [];
                }, function (payload, _this) {
                    console.log("用户信息备份失败：", payload);
                    _this.accBackup = [];
                    _this.$alert(payload, "用户信息备份失败！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    });
                })
            });
        },
        DBRestore: function () {
            connect.send({Name:"accountsRestore",Payload:{}}, function (payload, _this) {
                let accs = JSON.parse(payload);
                console.log("用户信息还原成功", accs.accounts);
                for (let i = 0; i < accs.accounts.length; i++) {
                    acc_db.write({
                        "address": accs.accounts[i].address,
                        "nickname": accs.accounts[i].nickname,
                        "fromBlock": 1,
                        "isVerifier": false
                    });
                }
            }, function (payload, _this) {
                console.log("用户信息还原失败：", payload);
                _this.$alert(payload, "用户信息还原失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    }
}
</script>

<style scoped>
.nickname_address {
    height: 100px;
}
.nickname_nickname {
    height: 80px;
}
</style>