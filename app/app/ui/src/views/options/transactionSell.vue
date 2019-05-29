<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-col :span="21" class="section-item">
            <s-f-t button-name="再加密数据" @password="reEncrypt"></s-f-t>
        </el-col>
        <el-col :span="3" class="section-item">
            <el-button size="mini" type="primary" @click="initTxS">刷新列表</el-button>
        </el-col>

        <el-table :data="this.$store.state.transactionsell.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border :height=height @current-change="currentChange">
            <el-table-column type="expand">
                <el-form slot-scope="props" label-position="left" class="tx-table-expand">
                    <el-form-item label="标题"><span>{{ props.row.Title }}</span></el-form-item>
                    <el-form-item label="价格"><span>{{ props.row.Price }}</span></el-form-item>
                    <el-form-item label="标签"><span>{{ props.row.Keys }}</span></el-form-item>
                    <el-form-item label="描述"><span>{{ props.row.Description }}</span></el-form-item>
                    <el-form-item label="状态"><span>{{ props.row.State }}</span></el-form-item>
                </el-form>
            </el-table-column>
            <el-table-column prop="Title" label="标题" show-overflow-tooltip></el-table-column>
            <el-table-column prop="State" label="状态" show-overflow-tooltip></el-table-column>
        </el-table>
        <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                       layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
        ></el-pagination>
    </section>
</template>

<script>
import {connect} from "../../utils/connect";
import {txSeller_db} from "../../utils/DBoptions";
import SFT from "../templates/simple_function_template.vue";
export default {
    name: "TransactionSell",
    data () {
        return {
            selectedTx: {},  // {tID: "", Buyer: "", Seller: "", MetaDataIDEncWithSeller: "", pID: ""}
            curPage: 1,
            pageSize: 6,
            total: 0,
            height: window.innerHeight - 170,
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn;},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize;},
        currentChange: function (curRow) {
            this.selectedTx = {
                TransactionID: curRow.TransactionID,
                Buyer: curRow.Buyer,
                Seller: curRow.Seller,
                PublishID: curRow.PublishID,
                MetaDataIDEncWithSeller: curRow.MetaDataIDEncWithSeller // WSConnect between go and js buy not show out to user.
            };
        },
        initTxS: function () {
            txSeller_db.init(this);
        },
        reEncrypt:function (pwd) {
            connect.send({ Name:"reEncrypt", Payload:{password: pwd, tID: this.selectedTx}}, function (payload, _this) {
                console.log("再加密数据成功", payload);
            }, function (payload, _this) {
                console.log("再加密数据失败：", payload);
                _this.$alert(payload, "再加密数据失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    },
    components: {
        SFT
    },
    computed: {
        listenTxSRefresh() {
            return this.$store.state.transactionsell;
        }
    },
    watch: {
        listenTxSRefresh: function () {
            this.curPage = 1;
            this.total = this.$store.state.transactionsell.length;
        }
    },
    created () {
        this.total = this.$store.state.transactionsell.length;
    }
}
</script>

<style>

</style>
