<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <section>
        <el-form class="pubForm" :model="pubData" label-position="left" label-width="25%" :style-height="height">
            <el-form-item label="输入密码:"><el-input v-model="password" show-password clearable></el-input></el-form-item>
            <el-form-item label="标题:"><el-input v-model="pubData.details.Title" clearable></el-input></el-form-item>
            <el-form-item label="价格:"><el-input v-model.number="pubData.Price" clearable></el-input></el-form-item>
            <el-form-item label="标签:"><el-input v-model="pubData.details.Keys" clearable></el-input></el-form-item>
            <el-form-item label="描述:">
                <el-input v-model="pubData.details.Description" type="textarea" :rows=2 clearable></el-input>
            </el-form-item>
            <el-form-item label="是否支持验证：">
                <el-switch v-model="SupportVerify" active-text="是" inactive-text="否"></el-switch>
            </el-form-item>
            <el-form-item label="数据:"><input class="el-input__inner" ref="selectedData" type="file"></el-form-item>
            <el-form-item label="证明:"><input class="el-input__inner" ref="selectedProofs" type="file" multiple></el-form-item>
            <el-form-item><el-button type="primary" @click="pubPrepare">Publish</el-button></el-form-item>
        </el-form>
    </section>
</template>

<script>
import {utils} from "../../utils";
export default {
    name: "PublishNewData",
    data () {
        return {
            height: window.innerHeight - 20,
            pubData: {
                details: {
                    Title: "",
                    Keys: "",
                    Description: "",
                    MetaDataExtension: "",
                    ProofDataExtensions: [],
                    Seller: ""
                },
                Price: 0
            },
            SupportVerify: false,
            IDs: {
                metaDataID: "",
                proofDataIDs: [],
                detailsID: "",
            },
            password: "",
            count: 0
        }
    },
    methods: {
        cancelClickFunc: function () {
            this.pubDialog = false;
            this.$message({
                type: "info",
                message: "取消发布新数据"
            });
        },
        pubPrepare: function () {
            this.count = this.$refs.selectedProofs.files.length;
            this.pubData.details.Seller = this.$store.state.account;
            this.setDataID();
            this.setProofIDs();
        },
        setDataID: function () {
            this.IDs.metaDataID = "";
            let _this = this;
            let data = this.$refs.selectedData.files[0];
            this.pubData.details.MetaDataExtension = data.name.slice(data.name.indexOf("."));
            let reader = new FileReader();
            reader.readAsArrayBuffer(data);
            reader.onload = function (evt) {
                utils.ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                    _this.IDs.metaDataID = result[0].hash;
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
        setProofIDs: function () {
            this.IDs.proofDataIDs = [];
            let _this = this;
            let proofs = this.$refs.selectedProofs.files;
            for (let i=0;i<proofs.length;i++) {
                this.pubData.details.ProofDataExtensions.push( proofs[i].name.slice(proofs[i].name.indexOf(".")) );
                let reader = new FileReader();
                reader.readAsArrayBuffer(proofs[i]);
                reader.onload = function (evt) {
                    utils.ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                        _this.IDs.proofDataIDs.push(result[0].hash);
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
        setDetailsID: function () {
            let _this = this;
            utils.ipfs.add(Buffer.from(JSON.stringify(this.pubData.details), "utf-8")).then(function (result) {
                _this.IDs.detailsID = result[0].hash;
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
            utils.send({Name:"publish", Payload: {password: pwd, supportVerify: this.SupportVerify,
                    price: this.pubData.Price, IDs: this.IDs}});
            utils.addCallbackFunc("publish.callback", function (payload, _this) {
                console.log("发布新数据成功", payload);
            });
            utils.addCallbackFunc("publish.callback.error", function (payload, _this) {
                console.log("发布新数据失败：", payload);
                _this.$alert(payload, "发布新数据失败！", {
                    confirmButtonText: "关闭",
                    showClose: false,
                    type: "error"
                });
            });
        }
    },
    watch: {
        count: function () {
            if (this.count === -1) {
                this.setDetailsID();
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
