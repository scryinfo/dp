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
            <el-col :span="14">
                用户昵称：{{ this.$store.state.nickname }}
            </el-col>
            <el-col :span="6">
                <el-input v-model="nickName" placeholder="nickname"></el-input>
            </el-col>
            <el-col :span="4" >
                <el-button size="mini" type="primary" class="section-item-right" @click="modifyNickname">修改昵称</el-button>
            </el-col>
        </el-row>
    </section>
</template>

<script>
import {acc_db} from "../../utils/DBoptions";
export default {
    name: "nickname.vue",
    data () {
        return {
            nickName: "" // todo: limit, can not modify nickName as 42chars string start with "0x".
        }
    },
    methods: {
        modifyNickname: function () {
            this.$store.state.nickname = this.nickName;
            let _nickname = this;
            acc_db.read(this.$store.state.account, function (accInstance) {
                acc_db.write({
                    address: accInstance.address,
                    nickname: _nickname.nickName,
                    fromBlock: accInstance.fromBlock,
                    isVerifier: accInstance.isVerifier
                });
            })
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
.nickname_button {

}
</style>