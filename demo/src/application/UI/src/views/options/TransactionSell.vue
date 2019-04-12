<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="reEncryptDialog = true" >再加密数据</el-button>
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

        <!-- dialogs -->
        <el-dialog :visible.sync="reEncryptDialog" title="输入密码：">
            <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
            <div slot="footer">
                <el-button @click="cancelClickFunc">取消</el-button>
                <el-button type="primary" @click="reEncrypt">确认</el-button>
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
            height: window.innerHeight - 170,
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
                message: "取消再加密数据"
            })
        },
        reEncrypt:function () {
            this.reEncryptDialog = false
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({ Name:"reEncrypt",Payload:{password: pwd, tID: this.selectedTx}}, function (message) {
                if (message.name !== "error") {
                    console.log("再加密数据成功", message)
                }else {
                    console.log("再加密数据失败：", message.payload)
                    _this.$alert(message.payload, "再加密数据失败！", {
                        confirmButtonText: "关闭",
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
