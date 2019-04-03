<template>
    <section>
        <div v-if="!showControl">
            <el-col :span="24" class="section-item">
                <el-button size="mini" type="primary" @click="RegisterDialog = true">Register as verifier</el-button>
            </el-col>

            <el-dialog :visible.sync="RegisterDialog" title="Input password for this account:">
                <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('register')">Cancel</el-button>
                    <el-button type="primary" @click="Register">Submit</el-button>
                </div>
            </el-dialog>
        </div>
        <div v-if="showControl">
            <el-col :span="24" class="section-item">
                <el-button size="mini" type="primary" @click="VoteDialog = true">Verify</el-button></el-col>
            <el-table :data="this.$store.state.transactionverifier.slice((curPage-1)*pageSize, curPage*pageSize)"
                      highlight-current-row border height=468 @current-change="currentChange">
                <el-table-column prop="TransactionID" label="TransactionID" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Title" label="Title" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Price" label="Price" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Keys" label="Keys" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Description" label="Description" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Seller" label="Seller" show-overflow-tooltip></el-table-column>
            </el-table>
            <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                           layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
            ></el-pagination>

            <el-dialog :visible.sync="VoteDialog" title="Give out your suggestion: ">
                <el-dialog :visible.sync="VoteDialog2" title="Input password for this account:" append-to-body>
                    <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                    <div slot="footer">
                        <el-button @click="cancelClickFunc('vote2')">Cancel</el-button>
                        <el-button type="primary" @click="Vote">Submit</el-button>
                    </div>
                </el-dialog>
                <p>{{this.$store.state.account}}</p>
                <div>Suggestion:&nbsp;&nbsp;&nbsp;
                    <el-switch v-model="verify.suggestion" active-text="Buy!" inactive-text="Not buy."></el-switch>
                    <el-input v-model="verify.comment" placeholder="comment" clearable></el-input>
                </div>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('vote')">Cancel</el-button>
                    <el-button type="primary" @click="VoteDialog2 = true">Input password</el-button>
                </div>
            </el-dialog>
        </div>
    </section>
</template>

<script>
import {acc_db} from "../../DBoptions";
export default {
    name: "Verify.vue",
    data () {
        return {
            showControl: false,
            RegisterDialog: false,
            VoteDialog: false,
            VoteDialog2: false,
            selectedTx: "",     // txID: ""
            password: "",
            verify: {
                suggestion: false,
                comment: ""
            },
            curPage: 1,
            pageSize: 6,
            total: 0
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize},
        currentChange: function (curRow) {this.selectedTx = curRow.TransactionID},
        cancelClickFunc: function (dialogName) {
            switch (dialogName) {
                case "register": this.RegisterDialog = false; break
                case "vote": this.VoteDialog = false; break
                case "vote2": this.VoteDialog2 = false; break
            }
            this.$message({
                type: "info",
                message: "Cancel " + dialogName + ". "
            })
        },
        refresh: function () {
            let _this = this
            acc_db.read(this.$store.state.account, function (accInstance) {
                _this.showControl = accInstance.isVerifier
            })
        },
        Register: function () {
            this.RegisterDialog = false
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({ Name:"register",Payload:{password: pwd}}, function (message) {
                if (message.name !== "error") {
                    console.log("Register as verifier success.", message)
                }else {
                    console.log("Node: register as verifier failed.", message.payload)
                    _this.$alert(message.payload, "Error: Register as verifier failed.", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        },
        Vote: function () {
            this.VoteDialog = false
            this.VoteDialog2 = false
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({ Name:"verify",Payload:{password: pwd, tID: this.selectedTx, verify: this.verify}}, function (message) {
                if (message.name !== "error") {
                    _this.verify = {suggestion: false, comment: ""}
                    console.log("Vote success.", message)
                }else {
                    console.log("Node: vote failed.", message.payload)
                    _this.$alert(message.payload, "Error: Vote failed.", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        }
    },
    computed: {
        listenTxVRefresh() {
            return this.$store.state.transactionverifier
        }
    },
    watch: {
        listenTxVRefresh: function () {
            this.curPage = 1
            this.total = this.$store.state.transactionverifier.length
        }
    },
    created () {
        this.total = this.$store.state.transactionverifier.length
        this.refresh()
    }
}
</script>

<style scoped>

</style>