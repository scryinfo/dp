<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="reEncryptDialog = true" >ReEncrypt</el-button>
        </el-col>

        <el-table :data="this.$store.state.transactionsell.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border height=468 @current-change="currentChange">
            <el-table-column type="expand">
                <el-form slot-scope="props" label-position="left" class="tx-table-expand">
                    <el-form-item label="TransactionID"><span>{{ props.row.TransactionID }}</span></el-form-item>
                    <el-form-item label="Title"><span>{{ props.row.Title }}</span></el-form-item>
                    <el-form-item label="Price"><span>{{ props.row.Price }}</span></el-form-item>
                    <el-form-item label="State"><span>{{ props.row.State }}</span></el-form-item>
                    <el-form-item label="Buyer"><span>{{ props.row.Buyer }}</span></el-form-item>
                    <el-form-item label="Seller"><span>{{ props.row.Seller }}</span></el-form-item>
                    <el-form-item label="Verifier1Response"><span>{{ props.row.Verifier1Response }}</span></el-form-item>
                    <el-form-item label="Verifier2Response"><span>{{ props.row.Verifier2Response }}</span></el-form-item>
                    <el-form-item label="ArbitrateResult"><span>{{ props.row.ArbitrateResult }}</span></el-form-item>
                </el-form>
            </el-table-column>
            <el-table-column prop="TransactionID" label="TransactionID" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Title" label="Title" show-overflow-tooltip></el-table-column>
            <el-table-column prop="State" label="State" show-overflow-tooltip></el-table-column>
        </el-table>
        <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                       layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
        ></el-pagination>

        <el-dialog :visible.sync="reEncryptDialog" title="Input password for this account:">
            <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
            <div slot="footer">
                <el-button @click="cancelClickFunc">Cancel</el-button>
                <el-button type="primary" @click="reEncrypt">Submit</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
export default {
    name: "TransactionSell",
    data () {
        return {
            selectedTx: {},  // {tID: "", Buyer: "", Seller: "", MetaDataIDEncWithSeller: "", pID: ""}
            curPage: 1,
            pageSize: 6,
            total: 0,
            password: "",
            reEncryptDialog: false
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize},
        currentChange: function (curRow) {
            this.selectedTx = {
                TransactionID: curRow.TransactionID,
                Buyer: curRow.Buyer,
                Seller: curRow.Seller,
                PublishID: curRow.PublishID,
                MetaDataIDEncWithSeller: curRow.MetaDataIDEncWithSeller // transmission between go and js buy not show out to user.
            }
        },
        cancelClickFunc: function () {
            this.reEncryptDialog = false
            this.$message({
                type: "info",
                message: "Cancel re-encrypt. "
            })
        },
        reEncrypt:function () {
            this.reEncryptDialog = false
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({ Name:"reEncrypt",Payload:{password: pwd, tID: this.selectedTx}}, function (message) {
                if (message.name !== "error") {
                    console.log("ReEncrypt data success.", message)
                }else {
                    console.log("Node: reEncrypt failed.", message.payload)
                    _this.$alert(message.payload, "Error: ReEncrypt data failed: ", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        }
    },
    computed: {
        listenTxSRefresh() {
            return this.$store.state.transactionsell
        }
    },
    watch: {
        listenTxSRefresh: function () {
            this.curPage = 1
            this.total = this.$store.state.transactionsell.length
        }
    },
    created () {
        this.total = this.$store.state.transactionsell.length
    }
}
</script>

<style>

</style>
