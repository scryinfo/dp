<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-form class="pubForm" :model="pubData" label-position="left" label-width="25%" :style-height="height">
            <el-form-item label="标题:"><el-input v-model="pubData.details.Title" clearable></el-input></el-form-item>
            <el-form-item label="价格:"><el-input type="number" :v-model.number="pubData.Price" clearable placeholder="uint is DDD"></el-input></el-form-item>
            <el-form-item label="标签:"><el-input v-model="pubData.details.Keys" clearable></el-input></el-form-item>
            <el-form-item label="描述:">
                <el-input v-model="pubData.details.Description" type="textarea" :rows=2 clearable></el-input>
            </el-form-item>
            <el-form-item label="是否支持验证：">
                <el-switch v-model="SupportVerify" active-text="是" inactive-text="否"></el-switch>
            </el-form-item>
            <el-form-item label="数据:"><input class="el-input__inner" ref="selectedData" type="file"></el-form-item>
            <el-form-item label="证明:"><input class="el-input__inner" ref="selectedProofs" type="file" multiple></el-form-item>
            <el-form-item><s-f-t button-name="Publish" button-size="medium" @password="pubPrepare"></s-f-t></el-form-item>
        </el-form>
    </section>
</template>

<script lang="ts">

import { Component,Vue, Watch } from 'vue-property-decorator'
import SFT from "../templates/simple_function_template.vue";
import connects from "../../utils/connect";

@Component({
  components: {
    SFT
  }
})
export default class publish extends Vue{
    height= window.innerHeight - 20;
    pubData= {
      details: {
        Title: "",
        Keys: "",
        Description: "",
        Seller: "",
        MetaDataExtension: "",
        ProofDataExtensions: []
      },
      Price: Number
    };
    SupportVerify= false;
    Ids= {
      metaDataId: "",
      proofDataIds: [],
      detailsId: "",
    };
    password= "";
    count= 0;

  selectedData:any = this.$refs.selectedProofs as Element[];
  selectedProofs:any = this.$refs.selectedProofs as Element[];

  ipfs= require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'});


  pubPrepare (pwd:string) {

    this.password = pwd;
    this.count = this.selectedProofs.length;
    this.pubData.details.Seller = this.$store.state.account;
    this.setDataId();
    this.setProofIds();
  }

  setDataId() {

    this.Ids.metaDataId = "";
    let _this = this;
    let data = this.selectedProofs.files[0];
    this.pubData.details.MetaDataExtension = data.name.slice(data.name.indexOf("."));
    let reader = new FileReader();
    reader.readAsArrayBuffer(data);
    reader.onload = function (evt:any) {
      _this.ipfs.add(Buffer.from(evt.target.result, "utf-8")).then( (result:any) =>{
        _this.Ids.metaDataId = result[0].hash;
        _this.count--;
      }).catch(function (err:any) {
        console.log("IPFS上传失败：", err);
        _this.$alert(err, "IPFS上传失败！", {
          confirmButtonText: "关闭",
          showClose: false,
          type: "error"
        });
      });
    }
  }

  setProofIds () {
    this.Ids.proofDataIds = [];
    this.pubData.details.ProofDataExtensions = [];
    let _this = this;
    let proofs = this.selectedProofs.files;
    for (let i=0;i<proofs.length;i++) {
      let b :never = proofs[i].name.slice(proofs[i].name.indexOf(".")) as never;
      this.pubData.details.ProofDataExtensions.push(b);
      let reader = new FileReader();
      reader.readAsArrayBuffer(proofs[i]);
      reader.onload = (evt)=> {
        if (evt.target!= null && evt.target.result!="") {
          _this.ipfs.add(Buffer.from(evt.target.result as string, "utf-8")).then( (result:any)=> {
            _this.Ids.proofDataIds.push(result[0].hash as never);
            _this.count--;
          }).catch(function (err:any) {
            console.log("IPFS上传失败：", err);
            _this.$alert(err, "IPFS上传失败！", {
              confirmButtonText: "关闭",
              showClose: false,
              type: "error"
            });
          });
        }
      }
    }
  }

  setDetailsId () {
    let _this = this;
    this.ipfs.add(Buffer.from(JSON.stringify(this.pubData.details), "utf-8")).then((result:any)=> {
      _this.Ids.detailsId = result[0].hash;
      _this.count--;
    }).catch(function (err:any) {
      console.log("IPFS上传失败：", err);
      _this.$alert(err, "IPFS上传失败！", {
        confirmButtonText: "关闭",
        showClose: false,
        type: "error"
      });
    });
  }

  pub () {
    let pwd = this.password;
    this.password = "";
    connects.send({Name:"publish", Payload: {password: pwd, supportVerify: this.SupportVerify,
          price: this.pubData.Price.toString(), Ids: this.Ids}},
      function (payload:any, _this:any) {
        console.log("发布新数据成功", payload);
      }, function (payload:any, _this:any) {
        console.log("发布新数据失败：", payload);
        _this.$alert(payload, "发布新数据失败！", {
          confirmButtonText: "关闭",
          showClose: false,
          type: "error"
        });
      });
  }

  @Watch('this.height',{immediate:true,deep:true})
  height1 () {
    this.height = window.innerHeight - 20;
  }

@Watch('this.count',{immediate:true,deep:true})
  count1() {
    if (this.count === -1) {
      this.setDetailsId();
    }
    if (this.count === -2) {
      this.pub();
      this.count = 0;
    }
  }
}

</script>

<style>
.pubForm {
    padding: 0 10% 0 10%;
    height: calc(100vh - 70px);
}
</style>
