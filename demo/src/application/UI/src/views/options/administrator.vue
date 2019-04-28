<template>
    <!-- 管理员操作界面，提供一些便捷功能供测试使用，正式版本删除本文件或删除相关路由设置。 -->
    <section class="administrator">
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" style="width: 200px" @click="welcome">Welcome</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" style="width: 200px" @click="resetChain">ResetChain</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" style="width: 200px" @click="initDL">InitDL</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" style="width: 200px" @click="initTx">InitTx</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" style="width: 200px" @click="testTxDBsConnect">TestTxDBsConnect</el-button></el-col>
    </section>
</template>

<script>
import {dl_db, acc_db, txBuyer_db, txSeller_db, txVerifier_db, db_options} from "../../DBoptions.js";
export default {
    name: "Administrator.vue",
    data () {
        return {

        }
    },
    methods: {
        welcome: function () {
            this.$notify({
                title: "通知: ",
                message: "欢迎使用我的程序！ ^_^ ",
                position: "top-left"
            });
            acc_db.init(this);
        },
        resetChain: function () {
            dl_db.reset();
            acc_db.reset();
            this.resetTxDBs();
            console.log("重置命令已接收");
        },
        initDL: function () {
            dl_db.init(this);
            console.log("数据列表初始化完成");
        },
        initTx: function () {
            txBuyer_db.init(this);
            txSeller_db.init(this);
            txVerifier_db.init(this);
            console.log("交易列表初始化完成");
        },
        testTxDBsConnect: function () {
            let result = "";
            if (txBuyer_db.db_name === txSeller_db.db_name && txSeller_db.db_name === txVerifier_db.db_name) {
                result = "数据库名： " + txBuyer_db.db_name;
            } else {
                result = "数据库名： " + txBuyer_db.db_name + " " + txSeller_db.db_name + " " + txVerifier_db.db_name;
            }
            console.log(result)
        },
        resetTxDBs: function () {
            let c = acc_db.db.transaction(acc_db.db_store_name,"readwrite").objectStore(acc_db.db_store_name).openCursor();
            c.onsuccess = function (evt) {
                let cursor = evt.target.result;
                if (cursor) {
                    db_options.userDBInit(cursor.value.address);
                    db_options.userDBClose();
                    cursor.continue();
                }
            }
        }
    }
}
</script>

<style scoped>
.administrator {
    background-color: lightgrey;
}
</style>