<template>
    <div>
        <el-row>
            <el-col :span="24" class="top">
                <el-col :span="18">一个不会起名字的人写的应用 ￣へ￣</el-col>
                <el-col :span="6">
                    <el-dropdown class="top-dropdown" trigger="click">
                        <span>{{acc}}</span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item @click.native="message">通知</el-dropdown-item>
                            <el-dropdown-item divided @click.native="logoutMsg">退出登录</el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                </el-col>
            </el-col>
            <el-col :span="24" >
                <el-col :span="4">
                    <aside class="aside">
                        <el-menu :default-active="$route.path" unique-opened router>
                            <template v-for="items in $router.options.routes">
                                <el-menu-item v-for="item in items.children" :index="item.path" :key="item.path"
                                              v-if="!item.hidden">{{item.name}}</el-menu-item>
                            </template>
                        </el-menu>
                    </aside>
                </el-col>
                <el-col :span="20">
                    <section class="section">
                        <router-view></router-view>
                    </section>
                </el-col>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import {dl_db, tx_db, acc_db} from "../DBoptions.js"
import {utils} from "../utils.js"
export default {
    name: "Home",
    data () {
        return {
            acc: ""
        }
    },
    methods: {
        message: function () {
            this.$router.push("/msg")
        },
        logoutMsg: function () {
            this.$confirm("确定退出登录吗？", "提示：", {
                confirmButtonText: "确认",
                cancelButtonText: "取消",
                type: "warning"
            }).then(() => {
                this.logout()
            }).catch(() => {
                this.$message({
                    type:"info",
                    message:"取消退出登录"
                })
            })
        },
        logout: function () {
            let _this = this
            astilectron.sendMessage({ Name:"logout", Payload: ""}, function (message) {
                if (message.name !== "error") {
                    _this.$router.push("/")
                }else {
                    console.log("退出登录失败：", message.payload)
                    _this.$alert(message.payload, "退出登录失败！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        }
    },
    created() {
        this.acc = this.$route.params.acc
        this.$store.state.account = this.$route.params.acc
        let _this = this
        dl_db.init(this)
        tx_db.init(this)
        utils.listen(this)
        acc_db.read(this.$route.params.acc, function (accinstance) {
            astilectron.sendMessage({ Name:"sdk.init", Payload: { fromBlock: accinstance.fromBlock } }, function (message) {
                if (message.name !== "error") {
                    console.log("设置初始区块成功", message)
                }else {
                    console.log("设置初始区块失败：", message.payload)
                    _this.$alert(message.payload, "设置初始区块失败！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        })
    }
}
</script>

<style scoped>
.top {
    background-color: grey;
    font-size: 20px;
    color: white;
    height: 50px;
    text-align: center;
    line-height: 50px;
}
.top-dropdown {
    color: lightgrey;
    font-size: 12px;
}
.aside {
    background-color: lightgrey;
    padding-top: 60px;
    height: 540px;
}
.el-menu-item {
    background-color: lightgrey;
}
.section {
    padding: 10px 10%;
    height: 580px;
}
.el-form-item {
    width: 100%;
}
</style>
