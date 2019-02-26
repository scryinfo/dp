<template>
    <section>
        <!--工具条，暂时还没有功能-->
        <el-col :span="24" style="padding-bottom: 0; background-color: lightgrey;">
            tool bar.
        </el-col>

        <el-table :data="this.$store.state.mytransaction" highlight-current-row border height=400>
            <el-table-column type="selection" width="50"></el-table-column>
            <el-table-column prop="Title" label="Title" show-overflow-tooltip></el-table-column>
            <el-table-column prop="TransactionID" label="TransactionID" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Seller" label="Seller" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Buyer" label="Buyer" show-overflow-tooltip></el-table-column>
            <el-table-column prop="State" label="State" show-overflow-tooltip></el-table-column>
        </el-table>
    </section>
</template>

<script>
    export default {
        name: "MyTransaction.vue",
        data () {
            return {

            }
        },
    }
</script>

<style scoped>

</style>
