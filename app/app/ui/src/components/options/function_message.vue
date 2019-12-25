<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
  <section>
<!--    <el-col :span="2" class="section-item-10">-->
<!--      <el-button size="mini" type="primary" @click="initEvtDL">标记为已读</el-button>-->
<!--    </el-col>-->
<!--    <el-col :span="2" class="section-item-10">-->
<!--      <el-button size="mini" type="primary" @click="initEvtDL">未读</el-button>-->
<!--    </el-col>-->
<!--    <el-col :span="3" class="section-item-10">-->
<!--      <c-f-t1  button-name="删除" dialog-title="是否确认删除？" @password="initEvtDL">-->
<!--        <div>-->
<!--          <p>是否确认删除：</p>-->
<!--          <el-switch active-text="是" inactive-text="否"></el-switch>-->
<!--        </div>-->
<!--      </c-f-t1>-->
<!--    </el-col>-->
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

<script>
  import moment from 'moment'
  import CFT1 from "../templates/comfrim_function_template.vue"
  import {utils} from "../../utils/utils";

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
        };
      },
      initEvtDL: function () {
        let str = this.input;
        utils.reacquireData("evtdl",str);
      },
      formatterDate:function (row,column) {
        // 获取单元格数据
        let data = row[column.property];
        if (data === undefined) {
          return '';
        }
        let dt = new Date(data*1000);
        return moment(dt).format("YYYY-MM-DD HH:mm:ss")
      },
      formatterBodys:function (row,column) {
        let data = row[column.property];
        if (data === undefined){
          return '';
        }
        let s = JSON.parse(data);
        return JSON.stringify(s);
      }
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
