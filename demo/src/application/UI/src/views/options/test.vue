<template>
    <div>
        <el-button type="primary" size="mini" @click="add">Add</el-button>
        <el-button type="primary" size="mini" @click="funcTest">Test</el-button>
        <el-table :data="testData.slice((curPage-1)*pageSize,curPage*pageSize)">
            <el-table-column label="ID" prop="ID"></el-table-column>>
            <el-table-column label="Title" prop="Title"></el-table-column>>
            <el-table-column label="Description" prop="Description"></el-table-column>
        </el-table>
        <el-pagination class="ep_css"
                :total="total" layout="sizes, total, prev, pager, next, jumper" @current-change="ccfunc"
                :page-sizes="[1, 2]" :page-size="pageSize" @size-change="scfunc"
        >
        </el-pagination>
    </div>
</template>

<script>
export default {
    name: "test.vue",
    data () {
        return {
            testData: [
                {ID: "1", Title: "title1", Description: "description1"},
                {ID: "2", Title: "title2", Description: "description2"},
                {ID: "3", Title: "title3", Description: "description3"},
                {ID: "4", Title: "title4", Description: "description4"},
                {ID: "5", Title: "title5", Description: "description5"},
                {ID: "6", Title: "title6", Description: "description6"},
                {ID: "7", Title: "title7", Description: "description7"},
                {ID: "8", Title: "title8", Description: "description8"},
                {ID: "9", Title: "title9", Description: "description9"}
            ],
            curPage: 1,
            pageSize: 2,
            total: 0,
            message: "Here is error details. "
        }
    },
    methods: {
        ccfunc: function (curPageReturn) {
            this.curPage = curPageReturn
        },
        scfunc: function (newPageSize) {
            this.pageSize = newPageSize
        },
        add: function () {
            this.testData.push({ID: "add", Title: "titleAdd", Description: "descriptionAdd"})
        },
        funcTest: function () {
            this.$alert(this.message, "Error: Buy data failed.", {
                confirmButtonText: "I've got it.",
                showClose: false,
                type: "error"
            })
        }
    },
    watch: {
        testData: function () {
            this.total = this.testData.length
            this.curPage = 1
        }
    },
    created() {
        this.total = this.testData.length
    }
}
</script>

<style scoped>
.ep_css {
    text-align: center;
}
</style>