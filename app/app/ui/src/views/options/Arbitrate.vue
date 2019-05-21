<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <div v-if="!showControl">
            <el-col :span="24" class="section-item">
                你还不是验证者！请前往验证界面注册成为验证者。
            </el-col>
        </div>
        <div v-if="showControl">
            <el-col :span="24" class="section-item">
                <el-button size="mini" type="primary" @click="decryptDialog = true">解密数据</el-button>
                <el-button size="mini" type="primary" @click="ArbitrateDialog = true">仲裁数据</el-button></el-col>
            <el-table :data="this.$store.state.transactionarbitrator.slice((curPage-1)*pageSize, curPage*pageSize)"
                      highlight-current-row border :height=height @current-change="currentChange">
                <el-table-column type="expand">
                    <el-form slot-scope="props" label-position="left" class="tx-table-expand">
                        <el-form-item label="标题"><span>{{ props.row.Title }}</span></el-form-item>
                        <el-form-item label="价格"><span>{{ props.row.Price }}</span></el-form-item>
                        <el-form-item label="标签"><span>{{ props.row.Keys }}</span></el-form-item>
                        <el-form-item label="描述"><span>{{ props.row.Description }}</span></el-form-item>
                        <el-form-item label="卖家"><span>{{ props.row.Seller }}</span></el-form-item>
                        <el-form-item label="状态"><span>{{ props.row.State }}</span></el-form-item>
                        <el-form-item label="验证者回复1"><span>{{ props.row.Verifier1Response }}</span></el-form-item>
                        <el-form-item label="验证者回复2"><span>{{ props.row.Verifier2Response }}</span></el-form-item>
                    </el-form>
                </el-table-column>
                <el-table-column prop="Title" label="标题" show-overflow-tooltip></el-table-column>
                <el-table-column prop="Price" label="价格" show-overflow-tooltip></el-table-column>
                <el-table-column prop="State" label="状态" show-overflow-tooltip></el-table-column>
            </el-table>
            <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                           layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
            ></el-pagination>

            <!-- dialogs -->
            <el-dialog :visible.sync="decryptDialog" title="输入密码：">
                <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('decrypt')">取消</el-button>
                    <el-button type="primary" @click="decrypt">确认</el-button>
                </div>
            </el-dialog>
            <el-dialog :visible.sync="ArbitrateDialog" title="仲裁数据：">
                <el-dialog :visible.sync="ArbitrateDialog2" title="输入密码：" append-to-body>
                    <p>{{this.$store.state.account}}</p><el-input v-model="password" show-password clearable></el-input>
                    <div slot="footer">
                        <el-button @click="cancelClickFunc('arbitrate2')">取消</el-button>
                        <el-button type="primary" @click="Arbitrate">确认</el-button>
                    </div>
                </el-dialog>
                <p>{{this.$store.state.account}}</p>
                <div>判断数据真实性：&nbsp;&nbsp;&nbsp;
                    <el-switch v-model="arbitrateResult" active-text="真" inactive-text="假"></el-switch>
                </div>
                <div slot="footer">
                    <el-button @click="cancelClickFunc('arbitrate')">取消</el-button>
                    <el-button type="primary" @click="ArbitrateDialog2 = true">输入密码</el-button>
                </div>
            </el-dialog>
        </div>
    </section>
</template>

<script>
import {acc_db} from "../../DBoptions";
export default {
    name: "Arbitrate.vue",
    data () {
        return {
            showControl: false,
            decryptDialog: false,
            ArbitrateDialog: false,
            ArbitrateDialog2: false,
            selectedTx: {},     // {txID: "", User: "", MetaDataIDEncrypt: ""}
            password: "",
            arbitrateResult: false,
            height: window.innerHeight - 170,
            curPage: 1,
            pageSize: 6,
            total: 0
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize},
        currentChange: function (curRow) {
            this.selectedTx = {
                TransactionID: curRow.TransactionID,
                User: curRow.Buyer,
                MetaDataIDEncrypt: curRow.MetaDataIDEncWithArbitrator,
            }
        },
        cancelClickFunc: function (dialogName) {
            let str = ""
            switch (dialogName) {
                case "decrypt": this.decryptDialog = false; str = "解密数据"; break
                case "arbitrate": this.ArbitrateDialog = false; str = "仲裁数据"; break
                case "arbitrate2": this.ArbitrateDialog2 = false; str = "仲裁数据"; break
            }
            this.$message({
                type: "info",
                message: "取消" + str
            })
        },
        refresh: function () {
            let _this = this
            acc_db.read(this.$store.state.account, function (accInstance) {
                _this.showControl = accInstance.isVerifier
            })
        },
        decrypt: function () {
            this.decryptDialog = false
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({ Name:"decrypt",Payload:{password: pwd, tID: this.selectedTx}}, function (message) {
                if (message.name !== "error") {
                    _this.$alert(message.payload, "原始数据：", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "info"
                    })
                    console.log("解密数据成功", message)
                }else {
                    console.log("解密数据失败：", message.payload)
                    _this.$alert(message.payload, "解密数据失败！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        },
        Arbitrate: function () {
            this.ArbitrateDialog = false
            this.ArbitrateDialog2 = false
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({ Name:"arbitrate",Payload:{password: pwd, tID: this.selectedTx, arbitrate: this.arbitrateResult}}, function (message) {
                if (message.name !== "error") {
                    console.log("仲裁成功", message)
                }else {
                    console.log("仲裁失败：", message.payload)
                    _this.$alert(message.payload, "仲裁失败！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        }
    },
    computed: {
        listenTxARefresh() {
            return this.$store.state.transactionarbitrator
        }
    },
    watch: {
        listenTxARefresh: function () {
            this.curPage = 1
            this.total = this.$store.state.transactionarbitrator.length
        }
    },
    created () {
        this.total = this.$store.state.transactionarbitrator.length
        this.refresh()
    }
}
</script>

<style scoped>

</style>