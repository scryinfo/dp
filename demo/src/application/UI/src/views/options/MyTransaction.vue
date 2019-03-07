<template>
    <section>
        <el-col :span="24" style="padding-bottom: 0; background-color: lightgrey;">
            <el-button size="mini" type="primary" @click="purchasePwd" >Purchase</el-button>
        </el-col>

        <el-table :data="this.$store.state.mytransaction" highlight-current-row border height=400 @selection-change="selectedChange">
            <el-table-column type="selection" width=50></el-table-column>
            <el-table-column type="expand">
                <el-form label-position="left" inline slot-scope="props" label-width="150">
                    <el-form-item label="transactionID"><span>{{ props.row.transactionID }}</span></el-form-item>
                    <el-form-item label="Title"><span>{{ props.row.Title }}</span></el-form-item>
                    <el-form-item label="State"><span>{{ props.row.State }}</span></el-form-item>
                    <el-form-item label="Buyer"><span>{{ props.row.Buyer }}</span></el-form-item>
                    <el-form-item label="Seller"><span>{{ props.row.Seller }}</span></el-form-item>
                    <el-form-item label="Verifier1Response"><span>{{ props.row.Verifier1Response }}</span></el-form-item>
                    <el-form-item label="Verifier2Response"><span>{{ props.row.Verifier2Response }}</span></el-form-item>
                    <el-form-item label="Verifier3Response"><span>{{ props.row.Verifier3Response }}</span></el-form-item>
                    <el-form-item label="ArbitrateResult"><span>{{ props.row.ArbitrateResult }}</span></el-form-item>
                </el-form>
            </el-table-column>
            <el-table-column prop="tID" label="TransactionID" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Title" label="Title" show-overflow-tooltip></el-table-column>
            <el-table-column prop="State" label="State" show-overflow-tooltip></el-table-column>
        </el-table>
    </section>
</template>

<script>
export default {
    name: "MyTransaction",
    data () {
        return {
            selectsMT: []  // {ID: ""}
        }
    },
    methods: {
        selectedChange: function (sels) {
            this.selectsMT = []
            for (let i=0;i<sels.length;i++) {
                this.selectsMT.push({ ID: sels[i].ID })
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
            // not support buy a group of data one time, give the first id for instead.
            astilectron.sendMessage({ Name:"purchase",Payload:{buyer: this.$store.state.account, password: pwd,
                    ids: this.selectsMT[0]} }, function (message) {
                if (message.payload) {
                    console.log("Purchase data succeed.")
                }else {
                    console.log("Node: purchase failed.", message)
                    alert("Purchase data failed.")
                }
            })
        }
    },
    created() {
        this.selectsMT = []
    }
}
</script>

<style scoped>
.el-form-item {
    width: 100%;
}
.el-form-item__label {
    color: #99a9bf;
}
</style>
