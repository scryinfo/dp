<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="buyDialog = true">Buy</el-button>
        </el-col>

        <el-table :data="this.$store.state.datalist.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border height=468 @current-change="currentChange">
            <el-table-column prop="Title" label="Title" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Price" label="Price" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Keys" label="Keys" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Description" label="Description" show-overflow-tooltip></el-table-column>
            <el-table-column prop="SupportVerify" label="SupportVerify" show-overflow-tooltip></el-table-column>
        </el-table>
        <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
            layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
        ></el-pagination>

        <el-dialog :visible.sync="buyDialog" title="Select if you want to start verify process? ">
            <el-dialog :visible.sync="buyDialog2" title="Input password for this account:" append-to-body>
                <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('buy2')">Cancel</el-button>
                    <el-button type="primary" @click="buy">Submit</el-button>
                </div>
            </el-dialog>
            <div v-if="selectedData.SupportVerify">
                <div>Start verify:&nbsp;&nbsp;&nbsp;<el-switch v-model="startVerify" active-text="Verify" inactive-text="Not verify"></el-switch></div>
            </div>
            <div v-if="!selectedData.SupportVerify">
                <p>Seller not support verifiy.<br/>Click "Input password" to buy data without verify or click cancel to cancel.</p>
            </div>
            <div slot="footer">
                <el-button @click="cancelClickFunc('buy')">Cancel</el-button>
                <el-button type="primary" @click="buyDialog2 = true">Input password</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
export default {
    name: "DataList",
    data () {
        return {
            selectedData: {},    // {pID: "", SupportVerify: false}
            curPage: 1,
            pageSize: 6,
            total: 0,
            password: "",
            buyDialog: false,
            buyDialog2: false,
            startVerify: false
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize},
        currentChange: function (curRow) {
            this.selectedData = {
                PublishID: curRow.PublishID,
                SupportVerify: curRow.SupportVerify
            }
        },
        cancelClickFunc: function (dialogName) {
            switch (dialogName) {
            case "buy": this.buyDialog = false; break
            case "buy2": this.buyDialog2 = false; break
            }
            this.$message({
                type: "info",
                message: "Cancel " + dialogName + ". "
            })
        },
        buy: function () {
            this.buyDialog = false
            this.buyDialog2 = false
            let _this = this
            let pwd = this.password
            this.password = ""
            let sv = this.startVerify
            this.startVerify = false
            astilectron.sendMessage({ Name:"buy",Payload:{password: pwd, startVerify: sv, pID: this.selectedData}}, function (message) {
                if (message.name !== "error") {
                    console.log("Buy data success.", message)
                }else {
                    console.log("Node: buy failed.", message.payload)
                    _this.$alert(message.payload, "Error: Buy data failed.", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        }
    },
    computed: {
        listenDLRefresh() {
            return this.$store.state.datalist
        }
    },
    watch: {
        listenDLRefresh: function () {
            this.curPage = 1
            this.total = this.$store.state.datalist.length
        }
    },
    created () {
        this.total = this.$store.state.datalist.length
    }
}
</script>

<style>

</style>
