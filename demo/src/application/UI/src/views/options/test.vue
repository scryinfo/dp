<template>
    <section>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="buyDialog = true">预购买</el-button>
        </el-col>

        <el-table :data="this.$store.state.datalist.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border :height=height @current-change="currentChange">
            <el-table-column prop="Title" label="标题" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Price" label="价格" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Keys" label="标签" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Description" label="描述" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Seller" label="卖家" show-overflow-tooltip></el-table-column>
        </el-table>
        <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                       layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
        ></el-pagination>

        <el-dialog :visible.sync="buyDialog" title="是否启动验证流程？">
            <el-dialog :visible.sync="buyDialog2" title="输入密码：" append-to-body>
                <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('buy2')">取消</el-button>
                    <el-button type="primary" @click="buy">确认</el-button>
                </div>
            </el-dialog>
            <div v-if="selectedData.SupportVerify">
                <div>是否启动验证流程：&nbsp;&nbsp;&nbsp;<el-switch v-model="startVerify" active-text="是" inactive-text="否"></el-switch></div>
            </div>
            <div v-if="!selectedData.SupportVerify">
                <p>卖家不支持验证。<br/>点击“输入密码”按钮直接购买数据</p>
            </div>
            <div slot="footer">
                <el-button @click="cancelClickFunc('buy')">取消</el-button>
                <el-button type="primary" @click="buyDialog2 = true">输入密码</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
export default {
    name: "test.vue",
    data () {
        return {
            selectedData: {},    // {pID: "", SupportVerify: false}
            curPage: 1,
            pageSize: 6,
            total: 0,
            password: "",
            height: window.innerHeight - 170,
            buyDialog: false,
            buyDialog2: false,
            startVerify: false,
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
                message: "取消预购买"
            })
        },
        buy: function () {
            this.buyDialog = false
            this.buyDialog2 = false
            let pwd = this.password
            this.password = ""
            let sv = this.startVerify
            this.startVerify = false
            console.log(pwd, sv, this.height)
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
        },
    },
    created () {
        this.total = this.$store.state.datalist.length
    }
}
</script>

<style scoped>

</style>