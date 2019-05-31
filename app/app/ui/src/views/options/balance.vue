<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-row class="section-item center token-height">
            <el-col :span="21">
                以太币余额：&nbsp;{{ ethBalance }}&nbsp;wei
                <span class="token-time">查询时间：{{ ethTime }}</span>
            </el-col>
            <el-col :span="3" class="section-item-right">
                <s-f-t button-name="余额查询" @password="getEthBalance"></s-f-t>
            </el-col>
        </el-row>
        <el-row class="section-item center token-height">
            <el-col :span="21">
                &nbsp;token&nbsp;余额：&nbsp;{{ tokenBalance }}&nbsp;DDD
                <span class="token-time">查询时间：{{ tokenTime }}</span>
            </el-col>
            <el-col :span="3" class="section-item-right">
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
            ethBalance: "-",
            ethTime: "-",

            tokenBalance: "-",
            tokenTime: "-"
        }
    },
    methods: {
        getEthBalance: function (pwd) {
            let _balance = this;
            connect.send({Name: "get.eth.balance", Payload: {password: pwd}}, function (payload, _this) {
                console.log("查询以太币余额成功：", payload.split("|")[0]);
                _balance.ethBalance = payload.split("|")[0];
                _balance.ethTime = payload.split("|")[1];
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
            let _balance = this;
            connect.send({Name: "get.token.balance", Payload: {password: pwd}}, function (payload, _this) {
                console.log("查询token余额成功：", payload.split("|")[0]);
                _balance.tokenBalance = payload.split("|")[0];
                _balance.tokenTime = payload.split("|")[1];
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
.token-height {
    height: 80px;
}
.token-time {
    margin-left: 20px;
    font-size: 10px;
}
</style>