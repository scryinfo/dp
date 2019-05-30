<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-row>
            <el-col :span="21" class="section-item before-name token-height">
                以太币余额：&nbsp;{{ this.$store.state.ethBalance }}&nbsp;wei
                <span class="token-time">查询时间：{{ this.$store.state.ethTime }}</span>
            </el-col>
            <el-col :span="3" class="section-item token-height">
                <s-f-t button-name="余额查询" @password="getEthBalance"></s-f-t>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="21" class="section-item before-name token-height">
                &nbsp;token&nbsp;余额：&nbsp;{{ this.$store.state.tokenBalance }}&nbsp;DDD
                <span class="token-time">查询时间：{{ this.$store.state.tokenTime }}</span>
            </el-col>
            <el-col :span="3" class="section-item token-height">
                <s-f-t button-name="余额查询" @password="getTokenBalance"></s-f-t>
            </el-col>
        </el-row>
    </section>
</template>

<script>
import {connect} from "../../utils/connect";
import SFT from "../templates/simple_function_template.vue";
export default {
    name: "balance.vue",
    data () {
        return {

        }
    },
    methods: {
        getEthBalance: function (pwd) {
            connect.send({Name: "get.eth.balance", Payload: {password: pwd}}, function (payload, _this) {
                console.log("查询以太币余额成功：", payload.split("|")[0]);
                _this.$store.state.ethBalance = payload.split("|")[0];
                _this.$store.state.ethTime = payload.split("|")[1];
            }, function (payload, _this) {
                console.log("查询以太币余额成功：", payload);
                _this.$alert(payload, "查询以太币余额失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        getTokenBalance: function (pwd) {
            connect.send({Name: "get.token.balance", Payload: {password: pwd}}, function (payload, _this) {
                console.log("查询token余额成功：", payload.split("|")[0]);
                _this.$store.state.tokenBalance = payload.split("|")[0];
                _this.$store.state.tokenTime = payload.split("|")[1];
            }, function (payload, _this) {
                console.log("查询token余额成功：", payload);
                _this.$alert(payload, "查询token余额失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    },
    components: {
        SFT
    }
}
</script>

<style scoped>
.before-name {
    padding-left: 30px;
}
.token-height {
    height: 50px;
    display: flex;
    align-items: center;
}
.token-time {
    margin-left: 20px;
    font-size: 10px;
}
</style>