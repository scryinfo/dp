<template>
    <div>
        <el-row>
            <el-col :span="24" class="top">
                <el-col :span="18">My Astilectron demo</el-col>
                <el-col :span="6">
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
                        <router-view></router-view>
                    </section>
                </el-col>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import {dl_db, tx_db} from "../DBoptions.js"
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
        }
    },
    created() {
        this.acc = this.$route.params.acc
        this.$store.state.account = this.$route.params.acc
        let _this = this
        dl_db.init(this)
        tx_db.init(this)
        document.addEventListener("astilectron-ready", function() {
            utils.listen(_this)
            DBoptions.init(_this)
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
    height: 440px;
}
.el-menu-item {
    background-color: lightgrey;
}
.section {
    padding: 10px 10%;
    height: 480px;
}
.el-form-item {
    width: 100%;
}
</style>
