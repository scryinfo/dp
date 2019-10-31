<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <div>
        <el-row>
            <el-col :span="24" class="top">
                <el-col :span="18">Dapp</el-col>
                <el-col :span="6">
                    <el-dropdown class="top-dropdown" show-timeout=200 trigger="hover">
                        <span>{{ this.$store.state.nickname }}</span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item @click.native="getBalance">余额查询</el-dropdown-item>
                            <el-dropdown-item @click.native="nickName">修改昵称</el-dropdown-item>
                            <el-dropdown-item @click.native="message">消息处理</el-dropdown-item>
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
import {connect} from "../utils/connect.js";
import {utils} from "../utils/utils.js";
export default {
    name: "home.vue",
    data () {
        return {

        }
    },
    methods: {
        getBalance: function () { this.$router.push("/blc"); },

        nickName: function () { this.$router.push("/ncn"); },

        message: function () { this.$router.push("/msg"); },

        logoutMsg: function () {
            this.$confirm("确定退出登录吗？", "提示：", {
                confirmButtonText: "确认",
                cancelButtonText: "取消",
                type: "warning"
            }).then(() => {
                this.logout();
            }).catch(() => {
                this.$message({
                    type:"info",
                    message:"取消退出登录"
                });
            });
        },

        logout: function () {
            connect.send({Name:"logout", Payload: ""}, function (payload, _this) {
                connect.cleanFuncMap();
                setTimeout(function () {
                    _this.$router.push("/");
                }, 500);
            }, function (payload, _this) {
                console.log("退出登录失败：", payload);
                _this.$alert(payload, "退出登录失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    },
    created() {
        utils.init();

        let _home = this;
        _home.$store.state.account = this.$route.params.acc;

        connect.send({Name:"currentUserDataUpdate", Payload: {address: this.$route.params.acc}}, function (payload, _this) {
            console.log("用户信息更新成功！", payload);
            _home.$store.state.nickname = payload;
        }, function (payload, _this) {
            console.log("用户信息更新失败：", payload);
            _this.$alert(payload, "用户信息更新失败！", {
                confirmButtonText: "关闭",
                showClose: false,
                type: "error"
            }).then(() => {
                _this.$router.push("/");
            });
        });
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
    height: calc(100vh - 110px);
}
.el-menu-item {
    background-color: lightgrey;
}
.section {
    padding: 10px 10%;
    height: calc(100vh - 70px);
}
.el-form-item {
    width: 100%;
}
</style>
