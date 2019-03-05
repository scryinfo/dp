<template>
    <section>
        <el-col :span="24" style="padding-bottom: 0; background-color: lightgrey;">
            <el-button @click="buy">Buy</el-button>
        </el-col>

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
                selectsDL: [
                    {ID: ""}
                ]
            }
        },
        methods: {
            buy: function () {
                astilectron.sendMessage({ Name:"buy",Payload:{buyer: this.$store.state.account, ids: this.selectsDL} },
                    function (message) {
                        if (message.payload) {
                            // DBoptions.getTransaction();
                            console.log("Buy data succeed.")
                        }else {
                            alert("Buy data failed.")
                        }
                })
            },
            selectedChange: function (sels) {
                this.selectsDL = []
                for (let i=0;i<sels.length;i++) {
                    this.selectsDL.push({ ID: sels[i].ID })
                }
            }
        }
    }
</script>

<style scoped>

</style>
