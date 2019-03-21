<template>
    <section>
        <el-col :span="24" class="section-item">
            Here will show all messages, should I save it in database?
        </el-col>
    </section>
</template>

<script>
export default {
    name: "Message"
}
</script>

<style>
.section-item {
    margin: 10px 0;
    padding: 10px;
    background-color: lightgrey;
}
.tx-table-expand label {
    width: 20%;
    color: #99a9bf;
}
.pagination {
    text-align: center;
}
</style>
