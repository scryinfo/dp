<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-col :span="21" class="section-item">
            <c-f-t button-name="预购买" dialog-title="是否启动验证流程？" @password="advancePurchase">
                <div v-if="selectedData.SupportVerify">
                    <p>是否启动验证流程：</p>
                    <el-switch v-model="startVerify" active-text="是" inactive-text="否"></el-switch>
                </div>
                <div v-if="!selectedData.SupportVerify"><p>卖家不支持验证。</p>点击“输入密码”按钮直接购买数据</div>
            </c-f-t>
        </el-col>
        <el-col :span="3" class="section-item">
            <el-button size="mini" type="primary" @click="initDL">刷新列表</el-button>
        </el-col>

        <el-table :data="this.$store.state.datalist.slice((curPage-1)*pageSize, curPage*pageSize)"
                  highlight-current-row border :height=height @current-change="currentChange">
            <el-table-column prop="PublishId" label="发布ID" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Title" label="标题" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Price" label="价格" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Keys" label="标签" show-overflow-tooltip></el-table-column>
            <el-table-column prop="Description" label="描述" show-overflow-tooltip></el-table-column>
            <el-table-column prop="SVDisplay" label="是否支持验证" show-overflow-tooltip></el-table-column>
        </el-table>
        <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
            layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
        ></el-pagination>
    </section>
</template>

<script>
import {connect} from "../../utils/connect.js";
import {utils} from "../../utils/utils.js";
import CFT from "../templates/complex_function_template.vue"
export default {
    name: "datalist.vue",
    data () {
        return {
            selectedData: {},    // {pId: "", SupportVerify: false, Price: 0}
            curPage: 1,
            pageSize: 6,
            total: 0,
            height: window.innerHeight - 170,
            startVerify: false,
        }
    },
    methods: {
        setCurPage: function (curPageReturn) {this.curPage = curPageReturn;},

        setPageSize: function (newPageSize) {this.pageSize = newPageSize;},

        currentChange: function (curRow) {
            this.selectedData = {
                PublishId: curRow.PublishId,
                SupportVerify: curRow.SupportVerify,
                Price: curRow.Price
            };
        },

        initDL: function () {
            utils.reacquireData("dl");
        },

        advancePurchase: function (pwd) {
            connect.send({Name:"advancePurchase",Payload:{password: pwd, startVerify: this.startVerify,
                        PublishId: this.selectedData.PublishId, price: this.selectedData.Price.toString()}},
                function (payload, _this) {
                console.log("预购买成功", payload);
            }, function (payload, _this) {
                console.log("预购买失败：", payload);
                _this.$alert(payload, "预购买失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    },
    components: {
        CFT
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
    created () {
        this.total = this.$store.state.datalist.length;

        this.initDL();
    }
}
</script>

<style>

</style>
