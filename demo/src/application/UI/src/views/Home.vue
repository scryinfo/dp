<template>
    <div>
        <el-row>
            <el-col :span="24" class="top">
                <el-col :span="20">My Astilectron demo</el-col>
                <el-col :span="4">
                    <el-dropdown class="top-dropdown" trigger="click">
                        <span>{{acc}}</span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item @click.native="message">Messages</el-dropdown-item>
                            <el-dropdown-item divided @click.native="logout">Logout</el-dropdown-item>
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
                        <div><el-col :span="24"><router-view></router-view></el-col></div>
                    </section>
                </el-col>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import {dl_db, mt_db, DBoptions} from "../DBoptions"
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
        logout: function () {
            this.$confirm("Are you sure to logout?", "Tips:", {
                confirmButtonText: "Yes",
                cancelButtonText: "No",
                type: "warning"
            }).then(() => {
                this.$router.push("/")
            }).catch(() => {
                this.$message({
                    type:"info",
                    message:"keep login.  ^_^ "
                })
            })
        },
        listen: function () {
            let _this = this
            astilectron.onMessage(function(message) {
                switch (message.name) {
                    case "welcome": console.log(message.payload); break
                    case "sdkInit": console.log(message.name + ": " + message.payload); break
                    case "sendMessage":
                        _this.$notify({
                            title: "Notify: ",
                            message: message.payload,
                            position: "top-left"
                        })
                        break
                }
            })
        }
    },
    created() {
        this.acc = this.$route.params.acc
        this.$store.state.account = this.$route.params.acc
        let _this = this
        dl_db.init(this)
        mt_db.init(this)
        document.addEventListener("astilectron-ready", function() {
            _this.listen()
            DBoptions.init(_this)
        })
    }
}
</script>

<style>
.top-dropdown {
    color: lightgrey;
    font-size: 12px;
}
.aside {
    background-color: lightgrey;
    padding-top: 60px;
    height: 440px;
}
.el-menu-item {
    background-color: lightgrey;
}
.section {
    padding: 10px 10%;
    height: 480px;
}
</style>
