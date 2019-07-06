<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->

<template>
    <!-- 管理员功能，提供一些便捷功能供测试使用，正式版本当做彩蛋随项目赠送。 -->
    <!-- ps: 彩蛋[cai dan] colorful egg? -> extra scene :) -->
    <section class="administrator">
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="welcome">Welcome</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="resetChain">ResetChain</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="initDL">InitDL</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="initTx">InitTx</el-button></el-col>
        <el-col :span="24" class="section-item">
            <el-button size="mini" type="primary" @click="testTxDBsConnect">TestTxDBsConnect</el-button></el-col>
    </section>
</template>

<script>
import {dl_db, acc_db, txBuyer_db, txSeller_db, txVerifier_db, txArbitrator_db, db_options} from "../../utils/DBoptions.js";
export default {
    name: "ES_administrator.vue",
    data () {
        return {

        }
    },
    methods: {
        welcome: function () {
            this.$notify({
                title: "彩蛋: ",
                message: "谢谢你使用我的程序！ :)  ",
                position: "top-left"
            });
            acc_db.init(this);
        },
        resetChain: async function () {
            dl_db.reset();
            acc_db.reset();
            await this.resetTxDBs();
            console.log("已重置app全部数据");
        },
        initDL: function () {
            dl_db.init(this);
            console.log("数据列表初始化完成");
        },
        initTx: function () {
            txBuyer_db.init(this);
            txSeller_db.init(this);
            txVerifier_db.init(this);
            txArbitrator_db.init(this);
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
                    db_options.deleteDB(cursor.value.address);
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