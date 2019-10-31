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
                <el-input v-model="nickname" placeholder="nickname, suggest different structure from address"></el-input>
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
import {connect} from "../../utils/connect";
export default {
    name: "nickname.vue",
    data () {
        return {
            nickname: "",
        }
    },
    methods: {
        modifyNickname: function () {
            if (this.validNickname()) {
                this.$alert("请勿使用标准地址类型（以'0x'开头的40位十六进制数）作为昵称。", "非法昵称！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
                return;
            }
            connect.send({Name:"modifyNickname",Payload:{Nickname: this.nickname}}, function (payload, _this) {
                console.log("修改昵称成功：", payload);
                _this.$store.state.nickname = _this.nickname;
            }, function (payload, _this) {
                console.log("修改昵称失败！", payload);
                _this.$alert(payload, "修改昵称失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        validNickname: function () {
            let reg = /^0x(\d|[a-f]){40}$/i; // match standard 40 digits Hexadecimal number
            return reg.test(this.nickname);
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