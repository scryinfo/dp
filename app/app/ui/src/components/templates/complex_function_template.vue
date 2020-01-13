<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <!-- complex function template has a two-layer nested dialog. -->
    <div class="inALine">
        <el-button :size="buttonSize" :type="buttonType" @click="pre(preParams)" :disabled="buttonDisabled">{{ buttonName }}</el-button>

        <el-dialog :visible.sync="dialog" :title="dialogTitle">
            <el-dialog :visible.sync="dialogInner" title="输入密码：" append-to-body>
                <p>{{this.$store.state.account}}</p>
                <el-input v-model="password" show-password clearable></el-input>
                <div slot="footer">
                    <el-button @click="cancelFunc('输入密码')">取消</el-button>
                    <el-button type="primary" @click="submitFunc()">确认</el-button>
                </div>
            </el-dialog>

            <slot></slot>
            <div slot="footer">
                <el-button @click="cancelFunc(buttonName)">取消</el-button>
                <el-button type="primary" @click="dialogInner = true">输入密码</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script lang="ts">
  import { Vue, Component, Prop } from 'vue-property-decorator'

  @Component
  export default class CFT extends Vue{
    @Prop(String) buttonSize: String | undefined  ;
    @Prop() buttonType: String | undefined ;
    @Prop() buttonName: String | undefined;
    @Prop()dialogTitle: String | undefined;
    @Prop()buttonDisabled: Boolean | undefined;
    @Prop(Array) preParams: Array<any> | undefined;


    dialog= false;
    dialogInner= false;
    password= "";

    pre (params:any) {
      this.dialog = true;
      if (!!params) {
        let result = [];
        for (let i = 0; i < params.length; i++) {
          result.push(params[i])
        }
        this.$emit("pre", result)
      }
    }

    cancelFunc (name:string) {
      if (name === "输入密码") {
        this.dialogInner = false;
      } else {
        this.dialog = false;
      }
      this.$message({
        type: "info",
        message: "取消" + name
      });
    }

    submitFunc () {
      this.dialogInner = false;
      this.dialog = false;
      let pwd = this.password;
      this.password = "";
      this.$emit("password", pwd);
    }
  }
</script>

<style scoped>
.inALine {
    float: left;
    margin: 0 5px;
}
</style>
