<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->

<template>
    <!-- 管理员功能，提供一些便捷功能供测试使用，正式版本当做彩蛋随项目赠送。 -->
    <!-- ps: 彩蛋[cai dan] colorful egg? -> extra scene :) -->
    <section class="administrator">
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="welcome">Welcome</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="logoutSimulate">Logout Simulate</el-button></el-col>
        <div>
            <el-col :span="5" class="section-item">
                <el-select size="mini" v-model="modifyItem" placeholder="Choose param" clearable allow-create filterable>
                    <el-option v-for="item in modifyList" :key="item.paramName" :value="item.paramName" :label="item.paramName"></el-option>
                </el-select>
            </el-col>
            <el-col :span="9" class="section-item">
                <el-input size="mini" v-model="newParamValue" placeholder="new param value" clearable></el-input>
            </el-col>
            <el-col :span="10" class="section-item">
                <el-button size="mini" type="primary" @click="modifyParam">ModifyParam</el-button>
            </el-col>
        </div>
    </section>
</template>

<script>
import {connect} from "../../utils/connect.js";
import home from "../home.vue";
export default {
    name: "ES_administrator.vue",
    data () {
        return {
            modifyItem: "",
            newParamValue: "", // avoid param bigger than float64.MAX, json unmarshal in go will wrong.
            modifyList: [
                {
                    paramName: "VerifierNum"
                }
            ]
        }
    },
    methods: {
        welcome: function () {
            this.$notify({
                title: "彩蛋: ",
                dangerouslyUseHTMLString: true,
                message: '谢谢你使用我的程序!&nbsp;<strong>:)</strong>',
                position: "top-left"
            });
        },

        logoutSimulate: function () {
            home.methods.logout();
        },

        modifyParam: function () {
            connect.send({Name: "modifyContractParam", Payload: {modifyContractParam: {paramName: this.modifyItem, paramValue: this.newParamValue}}},
                function (payload, _this) {
                    console.log("modify param success: ", payload); // payload is nothing now :( think if it need some param from go?
                }, function (payload, _this) {
                    console.log("modify param failed! ", payload);
                });
        }
    }
}
</script>

<style>
.administrator {
    background-color: lightgrey;
}
</style>