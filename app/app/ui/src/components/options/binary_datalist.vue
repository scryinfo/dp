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

<script lang="ts">
import { Component,Vue, Watch } from 'vue-property-decorator'
import {utils} from "../../utils/utils";
import connects from "../../utils/connect";
import CFT from "../templates/complex_function_template.vue"

  @Component({
    components: {
      CFT
    }
  })

  export default class Datalist extends Vue{

    selectedData : any = {} ;    // {pId: "", SupportVerify: false, Price: 0}
    curPage= 1;
    pageSize= 6;
    total= 0;
    height= window.innerHeight - 170;
    startVerify= false;

    setCurPage (curPageReturn:number) {this.curPage = curPageReturn;}

    setPageSize (newPageSize:number) {this.pageSize = newPageSize;}

    currentChange (curRow:any) {
      this.selectedData = {
        PublishId: curRow.PublishId,
        SupportVerify: curRow.SupportVerify,
        Price: curRow.Price
      };
    }

    initDL () {
      utils.reacquireData("dl");
    }

    advancePurchase (pwd:string) {
      connects.send({Name:"advancePurchase",Payload:{password: pwd, startVerify: this.startVerify,
            PublishId: this.selectedData.PublishId, price: this.selectedData.Price.toString()}},
        function (payload:any, _this:any) {
          console.log("预购买成功", payload);
        }, function (payload:any, _this:any) {
          console.log("预购买失败：", payload);
          _this.$alert(payload, "预购买失败！", {
            confirmButtonText: "关闭",
            showClose: false,
            type: "error"
          });
        });
    }

    created () {
      this.total = this.$store.state.datalist.length;

      this.initDL();
    }

    get listenDLRefresh(){
      return this.$store.state.datalist;
    }

    @Watch('this.$store.state.datalist',{immediate:true,deep:true})
    listenDLRefresh1 () {
      this.curPage = 1;
      this.total = this.$store.state.datalist.length;
    }

  }
</script>

<style>

</style>
