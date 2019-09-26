<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <div v-if="!showControl">
            <el-col :span="24" class="section-item">
                你还不是验证者！请前往验证界面注册成为验证者。
            </el-col>
        </div>
        <div v-if="showControl">
            <el-col :span="21" class="section-item">
                <s-f-t button-name="解密数据" @password="decrypt" :button-disabled="buttonDisabled(3)"></s-f-t>
                <c-f-t button-name="仲裁数据" dialog-title="仲裁数据：" @password="Arbitrate" :button-disabled="buttonDisabled(3)">
                    <p>判断数据真实性：</p>
                    <p><el-switch v-model="arbitrateResult" active-text="真" inactive-text="假"></el-switch></p>
                </c-f-t>
            </el-col>
            <el-col :span="3" class="section-item section-item-right">
                <el-button size="mini" type="primary" @click="initTxA">刷新列表</el-button>
            </el-col>

            <el-table :data="this.$store.state.transactionarbitrator.slice((curPage-1)*pageSize, curPage*pageSize)"
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
import {acc_db, txArbitrator_db} from "../../utils/DBoptions.js";
import {utils} from "../../utils/utils.js";
import SFT from "../templates/simple_function_template.vue";
import CFT from "../templates/complex_function_template.vue";
export default {
    name: "transaction_4_arbitrator.vue",
    data () {
        return {
            selectedTx: {},     // {txID: "", User: "", MetaDataIDEncrypt: ""}
            arbitrateResult: false,
            height: window.innerHeight - 170,
            showControl: false,
            txState: "Begin",
            curPage: 1,
            pageSize: 6,
            total: 0
        }
    },
    methods: {
        setCurPage: function (curPageReturn) { this.curPage = curPageReturn; },
        setPageSize: function (newPageSize) { this.pageSize = newPageSize; },
        currentChange: function (curRow) {
            this.selectedTx = {
                TransactionID: curRow.TransactionID,
                User: this.$store.state.account,
                MetaDataIDEncrypt: curRow.MetaDataIDEncWithArbitrator,
                MetaDataExtension: curRow.MetaDataExtension
            };
            this.txState = curRow.State;
        },
        buttonDisabled: function (funcNum) {
            return utils.functionDisabled(funcNum, this.txState);
        },
        initTxA: function () {
            txArbitrator_db.init(this);
        },
        refresh: function () {
            let _this = this;
            acc_db.read(this.$store.state.account, function (accInstance) {
                _this.showControl = accInstance.isVerifier;
            })
        },
        decrypt: function (pwd) {
            connect.send({Name:"decrypt", Payload:{password: pwd, tID: this.selectedTx}}, function (payload, _this) {
                console.log("解密数据成功", payload);
                _this.$alert(payload, "原始数据：", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "info"
                });
            }, function (payload, _this) {
                console.log("解密数据失败：", payload);
                _this.$alert(payload, "解密数据失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        Arbitrate: function (pwd) {
            connect.send({Name: "arbitrate", Payload: {password: pwd, tID: this.selectedTx, arbitrateResult: this.arbitrateResult}}, function (payload, _this) {
                console.log("仲裁成功", payload);
            }, function (payload, _this) {
                console.log("仲裁失败：", payload);
                _this.$alert(payload, "仲裁失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            })
        }
    },
    components: {
        SFT,
        CFT
    },
    computed: {
        listenTxARefresh() {
            return this.$store.state.transactionarbitrator
        }
    },
    watch: {
        listenTxARefresh: function () {
            this.curPage = 1
            this.total = this.$store.state.transactionarbitrator.length
        }
    },
    created () {
        this.total = this.$store.state.transactionarbitrator.length
        this.refresh()
    }
}
</script>

<style scoped>

</style>