<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-col :span="21" class="section-item">
            <s-f-t button-name="取消交易" button-type="danger" @password="cancelBuying"></s-f-t>
            <s-f-t button-name="购买数据" @password="purchase"></s-f-t>
            <s-f-t button-name="解密数据" @password="decrypt"></s-f-t>
            <c-f-t button-name="确认数据" dialog-title="判断原始数据真实性：" @password="confirm"
                   :pre-params="[this.selectedTx.SupportVerify]" @pre="confirmPre">
                <p v-if="supportVerify">判断原始数据真实性，如果你认为原始数据是假的，我们将为你启动仲裁流程：</p>
                <p v-if="!supportVerify">判断原始数据真实性，点击“输入密码”按钮完成交易。</p>
                <p><el-switch v-model="confirmData" active-text="真" inactive-text="假"></el-switch></p>
            </c-f-t>
            <c-f-t button-name="评价验证者" dialog-title="评价验证者：" @password="credit" @pre="creditPre"
                   :pre-params="[this.selectedTx.Verifier1Response !== '', this.selectedTx.Verifier2Response !== '']">
                <p>验证者1:</p>
                <p><el-slider v-model="verifier1Credit" max="5" v-if="verifier1Revert" show-input></el-slider>
                <span v-if="!verifier1Revert">交易未进入验证流程或验证者未回复</span></p>
                <p>验证者2:</p>
                <p><el-slider v-model="verifier2Credit" max="5" v-if="verifier2Revert" show-input></el-slider>
                <span v-if="!verifier2Revert">交易未进入验证流程或验证者未回复</span></p>
            </c-f-t>
        </el-col>
        <el-col :span="3" class="section-item section-item-right">
            <el-button size="mini" type="primary" @click="initTxB">刷新列表</el-button>
        </el-col>

        <el-table :data="this.$store.state.transactionbuy.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border :height=height @current-change="currentChange">
            <el-table-column type="expand">
                <el-form slot-scope="props" label-position="left" class="tx-table-expand">
                    <el-form-item label="标题"><span>{{ props.row.Title }}</span></el-form-item>
                    <el-form-item label="价格"><span>{{ props.row.Price }}</span></el-form-item>
                    <el-form-item label="标签"><span>{{ props.row.Keys }}</span></el-form-item>
                    <el-form-item label="描述"><span>{{ props.row.Description }}</span></el-form-item>
                    <el-form-item label="卖家"><span>{{ props.row.Seller }}</span></el-form-item>
                    <el-form-item label="状态"><span>{{ props.row.State }}</span></el-form-item>
                    <el-form-item label="验证者回复1"><span>{{ props.row.Verifier1Response }}</span></el-form-item>
                    <el-form-item label="验证者回复2"><span>{{ props.row.Verifier2Response }}</span></el-form-item>
                    <el-form-item label="仲裁结果"><span>{{ props.row.ArbitrateResult }}</span></el-form-item>
                </el-form>
            </el-table-column>
            <el-table-column prop="Title" label="标题" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Price" label="价格" show-overflow-tooltip></el-table-column>
            <el-table-column prop="State" label="状态" show-overflow-tooltip></el-table-column>
        </el-table>
        <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                       layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
        ></el-pagination>
    </section>
</template>

<script>
import {connect} from "../../utils/connect";
import {txBuyer_db} from "../../utils/DBoptions"
import SFT from "../templates/simple_function_template.vue"
import CFT from "../templates/complex_function_template.vue"
export default {
    name: "transaction_buy.vue",
    data () {
        return {
            selectedTx: {},  // {tID: "", User: "", MetaDataIDEncrypt: "", MetaDataExtension: "",
                             //  Verifier1Response: "", Verifier2Response: "", SupportVerify: false}
            curPage: 1,
            pageSize: 6,
            total: 0,
            height: window.innerHeight - 170,
            supportVerify: false,
            confirmData: false,
            verifier1Revert: false,
            verifier1Credit: 5,
            verifier2Revert: false,
            verifier2Credit: 5
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn;},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize;},
        currentChange: function (curRow) {
            this.selectedTx = {
                TransactionID: curRow.TransactionID,
                User: curRow.Buyer,
                Price: curRow.Price,
                MetaDataIDEncrypt: curRow.MetaDataIDEncWithBuyer,
                MetaDataExtension: curRow.MetaDataExtension,
                SupportVerify: curRow.SupportVerify,
                Verifier1Response: curRow.Verifier1Response,
                Verifier2Response: curRow.Verifier2Response
            };
        },
        initTxB: function () {
            txBuyer_db.init(this);
        },
        cancelBuying: function (pwd) {
            connect.send({Name:"cancel", Payload:{password: pwd, tID: this.selectedTx}}, function (payload, _this) {
                console.log("取消交易成功", payload);
            }, function (payload, _this) {
                console.log("取消交易失败：", payload);
                _this.$alert(payload, "取消交易失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        purchase: function (pwd) {
            connect.send({Name:"purchase", Payload:{password: pwd, tID: this.selectedTx}}, function (payload, _this) {
                console.log("购买数据成功", payload);
            }, function (payload, _this) {
                console.log("购买数据失败：", payload);
                _this.$alert(payload, "购买数据失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
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
        confirmPre: function (array) {
            this.supportVerify = array[0];
        },
        confirm: function (pwd) {
            connect.send({Name:"confirm", Payload:{password: pwd, tID: this.selectedTx, confirmData: this.confirmData}}, function (payload, _this) {
                _this.supportVerify = false;
                console.log("确认数据成功", payload);
            }, function (payload, _this) {
                console.log("确认数据失败：", payload);
                _this.$alert(payload, "确认数据失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        creditPre: function (array) {
            this.verifier1Revert = array[0];
            this.verifier2Revert = array[1];
        },
        credit: function (pwd) {
            connect.send({Name:"credit", Payload:{password: pwd, tID: this.selectedTx, credit: {
                        verifier1Revert: this.verifier1Revert, verifier1Credit: this.verifier1Credit,
                        verifier2Revert: this.verifier2Revert, verifier2Credit: this.verifier2Credit}}},
                function (payload, _this) {
                    _this.verifier1Revert = false;
                    _this.verifier2Revert = false;
                    console.log("评价验证者成功", payload);
                }, function (payload, _this) {
                    console.log("评价验证者失败：", payload);
                    _this.$alert(payload, "评价验证者失败！", {
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
        listenTxBRefresh() {
            return this.$store.state.transactionbuy;
        }
    },
    watch: {
        listenTxBRefresh: function () {
            this.curPage = 1;
            this.total = this.$store.state.transactionbuy.length;
        }
    },
    created () {
        this.total = this.$store.state.transactionbuy.length;
    }
}
</script>

<style>

</style>
