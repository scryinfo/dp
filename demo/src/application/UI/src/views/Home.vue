<template>
    <div>
        <el-row>
            <el-col :span="24" class="top">
                <el-col :span="20">My Astilectron demo</el-col>
                <el-col :span="4">
                    <el-dropdown class="top-dropdown">
                        <span>{{acc}}</span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item>Settings</el-dropdown-item>
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
                                              v-if="!item.hidden" class="el-menu-item">{{item.name}}</el-menu-item>
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
import {lfg} from '../listenFromGo'
import {dl_db, mt_db, options} from "../options"
export default {
    name: "Home.vue",
    data () {
        return {
            acc: ""
        }
    },
    methods: {
        logout: function () {
            this.$confirm('Are you sure to logout?', 'Tips:', {
                confirmButtonText: "Yes",
                cancelButtonText: "No",
                type: "warning"
            }).then(() => {
                this.$router.push('/')
            }).catch(() => {
                this.$message({
                    type:"info",
                    message:"keep login. ^_^"
                })
            })
        }
    },
    created() {
        this.acc = this.$route.params.acc
        this.$store.state.account = this.$route.params.acc
        let _this = this
        dl_db.init(this)
        mt_db.init(this)
        document.addEventListener('astilectron-ready', function() {
            lfg.listen()
            options.init(_this)
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
