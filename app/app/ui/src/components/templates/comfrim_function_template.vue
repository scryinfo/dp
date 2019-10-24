<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <!-- complex function template has a two-layer nested dialog. -->
    <div class="inALine">
        <el-button :size="buttonSize" :type="buttonType" @click="pre(preParams)" :disabled="buttonDisabled">{{ buttonName }}</el-button>

        <el-dialog :visible.sync="dialog" :title="dialogTitle">
            <div slot="footer">
                <el-button @click="cancelFunc(buttonName)">取消</el-button>
                <el-button type="primary" @click="dialogInner = true">确定删除</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
export default {
    name: "CFT1",
    data () {
        return {
            dialog: false,
            dialogInner: false,
            delID: ""
        }
    },
    props: {
        buttonSize: {
            type: String,
            default: "mini"
        },
        buttonType: {
            type: String,
            default: "primary"
        },
        buttonName: {
            type: String,
            default: "default button name"
        },
        dialogTitle: {
            type: String,
            default: "default dialog title"
        },
        buttonDisabled: {
            type: Boolean,
            default: false
        },
        preParams: Array,
    },
    methods: {
        pre: function (params) {
            this.dialog = true;
            if (!!params) {
                let result = [];
                for (let i = 0; i < params.length; i++) {
                    result.push(params[i])
                }
                this.$emit("pre", result)
            }
        },
        cancelFunc: function (name) {
            this.dialog = false;
            this.$message({
                type: "info",
                message: "取消" + name
            });
        },
        submitFunc: function () {
            this.dialogInner = false;
            this.dialog = false;
            let pwd = this.password;
            this.password = "";
            this.$emit("password", pwd);
        }
    }
}
</script>

<style scoped>
.inALine {
    float: left;
    margin: 0 5px;
}
</style>