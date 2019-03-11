<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="purchasePwd">Purchase</el-button>
            <el-button size="mini" type="primary" @click="cancelMsg">Cancel</el-button>
        </el-col>

        <el-table :data="this.$store.state.transactionbuy" highlight-current-row border height=400 @selection-change="selectedChange">
            <el-table-column type="selection" width=50></el-table-column>
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
    </section>
</template>

<script>
export default {
    name: "TransactionBuy",
    data () {
        return {
            selectsTx: []  // {ID: ""}
        }
    },
    methods: {
        selectedChange: function (sels) {
            this.selectsTx = []
            for (let i=0;i<sels.length;i++) {
                this.selectsTx.push( sels[i].TransactionID )
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
            // not support buy a group of data one time, give the first id for instead.
            astilectron.sendMessage({ Name:"purchase",Payload:{password: pwd, ids: this.selectsTx[0]} }, function (message) {
                if (message.name !== "error") {
                    _this.selectsTx = []
                    console.log("Purchase data success.")
                }else {
                    console.log("Node: purchase failed.", message)
                    alert("Purchase data failed.")
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
            console.log("Node: cancel buying function, which has not realized.")
            // cancel buying and close the transaction, sdk is also blank.
        }
    }
}
</script>

<style>

</style>
