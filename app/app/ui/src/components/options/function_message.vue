<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-col :span="2" class="section-item-10">
            <el-button size="mini" type="primary" @click="initDL">标记为已读</el-button>
        </el-col>
        <el-col :span="2" class="section-item-10">
            <el-button size="mini" type="primary" @click="initDL">未读</el-button>
        </el-col>
        <el-col :span="3" class="section-item-10">
            <c-f-t1  button-name="删除" dialog-title="是否确认删除？" @password="delEvent">
                <div>
                    <p>是否确认删除：</p>
                    <el-switch active-text="是" inactive-text="否"></el-switch>
                </div>
            </c-f-t1>
        </el-col>
        <el-col :span="11" class="section-item-10">
            <el-input class="txtHeight"
                    placeholder="请输入内容"
                    v-model="input"
                    clearable>
            </el-input>
        </el-col>
        <el-col :span="3" class="section-item-10">
            <el-button  size="mini" type="primary" @click="initDL">搜索</el-button>
            </el-col>
        <el-col :span="3" class="section-item-10">
            <el-button size="mini" type="primary" @click="initDL">刷新列表</el-button>
        </el-col>
        <el-table :data="this.$store.state.datalist.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border :height=height @current-change="currentChange">
            <el-table-column prop="MsgID" label="时间ID" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Title" label="标题" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Price" label="内容" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Keys" label="发布人" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Description" label="状态" show-overflow-tooltip></el-table-column>
            <el-table-column prop="SVDisplay" label="时间" show-overflow-tooltip></el-table-column>
        </el-table>
        <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                       layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
        ></el-pagination>
    </section>
</template>

<script>
    import {connect} from "../../utils/connect";
    import {dl_db} from "../../utils/DBoptions";
    import CFT1 from "../templates/comfrim_function_template.vue"
export default {
    name: "message.vue",
    data () {
        return {
            selectedData: {},    // {pID: "", SupportVerify: false, Price: 0}
            curPage: 1,
            pageSize: 6,
            total: 0,
            height: window.innerHeight - 170,
            startVerify: false,
            input:'',
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn;},
        setPageSize: function (newPageSize) {this.pageSize = newPageSize;},
        currentChange: function (curRow) {
            this.selectedData = {
                PublishID: curRow.PublishID,
                SupportVerify: curRow.SupportVerify,
                Price: curRow.Price
            };
        },
    },
    computed: {
        listenDLRefresh() {
            return this.$store.state.datalist;
        }
    },
    watch: {
        listenDLRefresh: function () {
            this.curPage = 1;
            this.total = this.$store.state.datalist.length;
        }
    },
    components: {
        CFT1
    },
}
</script>

<style>
.section-item-10 {
    margin: 10px 0;
    padding: 10px  10px 10px 10px;
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
    .el-input__inner{
        height:28px !important;
    }
</style>
