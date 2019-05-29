<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <!-- 功能测试界面，针对添加的功能制作的小demo，完成开发后删除本文件或删除相关路由设置 -->
    <div>
        <c-f-t :button-name="buttonName" button-type="primary" :pre-params="[this.key === 10]" @pre="f2" @password="f1"
               :dialog-title="dialogTitle" class="inALine">
            <el-input v-model="key"></el-input>
        </c-f-t>
        <s-f-t :button-name="buttonName" button-type="danger" :pre-params="[this.key === 11]" @pre="f2" @password="f1" :dialog-title="dialogTitle">
            <el-slider v-model="key"></el-slider>
        </s-f-t>
    </div>
</template>

<script>
import CFT from "../templates/complex_function_template.vue";
import SFT from "../templates/simple_function_template.vue";
export default {
    name: "home",
    data () {
        return {
            buttonName: "A function button. ",
            dialogTitle: "Test title",
            key: 10
        }
    },
    methods: {
        f1: function (pwd) {
            console.log("Click function passed. ", pwd);
        },
        f2: function (params) {
            console.log(params);
        }
    },
    components: {
        CFT,
        SFT
    },
    created () {
        this.$store.state.account = "0x123456...[40chars]"
    }
}
</script>

<style scoped>
.inALine {
    float: left;
}
</style>