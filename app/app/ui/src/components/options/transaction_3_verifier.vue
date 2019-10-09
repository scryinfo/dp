<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <div v-if="!showControl">
            <el-col :span="24" class="section-item">
                <s-f-t button-name="注册成为验证者" @password="register"></s-f-t>
            </el-col>
        </div>
        <div v-if="showControl">
            <el-col :span="21" class="section-item">
                <c-f-t button-name="验证数据" dialog-title="验证数据：" @password="vote" :button-disabled="buttonDisabled(2)">
                    <p>是否建议购买：</p>
                    <p><el-switch v-model="verify.suggestion" active-text="是" inactive-text="否"></el-switch></p>
                    <p><el-input v-model="verify.comment" placeholder="评论" clearable></el-input></p>
                </c-f-t>
            </el-col>
            <el-col :span="3" class="section-item">
                <el-button size="mini" type="primary" @click="initTxV">刷新列表</el-button>
            </el-col>
            <el-table :data="this.$store.state.transactionverifier.slice((curPage-1)*pageSize, curPage*pageSize)"
                      highlight-current-row border :height=height @current-change="currentChange">
                <el-table-column prop="PublishID" label="数据ID" show-overflow-tooltip></el-table-column>
                <el-table-column prop="TransactionID" label="交易ID" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Title" label="标题" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Price" label="价格" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Keys" label="标签" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Description" label="描述" show-overflow-tooltip></el-table-column>
            </el-table>
            <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                           layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
            ></el-pagination>
        </div>
    </section>
</template>

<script>
import {connect} from "../../utils/connect.js";
import {acc_db, tx_db} from "../../utils/DBoptions.js";
import {utils} from "../../utils/utils.js";
import SFT from "../templates/simple_function_template.vue";
import CFT from "../templates/complex_function_template.vue";
export default {
    name: "transaction_3_verifier.vue",
    data () {
        return {
            selectedTx: "",     // txID: ""
            curPage: 1,
            pageSize: 6,
            total: 0,
            height: window.innerHeight - 170,
            showControl: false,
            txState: "Begin",
            verify: {
                suggestion: false,
                comment: ""
            }
        }
    },
    methods: {
        setCurPage: function (curPageReturn) { this.curPage = curPageReturn; },
        setPageSize: function (newPageSize) { this.pageSize = newPageSize; },
        currentChange: function (curRow) {
            this.selectedTx = curRow.TransactionID;
            this.txState = curRow.State;
        },
        buttonDisabled: function (funcNum) {
            return utils.functionDisabled(funcNum, this.txState);
        },
        initTxV: function () {
            tx_db.initVerifier(this);
        },
        refresh: function () {
            let _this = this;
            acc_db.read(this.$store.state.account, function (accInstance) {
                _this.showControl = accInstance.isVerifier;
            });
        },
        register: function (pwd) {
            connect.send({Name:"register", Payload:{password: pwd}}, function (payload, _this) {
                console.log("注册成为验证者成功", payload);
            }, function (payload, _this) {
                console.log("注册成为验证者失败：", payload);
                _this.$alert(payload, "注册成为验证者失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        vote: function (pwd) {
            connect.send({Name:"vote", Payload:{password: pwd, tID: this.selectedTx, verify: this.verify}},
                function (payload, _this) {
                _this.verify = {suggestion: false, comment: ""};
                console.log("验证成功", payload);
            }, function (payload, _this) {
                console.log("验证失败：", payload);
                _this.$alert(payload, "验证失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    },
    components: {
        SFT,
        CFT
    },
    computed: {
        listenTxVRefresh() {
            return this.$store.state.transactionverifier;
        }
    },
    watch: {
        listenTxVRefresh: function () {
            this.curPage = 1;
            this.total = this.$store.state.transactionverifier.length;
        }
    },
    created () {
        this.total = this.$store.state.transactionverifier.length;
        this.refresh();
    }
}
</script>

<style scoped>

</style>