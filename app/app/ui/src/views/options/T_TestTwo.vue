<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <!-- 功能测试界面2，针对添加的功能制作的小demo，完成开发后删除本文件或删除相关路由设置 -->
    <div>
        <el-button :type="buttonType" @click="dialog = true">{{ buttonName }}</el-button>

        <el-dialog :visible.sync="dialog" :title="dialogTitle + ': ' ">
            <slot></slot>
            <div slot="footer">
                <el-button @click="cancelFunc(buttonName)">Cancel</el-button>
                <el-button type="primary" @click="submitFunc(clickFunc)">Submit</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
export default {
    name: "child",
    data () {
        return {
            input: "default string. ",
            dialog: false,
            value: "default string. "
        }
    },
    props: {
        buttonName: String,
        buttonType: String,
        dialogBody: String,
        dialogTitle: String,
        pre: Function
    },
    methods: {
        submitFunc: function (clickFunc) {
            clickFunc();
            this.dialog = false;
        },
        cancelFunc: function (buttonName) {
            this.dialog = false;
            this.$message({
                type: "info",
                message: "取消" + buttonName
            });
        }
    }
}
</script>

<style scoped>

</style>