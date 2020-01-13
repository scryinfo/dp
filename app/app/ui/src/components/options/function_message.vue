<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
  <section>
    <el-col :span="18" class="section-item-10">
      <el-input class="txtHeight"
                placeholder="请输入内容"
                v-model="input"
                clearable>
      </el-input>
    </el-col>
    <el-col :span="3" class="section-item-10">
      <el-button size="mini" type="primary" @click="initEvtDL">搜索</el-button>
    </el-col>
    <el-col :span="3" class="section-item-10">
      <el-button size="mini" type="primary" @click="initEvtDL">刷新列表</el-button>
    </el-col>
    <el-table :data="this.$store.state.datalist.slice((curPage-1)*pageSize, curPage*pageSize)"
              highlight-current-row border :height=height @current-change="currentChange">
      <el-table-column min-width="5%" prop="ID" label="ID" show-overflow-tooltip></el-table-column>
      <el-table-column min-width="7%" prop="EventName" label="类型" show-overflow-tooltip></el-table-column>
      <el-table-column min-width="15%" prop="CreatedTime" label="时间" show-overflow-tooltip :formatter="formatterDate"></el-table-column>
      <el-table-column min-width="73%" prop="EventBodys" label="内容" show-overflow-tooltip :formatter="formatterBodys"></el-table-column>    </el-table>
    <el-pagination class="pagination" @current-change="setCurPage" @size-change="setPageSize" :total="total"
                   layout="sizes, total, prev, pager, next, jumper" :page-sizes="[5, 6]" :page-size="pageSize"
    ></el-pagination>
  </section>
</template>

<script lang="ts">
    import { Component,Vue,Watch } from 'vue-property-decorator'
  import CFT1 from "../templates/comfrim_function_template.vue"
  import {utils} from "../../utils/utils";
    import moment from "moment";

    @Component({
        components: {
            CFT1
        }
    })

  export default class Message extends  Vue{
      selectedData= {};    // {pID: "", SupportVerify: false, Price: 0}
      curPage= 1;
      pageSize= 6;
      total= 0;
      height= window.innerHeight - 170;
      startVerify= false;
      input='';

      setCurPage (curPageReturn:number) {this.curPage = curPageReturn;};
      setPageSize (newPageSize:number) {this.pageSize = newPageSize;};
      currentChange (curRow:any) {
          this.selectedData = {
          };
      };
      initEvtDL () {
          let str = this.input;
          utils.reacquireData("evtdl",str);
      };
      formatterDate (row:any,column: { property: string | number; }) {
          // 获取单元格数据
          let data = row[column.property];
          if (data === undefined) {
              return '';
          }
          let dt = new Date(data*1000);
          return moment(dt).format("YYYY-MM-DD HH:mm:ss")
      };
      formatterBodys (row:any,column: { property: string | number; }) {
          let data = row[column.property];
          if (data === undefined){
              return '';
          }
          let s = JSON.parse(data);
          return JSON.stringify(s);
      }

        @Watch('this.$store.state.datalist',{immediate:true,deep:true})
        listenDLRefresh1 () {
            this.curPage = 1;
            this.total = this.$store.state.datalist.length;
        }

        get listenDLRefresh(){
            return this.$store.state.datalist;
        }

        created () {
            this.total = this.$store.state.datalist.length;
            this.initEvtDL();
        }
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
  .txtHeight .el-input__inner {
    height:28px !important;
  }
  /*.el-input__inner{*/
  /*  height:28px !important;*/
  /*}*/
</style>
