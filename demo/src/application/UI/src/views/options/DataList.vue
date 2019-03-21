<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="buyPwd">Buy</el-button>
        </el-col>

        <el-table :data="this.$store.state.datalist.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border height=368 @current-change="currentChange">
            <el-table-column prop="Title" label="Title" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Price" label="Price" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Keys" label="Keys" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Description" label="Description" show-overflow-tooltip></el-table-column>
            <el-table-column prop="SupportVerify" label="SupportVerify" show-overflow-tooltip></el-table-column>
        </el-table>
        <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
            layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
        ></el-pagination>
    </section>
</template>

<script>
export default {
    name: "DataList",
    data () {
        return {
            selectedData: "",    // {pID: ""}
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
            this.selectedData = curRow.PublishID
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
            astilectron.sendMessage({ Name:"buy",Payload:{password: pwd, pID: this.selectedData} }, function (message) {
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
