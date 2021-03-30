<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-col :span="24" class="section-item">
            通知信箱。
        </el-col>
    </section>
</template>

<script>
export default {
    name: "message.vue"
}
</script>

<style>
.section-item {
    margin: 10px 0;
    padding: 10px 30px;
    background-color: lightgrey;
}
.section-item-right {
    float: right;
}
.tx-table-expand label {
    width: 20%;
    color: #99a9bf;
}
.pagination {
    text-align: center;
}
.center {
    display: flex;
    align-items: center;
}
</style>
