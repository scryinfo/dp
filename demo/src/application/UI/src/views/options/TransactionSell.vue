<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="reEncryptPwd" >ReEncrypt</el-button>
        </el-col>

        <el-table :data="this.$store.state.transactionsell" highlight-current-row border height=400 @selection-change="currentChange">
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
    name: "TransactionSell",
    data () {
        return {
            selectedTx: {}  // {tID: "", Buyer: "", Seller: "", MetaDataIDEncWithSeller: "", pID: ""}
        }
    },
    methods: {
        currentChange: function (curRow) {
            this.selectedTx = {
                ID: curRow.TransactionID,
                Buyer: curRow.Buyer,
                Seller: curRow.Seller,
                PublishID: curRow.PublishID,
                MetaDataIDEncWithSeller: curRow.MetaDataIDEncWithSeller // transmission between go and js buy not show out to user.
            }
        },
        reEncryptPwd:function () {
            this.$prompt(this.$store.state.account, "Input password for this account:", {
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel"
            }).then(({ value }) => {
                this.reEncrypt(value)
            }).catch(() => {
                this.$message({
                    type: "info",
                    message: "Cancel reEncrypt."
                })
            })
        },
        reEncrypt:function (pwd) {
            // not support buy a group of data one time, give the first id for instead.
            astilectron.sendMessage({ Name:"reEncrypt",Payload:{password: pwd, tID: this.selectedTx}}, function (message) {
                if (message.name !== "error") {
                    console.log("ReEncrypt data success.", message)
                }else {
                    console.log("Node: reEncrypt failed.", message)
                    alert("ReEncrypt data failed.")
                }
            })
        }
    }
}
</script>

<style>

</style>
