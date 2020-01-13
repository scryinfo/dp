<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-row class="section-item center token-item">
            <el-col :span="2">以太币</el-col>
            <el-col :span="2">余额：</el-col>
            <el-col :span="5" class="token-middle-font">{{ this.$store.state.balance[0].Balance }}</el-col>
            <el-col :span="2">wei</el-col>
            <el-col :span="2" class="token-small-font">查询时间：</el-col>
            <el-col :span="8" class="token-small-font">{{ this.$store.state.balance[0].Time }}</el-col>
            <el-col :span="3" class="section-item-right">
                <s-f-t button-name="余额查询" button-size="small" @password="getEthBalance"></s-f-t>
            </el-col>
        </el-row>
        <el-row class="section-item center token-item">
            <el-col :span="2">token</el-col>
            <el-col :span="2">余额：</el-col>
            <el-col :span="5" class="token-middle-font">{{ this.$store.state.balance[1].Balance }}</el-col>
            <el-col :span="2">DDD</el-col>
            <el-col :span="2" class="token-small-font">查询时间：</el-col>
            <el-col :span="8" class="token-small-font">{{ this.$store.state.balance[1].Time }}</el-col>
            <el-col :span="3" class="section-item-right">
                <s-f-t button-name="余额查询" button-size="small" @password="getTokenBalance"></s-f-t>
            </el-col>
        </el-row>
    </section>
</template>

<script lang="ts">
    import { Component,Vue, } from 'vue-property-decorator'
import connects from "../../utils/connect";
import SFT from "../templates/simple_function_template.vue";

    @Component({
        // 所有的组件选项都可以放在这里
        components: {
            SFT
        }
    })

export default class Balance extends Vue{
    getEthBalance (pwd:string) {
        let _balance = this;
        connects.send({Name: "getEthBalance", Payload: {password: pwd}}, function (payload, _this) {
            console.log("查询以太币余额成功：", payload.split("|")[0]);
            _balance.$store.state.balance[0].Balance = payload.split("|")[0];
            _balance.$store.state.balance[0].Time = payload.split("|")[1];
        }, function (payload, _this) {
            console.log("查询以太币余额失败：", payload);
            _this.$alert(payload, "查询以太币余额失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    }

    getTokenBalance (pwd:string) {
        let _balance = this;
        connects.send({Name: "getTokenBalance", Payload: {password: pwd}}, function (payload, _this) {
            console.log("查询token余额成功：", payload.split("|")[0]);
            _balance.$store.state.balance[1].Balance = payload.split("|")[0];
            _balance.$store.state.balance[1].Time = payload.split("|")[1];
        }, function (payload, _this) {
            console.log("查询token余额失败：", payload);
            _this.$alert(payload, "查询token余额失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            });
        });
    }
}
</script>

<style scoped>
.token-item {
    height: 80px;
    font-size: 20px;
}
.token-middle-font {
    font-size: 16px;
}
.token-small-font {
    font-size: 12px;
}
</style>
