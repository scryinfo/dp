<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="danger" @click="cancelDialog = true">Cancel</el-button>
            <el-button size="mini" type="primary" @click="purchaseDialog = true">Purchase</el-button>
            <el-button size="mini" type="primary" @click="decryptDialog = true">Decrypt</el-button>
            <el-button size="mini" type="primary" @click="confirmDialog = true">Confirm</el-button>
            <el-button size="mini" type="primary" @click="creditPre">Credit</el-button>
        </el-col>

        <el-table :data="this.$store.state.transactionbuy.slice((curPage-1)*pageSize, curPage*pageSize)"
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

        <!-- Dialogs -->
        <el-dialog :visible.sync="cancelDialog" title="Input password for this account:">
            <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
            <div slot="footer">
                <el-button @click="cancelClickFunc('cancel')">Cancel</el-button>
                <el-button type="primary" @click="cancelBuying">Submit</el-button>
            </div>
        </el-dialog>
        <el-dialog :visible.sync="purchaseDialog" title="Input password for this account:">
            <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
            <div slot="footer">
                <el-button @click="cancelClickFunc('purchase')">Cancel</el-button>
                <el-button type="primary" @click="purchase">Submit</el-button>
            </div>
        </el-dialog>
        <el-dialog :visible.sync="decryptDialog" title="Input password for this account:">
            <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
            <div slot="footer">
                <el-button @click="cancelClickFunc('decrypt')">Cancel</el-button>
                <el-button type="primary" @click="decrypt">Submit</el-button>
            </div>
        </el-dialog>
        <el-dialog :visible.sync="confirmDialog" title="Confirm the mata data: ">
            <el-dialog :visible.sync="confirmDialog2" title="Input password for this account: ">
                <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('confirm2')">Cancel</el-button>
                    <el-button type="primary" @click="confirm">Submit</el-button>
                </div>
            </el-dialog>
            <div>Confirm the meta data: (Arbitrate process will start if you think it is fake.)&nbsp;&nbsp;&nbsp;
                <el-switch v-model="confirmData" active-text="True" inactive-text="Fake"></el-switch></div>
            <div slot="footer">
                <el-button @click="cancelClickFunc('confirm')">Cancel</el-button>
                <el-button type="primary" @click="confirmDialog2 = true">Input password</el-button>
            </div>
        </el-dialog>
        <el-dialog :visible.sync="creditDialog" title="Credit to verifiers:">
            <el-dialog :visible.sync="creditDialog2" title="Input password for this account:" append-to-body>
                <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('credit2')">Cancel</el-button>
                    <el-button type="primary" @click="credit">Submit</el-button>
                </div>
            </el-dialog>
            <div>Verifier1:
                <el-slider v-model="verifier1Credit" max="5" v-if="verifier1Revert" show-input></el-slider>
                <span v-if="!verifier1Revert">Not support verify or verifier not revert.</span>
            </div>
            <div>Verifier2:
                <el-slider v-model="verifier2Credit" max="5" v-if="verifier2Revert" show-input></el-slider>
                <span v-if="!verifier2Revert">Not support verify or verifier not revert.</span>
            </div>
            <div slot="footer">
                <el-button @click="cancelClickFunc('credit')">Cancel</el-button>
                <el-button type="primary" @click="creditDialog2 = true">Input password</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
export default {
    name: "TransactionBuy",
    data () {
        return {
            selectedTx: {},  // {tID: "", Buyer: "", MetaDataIDEncWithBuyer: "", MetaDataExtension: "", Verifier1: "", Verifier2: ""}
            curPage: 1,
            pageSize: 6,
            total: 0,
            password: "",
            cancelDialog: false,
            purchaseDialog: false,
            decryptDialog: false,
            confirmDialog: false,
            confirmDialog2: false,
            creditDialog: false,
            creditDialog2: false,
            verifier1Revert: false,
            verifier1Credit: 0,
            verifier2Revert: false,
            verifier2Credit: 0,
            confirmData: false
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize},
        currentChange: function (curRow) {
            this.selectedTx = {
                TransactionID: curRow.TransactionID,
                Buyer: curRow.Buyer,
                MetaDataIDEncWithBuyer: curRow.MetaDataIDEncWithBuyer,
                MetaDataExtension: curRow.MetaDataExtension,
                Verifier1: curRow.Verifier1,
                Verifier1Response: curRow.Verifier1Response,
                Verifier2: curRow.Verifier2,
                Verifier2Response: curRow.Verifier2Response
            }
        },
        cancelClickFunc: function (dialogName) {
            switch (dialogName) {
                case "cancel": this.cancelDialog = false; break
                case "purchase": this.purchaseDialog = false; break
                case "decrypt": this.decryptDialog = false; break
                case "confirm": this.cancelDialog = false; break
                case "confirm2": this.confirmDialog2 = false; break
                case "credit": this.creditDialog = false; break
                case "credit2": this.creditDialog2 = false; break
            }
            this.$message({
                type: "info",
                message: "Cancel " + dialogName + ". "
            })
        },
        cancelBuying: function () {
            this.cancelDialog = false
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({ Name:"cancel",Payload:{password: pwd, tID: this.selectedTx}}, function (message) {
                if (message.name !== "error") {
                    console.log("Cancel transaction success.", message)
                }else {
                    console.log("Node: cancel transaction failed.", message.payload)
                    _this.$alert(message.payload, "Error: Cancel transaction failed: ", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        },
        purchase: function () {
            this.purchaseDialog = false
            let _this = this
            let pwd = this.password
            this.password = ""
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
        decrypt: function () {
            this.decryptDialog = false
            let _this = this
            let pwd = this.password
            this.password = ""
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
        confirm: function () {
            this.confirmDialog = false
            this.confirmDialog2 = false
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({ Name:"confirm",Payload:{password: pwd, tID: this.selectedTx,
                    confirmData: true // 'this.startArbitrate' should, but arbitrate not implement.
            }}, function (message) {
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
        },
        creditPre: function () {
            if (this.selectedTx.Verifier1Response !== "") {
                this.verifier1Revert = true
            }
            if (this.selectedTx.Verifier2Response !== "") {
                this.verifier2Revert = true
            }
            this.creditDialog = true
        },
        credit: function () {
            this.creditDialog = false
            this.creditDialog2 = false
            let pwd = this.password
            this.password = ""
            let _this = this
            astilectron.sendMessage({ Name:"credit",Payload:{password: pwd, tID: this.selectedTx, credit: {
                        verifier1Revert: this.verifier1Revert, verifier1Credit: this.verifier1Credit,
                        verifier2Revert: this.verifier2Revert, verifier2Credit: this.verifier2Credit}}}, function (message) {
                if (message.name !== "error") {
                    console.log("Credit data success.", message)
                }else {
                    console.log("Node: credit failed.", message.payload)
                    _this.$alert(message.payload, "Error: Credit data failed: ", {
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
