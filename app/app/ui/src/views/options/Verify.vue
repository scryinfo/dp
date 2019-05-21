<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <div v-if="!showControl">
            <el-col :span="24" class="section-item">
                <el-button size="mini" type="primary" @click="RegisterDialog = true">注册成为验证者</el-button>
            </el-col>

            <el-dialog :visible.sync="RegisterDialog" title="输入密码：">
                <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('register')">取消</el-button>
                    <el-button type="primary" @click="Register">确认</el-button>
                </div>
            </el-dialog>
        </div>
        <div v-if="showControl">
            <el-col :span="20" class="section-item">
                <el-button size="mini" type="primary" @click="VoteDialog = true">验证数据</el-button>
            </el-col>
            <el-col :span="4" class="section-item">
                <el-button size="mini" type="primary" @click="initTxV">刷新列表</el-button>
            </el-col>
            <el-table :data="this.$store.state.transactionverifier.slice((curPage-1)*pageSize, curPage*pageSize)"
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

            <!-- dialogs -->
            <el-dialog :visible.sync="VoteDialog" title="验证数据：">
                <el-dialog :visible.sync="VoteDialog2" title="输入密码：" append-to-body>
                    <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                    <div slot="footer">
                        <el-button @click="cancelClickFunc('vote2')">取消</el-button>
                        <el-button type="primary" @click="Vote">确认</el-button>
                    </div>
                </el-dialog>
                <p>{{this.$store.state.account}}</p>
                <div>是否建议购买：&nbsp;&nbsp;&nbsp;
                    <el-switch v-model="verify.suggestion" active-text="是" inactive-text="否"></el-switch>
                    <el-input v-model="verify.comment" placeholder="评论" clearable></el-input>
                </div>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('vote')">取消</el-button>
                    <el-button type="primary" @click="VoteDialog2 = true">输入密码</el-button>
                </div>
            </el-dialog>
        </div>
    </section>
</template>

<script>
import {utils} from "../../utils";
import {acc_db, txVerifier_db} from "../../DBoptions";
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
            height: window.innerHeight - 170,
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
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn;},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize;},
        currentChange: function (curRow) {this.selectedTx = curRow.TransactionID;},
        cancelClickFunc: function (dialogName) {
            let str = "";
            switch (dialogName) {
                case "register": this.RegisterDialog = false; str = "注册成为验证者"; break;
                case "vote": this.VoteDialog = false; str = "验证数据"; break;
                case "vote2": this.VoteDialog2 = false; str = "验证数据"; break;
            }
            this.$message({
                type: "info",
                message: "取消" + str
            });
        },
        initTxV: function () {
            txVerifier_db.init(this);
        },
        refresh: function () {
            let _this = this;
            acc_db.read(this.$store.state.account, function (accInstance) {
                _this.showControl = accInstance.isVerifier;
            });
        },
        Register: function () {
            this.RegisterDialog = false;
            let pwd = this.password;
            this.password = "";
            utils.send({Name:"register", Payload:{password: pwd}});
            utils.addCallbackFunc("register.callback", function (payload, _this) {
                console.log("注册成为验证者成功");
            });
            utils.addCallbackFunc("register.callback.error", function (payload, _this) {
                console.log("注册成为验证者失败：", payload);
                _this.$alert(payload, "注册成为验证者失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },
        Vote: function () {
            this.VoteDialog = false;
            this.VoteDialog2 = false;
            let pwd = this.password;
            this.password = "";
            utils.send({Name:"verify", Payload:{password: pwd, tID: this.selectedTx, verify: this.verify}});
            utils.addCallbackFunc("verify.callback", function (payload, _this) {
                _this.verify = {suggestion: false, comment: ""};
                console.log("验证成功", payload);
            });
            utils.addCallbackFunc("verify.callback.error", function (payload, _this) {
                console.log("验证失败：", payload);
                _this.$alert(payload, "验证失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    },
    computed: {
        listenTxVRefresh() {
            return this.$store.state.transactionverifier;
        }
    },
    watch: {
        listenTxVRefresh: function () {
            this.curPage = 1;
            this.total = this.$store.state.transactionverifier.length;
        }
    },
    created () {
        this.total = this.$store.state.transactionverifier.length;
        this.refresh();
    }
}
</script>

<style scoped>

</style>