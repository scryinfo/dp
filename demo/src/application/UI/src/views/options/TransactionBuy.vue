<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="danger" @click="cancelMsg" plain>Cancel</el-button>
            <el-button size="mini" type="primary" @click="purchasePwd">Purchase</el-button>
            <el-button size="mini" type="primary" @click="decryptPwd">Decrypt</el-button>
            <el-button size="mini" type="primary" @click="confirmPwd">Confirm</el-button>
        </el-col>

        <el-table :data="this.$store.state.transactionbuy.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border height=368 @current-change="currentChange">
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
    </section>
</template>

<script>
export default {
    name: "TransactionBuy",
    data () {
        return {
            selectedTx: {},  // {tID: "", Buyer: "", MetaDataIDEncWithBuyer: "", MetaDataExtension: ""}
            curPage: 1,
            pageSize: 6,
            total: 0
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {
            this.curPage = curPageReturn
        },
        setPageSize: function (newPageSize) {
            this.pageSize = newPageSize
        },
        currentChange: function (curRow) {
            this.selectedTx = {
                TransactionID: curRow.TransactionID,
                Buyer: curRow.Buyer,
                MetaDataIDEncWithBuyer: curRow.MetaDataIDEncWithBuyer,
                MetaDataExtension: curRow.MetaDataExtension
            }
        },
        purchasePwd:function () {
            this.$prompt(this.$store.state.account, "Input password for this account:", {
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel"
            }).then(({ value }) => {
                // login.verify
                this.purchase(value)
            }).catch(() => {
                this.$message({
                    type: "info",
                    message: "Cancel purchase."
                })
            })
        },
        purchase:function (pwd) {
            let _this = this
            astilectron.sendMessage({ Name:"purchase",Payload:{password: pwd, tID: this.selectedTx}}, function (message) {
                if (message.name !== "error") {
                    console.log("Purchase data success.", message)
                }else {
                    console.log("Node: purchase failed.", message.payload)
                    _this.$alert(message.payload, "Error: Purchase data failed: ", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        },
        cancelMsg:function () {
            this.$confirm("Make sure to cancel buying and close the transaction?", "Tips:", {
                confirmButtonText: "Yes",
                cancelButtonText: "No",
                type: "warning"
            }).then(() => {
                this.cancelBuying()
            }).catch(() => {
                this.$message({
                    type:"info",
                    message:"Cancel close."
                })
            })
        },
        cancelBuying:function () {
            console.log("Node: cancel buying has not implemented.")
            // cancel buying and close the transaction, sdk is not finish.
        },
        decryptPwd:function () {
            this.$prompt(this.$store.state.account, "Input password for this account:", {
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel"
            }).then(({ value }) => {
                this.decrypt(value)
            }).catch(() => {
                this.$message({
                    type: "info",
                    message: "Cancel decrypt."
                })
            })
        },
        decrypt:function (pwd) {
            let _this = this
            astilectron.sendMessage({ Name:"decrypt",Payload:{password: pwd, tID: this.selectedTx}}, function (message) {
                if (message.name !== "error") {
                    _this.$alert(message.payload, "Meta data: ", {
                        confirmButtonText: "Close",
                        showClose: false,
                        type: "info"
                    })
                    console.log("Decrypt data success.", message)
                }else {
                    console.log("Node: decrypt failed.", message.payload)
                    _this.$alert(message.payload, "Error: Decrypt data failed: ", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        },
        confirmPwd:function () {
            this.$prompt(this.$store.state.account, "Input password and confirm if the meta data is true?", {
                distinguishCancelAndClose: true, // not implement
                confirmButtonText: "True, close transaction. ",
                cancelButtonText: "Fake, start arbitrate. "
            }).then((pwd) => {
                // think if it is necessary to add another pop box for user can make sure twice?
                this.confirm(pwd.value, true)
            }).catch((pwd) => {
                // arbitrate is not implement, however user confirm the meta data, it will close transaction.
                this.confirm(pwd.value, true)
            })
        },
        confirm:function (pwd, judge) {
            let _this = this
            astilectron.sendMessage({ Name:"confirm",Payload:{password: pwd, tID: this.selectedTx, startArbitrate: judge}},
                function (message) {
                if (message.name !== "error") {
                    console.log("Confirm data success.", message)
                }else {
                    console.log("Node: confirm failed.", message.payload)
                    _this.$alert(message.payload, "Error: Confirm data failed: ", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        }
    },
    computed: {
        listenTxBRefresh() {
            return this.$store.state.transactionbuy
        }
    },
    watch: {
        listenTxBRefresh: function () {
            this.curPage = 1
            this.total = this.$store.state.transactionbuy.length
        }
    },
    created () {
        this.total = this.$store.state.transactionbuy.length
    }
}
</script>

<style>

</style>
