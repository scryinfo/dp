<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-form class="pubForm" :model="pubData" label-position="left" label-width="25%" :style-height="height">
            <el-form-item label="标题:"><el-input v-model="pubData.details.Title" clearable></el-input></el-form-item>
            <el-form-item label="价格:"><el-input type="number" v-model.number="pubData.Price" clearable placeholder="uint is DDD"></el-input></el-form-item>
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

<script>
import {connect} from "../../utils/connect";
import SFT from "../templates/simple_function_template.vue";
export default {
    name: "publish.vue",
    data () {
        return {
            height: window.innerHeight - 20,
            pubData: {
                details: {
                    Title: "",
                    Keys: "",
                    Description: "",
                    Seller: "",
                    MetaDataExtension: "",
                    ProofDataExtensions: []
                },
                Price: Number
            },
            SupportVerify: false,
            Ids: {
                metaDataId: "",
                proofDataIds: [],
                detailsId: "",
            },
            password: "",
            count: 0
        }
    },
    methods: {
        pubPrepare: function (pwd) {
            this.password = pwd;
            this.count = this.$refs.selectedProofs.files.length;
            this.pubData.details.Seller = this.$store.state.account;
            this.setDataId();
            this.setProofIds();
        },

        setDataId: function () {
            this.Ids.metaDataId = "";
            let _this = this;
            let data = this.$refs.selectedData.files[0];
            this.pubData.details.MetaDataExtension = data.name.slice(data.name.indexOf("."));
            let reader = new FileReader();
            reader.readAsArrayBuffer(data);
            reader.onload = function (evt) {
                connect.ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                    _this.Ids.metaDataId = result[0].hash;
                    _this.count--;
                }).catch(function (err) {
                    console.log("IPFS上传失败：", err);
                    _this.$alert(err, "IPFS上传失败！", {
                        confirmButtonText: "关闭",
                        showClose: false,
                        type: "error"
                    });
                });
            }
        },

        setProofIds: function () {
            this.Ids.proofDataIds = [];
            this.pubData.details.ProofDataExtensions = [];
            let _this = this;
            let proofs = this.$refs.selectedProofs.files;
            for (let i=0;i<proofs.length;i++) {
                this.pubData.details.ProofDataExtensions.push(proofs[i].name.slice(proofs[i].name.indexOf(".")));
                let reader = new FileReader();
                reader.readAsArrayBuffer(proofs[i]);
                reader.onload = function (evt) {
                    connect.ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                        _this.Ids.proofDataIds.push(result[0].hash);
                        _this.count--;
                    }).catch(function (err) {
                        console.log("IPFS上传失败：", err);
                        _this.$alert(err, "IPFS上传失败！", {
                            confirmButtonText: "关闭",
                            showClose: false,
                            type: "error"
                        });
                    });
                }
            }
        },

        setDetailsId: function () {
            let _this = this;
            connect.ipfs.add(Buffer.from(JSON.stringify(this.pubData.details), "utf-8")).then(function (result) {
                _this.Ids.detailsId = result[0].hash;
                _this.count--;
            }).catch(function (err) {
                console.log("IPFS上传失败：", err);
                _this.$alert(err, "IPFS上传失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        },

        pub: function () {
            let pwd = this.password;
            this.password = "";
            connect.send({Name:"publish", Payload: {password: pwd, supportVerify: this.SupportVerify,
                    price: this.pubData.Price.toString(), Ids: this.Ids}},
                function (payload, _this) {
                console.log("发布新数据成功", payload);
            }, function (payload, _this) {
                console.log("发布新数据失败：", payload);
                _this.$alert(payload, "发布新数据失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    },
    components: {
        SFT
    },
    watch: {
        count: function () {
            if (this.count === -1) {
                this.setDetailsId();
            }
            if (this.count === -2) {
                this.pub();
                this.count = 0;
            }
        },

        height: function () {
            this.height = window.innerHeight - 20;
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
