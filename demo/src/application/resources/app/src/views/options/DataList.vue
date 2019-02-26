<template>
    <section>
        <!--工具条，暂时还没有功能-->
        <el-col :span="24" style="padding-bottom: 0; background-color: lightgrey;">
            <el-button @click="buy">Buy</el-button>
        </el-col>

        <!--这里暂时会将选中行的全部数据都带出来，后面修改成只带唯一性数据-->
        <el-table :data="this.$store.state.datalist" highlight-current-row border height=400 @selection-change="selectedChange">
            <el-table-column type="selection" width="50"></el-table-column>
            <el-table-column prop="Title" label="Title" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Price" label="Price" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Keys" label="Keys" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Description" label="Description" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Owner" label="Owner" show-overflow-tooltip></el-table-column>
        </el-table>
    </section>
</template>

<script>
    export default {
        name: "DataList.vue",
        data () {
            return {
                selectsDL: []
            }
        },
        methods: {
            buy: function () {
                astilectron.sendMessage({Name:"buy",Payload:{buyer:this.account,ids:this.selectsDL}},
                    function (message) {
                        if (message.payload) {
                            // options.getTransaction();
                            console.log("Buy data succeed.")
                        }else {
                            alert("Buy data failed.")
                        }
                    })
            },
            selectedChange: function (sels) {
                this.selectsDL = sels
            }
        }
    }
</script>

<style scoped>

</style>
