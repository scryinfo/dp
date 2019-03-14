<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="buyPwd">Buy</el-button>
        </el-col>

        <el-table :data="this.$store.state.datalist" highlight-current-row border height=400 @selection-change="selectedChange">
            <el-table-column type="selection" width=50></el-table-column>
            <el-table-column prop="Title" label="Title" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Price" label="Price" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Keys" label="Keys" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Description" label="Description" show-overflow-tooltip></el-table-column>
            <el-table-column prop="SupportVerify" label="SupportVerify" show-overflow-tooltip></el-table-column>
        </el-table>
    </section>
</template>

<script>
export default {
    name: "DataList",
    data () {
        return {
            selectsDL: []    // {ID: ""}
        }
    },
    methods: {
        selectedChange: function (sels) {
            this.selectsDL = []
            for (let i=0;i<sels.length;i++) {
                this.selectsDL.push( sels[i].PublishID )
            }
        },
        buyPwd: function () {
            this.$prompt(this.$store.state.account, "Input password for this account:", {
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel"
            }).then((pwd) => {
                this.buy(pwd.value)
            }).catch(() => {
                this.$message({
                    type: "info",
                    message: "Cancel pre-buy."
                })
            })
        },
        buy: function (pwd) {
            let _this = this
            // not support buy a group of data one time, give the first id for instead.
            astilectron.sendMessage({ Name:"buy",Payload:{password: pwd, ids: this.selectsDL[0]} }, function (message) {
                    if (message.name !== "error") {
                        _this.selectsDL = []
                        console.log("Buy data success.")
                    }else {
                        console.log("Node: buy failed.", message)
                        alert("Buy data failed.")
                    }
            })
        }
    }
}
</script>

<style>

</style>
