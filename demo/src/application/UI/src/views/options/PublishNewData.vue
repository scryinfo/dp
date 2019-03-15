<template>
    <section>
        <el-form class="pubForm" :model="pubData" label-position="left" label-width="15%">
            <el-form-item label="Title:"><el-input v-model="pubData.details.Title" clearable></el-input></el-form-item>
            <el-form-item label="Price:">
                <el-input v-model.number="pubData.Price" placeholder="Unit is DDD" clearable></el-input></el-form-item>
            <el-form-item label="Keys:"><el-input v-model="pubData.details.Keys" type="textarea" clearable :rows=2></el-input></el-form-item>
            <el-form-item label="Description:">
                <el-input v-model="pubData.details.Description" type="textarea" :rows=3 clearable></el-input></el-form-item>
            <el-form-item label="Data:"><el-input ref="selectedData" type="file"></el-input></el-form-item>
            <el-form-item label="Proofs:"><el-input ref="selectedProofs" type="file" multiple></el-input></el-form-item>
            <el-form-item>
                <el-button type="primary" @click="pubPwd">Publish</el-button>
            </el-form-item>
        </el-form>
    </section>
</template>

<script>
import {dl_db} from "../../DBoptions"
export default {
    name: "PublishNewData",
    data () {
        return {
            pubData: {
                details: {
                    Title: "",
                    Keys: "",
                    Description: "",
                    metaDataExtension: "",
                    proofDataExtensions: []
                },
                Price: 0,
                SupportVerify: false
            },
            send: {
                metaDataID: "",
                proofDataIDs: [],
                detailsID: "",
                price: 0,
                supportVerify: false,  // this.send.SupportVerify = this.pubData.SupportVerify
                password: ""
            },
            count: 0
        }
    },
    methods: {
        pubPwd: function () {
            this.$prompt(this.$store.state.account, "Input password for this account:", {
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel"
            }).then( (pwd) => {
                // unnecessary login.verify?
                this.count = this.$refs.selectedProofs.$refs.input.files.length + 1
                this.send.price = this.pubData.Price
                this.send.supportVerify = this.pubData.SupportVerify
                this.send.password = pwd.value
                this.setIDs()
            }).catch((err) => {
                this.$message({
                    type: "info",
                    message: "Cancel publish. " + err
                })
            })
        },
        setIDs: function() {
            let ipfsAPI = require("ipfs-http-client")
            let ipfs = ipfsAPI({host: 'localhost', port: '5001', protocol: 'http'})
            this.setDataID(ipfs, this)
            this.setProofIDs(ipfs, this)
            this.setDetailsID(ipfs, this)
        },
        setDataID: function (ipfs, _this) {
            let data = this.$refs.selectedData.$refs.input.files[0]
            this.pubData.details.metaDataExtension = data.name.split(".").pop()
            let reader = new FileReader()
            reader.readAsArrayBuffer(data)
            reader.onload = function (evt) {
                ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                    _this.send.metaDataID = result[0].hash
                    _this.count--
                }).catch(function (err) {
                    console.log("Node: add.metaData.failed. ", err)
                    alert("Add meta data failed. ", err)
                })
            }
        },
        setProofIDs: function (ipfs, _this) {
            _this.send.proofDataIDs = []
            let proofs = this.$refs.selectedProofs.$refs.input.files
            for (let i=0;i<proofs.length;i++) {
                this.pubData.details.proofDataExtensions.push( proofs[i].name.split(".").pop() )
                let reader = new FileReader()
                reader.readAsArrayBuffer(proofs[i])
                reader.onload = function (evt) {
                    ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                        _this.send.proofDataIDs.push(result[0].hash)
                        _this.count--
                    }).catch(function (err) {
                        console.log("Node: add.proofsData.failed. ", err)
                        alert("Add proofs data failed. ", err)
                    })
                }
            }
            return true
        },
        setDetailsID: function (ipfs, _this) {
            ipfs.add(Buffer.from(JSON.stringify(this.pubData.details), "utf-8")).then(function (result) {
                _this.send.detailsID = result[0].hash
                _this.count--
            }).catch(function (err) {
                console.log("Node: add.detailsData.failed. ", err)
                alert("Add details data failed. ", err)
            })
            return true
        },
        pub: function () {
            let _this = this
            astilectron.sendMessage({Name:"publish",Payload: this.send}, function (message) {
                if (message.name !== "error") {
                    dl_db.write({
                        Title: _this.pubData.details.Title,
                        Keys: _this.pubData.details.Keys,
                        Description: _this.pubData.details.Description,
                        Price: _this.pubData.Price,
                        Seller: _this.$store.state.account,
                        SupportVerify: _this.pubData.SupportVerify,
                        MetaDataExtension: _this.pubData.details.metaDataExtension,
                        ProofDataExtensions: _this.pubData.details.proofDataExtensions,
                        PublishID: message.payload
                    })
                    dl_db.init(_this)
                    console.log("Publish new data success.")
                    // reset datas.
                }else {
                    console.log("Node: publish.newData failed. ", message)
                    alert("Publish data failed: ", message.payload)
                }
            })
        }
    },
    watch: {
        count: function () {
            if (this.count === -1) {
                this.pub()
                this.count = 0
            }
        }
    }
}
</script>

<style>
.pubForm {
    padding: 0 20% 0 10%;
}
</style>
