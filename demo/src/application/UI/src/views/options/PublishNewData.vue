<template>
    <section>
        <el-form class="pubForm" :model="pubData" label-position="left" label-width="25%">
            <el-form-item label="Password:"><el-input v-model="password" show-password clearable></el-input></el-form-item>
            <el-form-item label="Title:"><el-input v-model="pubData.details.Title" clearable></el-input></el-form-item>
            <el-form-item label="Price:"><el-input v-model.number="pubData.Price" clearable></el-input></el-form-item>
            <el-form-item label="Keys:"><el-input v-model="pubData.details.Keys" clearable></el-input></el-form-item>
            <el-form-item label="Description:">
                <el-input v-model="pubData.details.Description" type="textarea" :rows=2 clearable></el-input>
            </el-form-item>
            <el-form-item label="Support verify:">
                <el-switch v-model="SupportVerify" active-text="Yes" inactive-text="No"></el-switch>
            </el-form-item>
            <el-form-item label="Data:"><el-input ref="selectedData" type="file"></el-input></el-form-item>
            <el-form-item label="Proofs:"><el-input ref="selectedProofs" type="file" multiple></el-input></el-form-item>
            <el-form-item><el-button type="primary" @click="pubPrepare">Publish</el-button></el-form-item>
        </el-form>
    </section>
</template>

<script>
export default {
    name: "PublishNewData",
    data () {
        return {
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
            this.pubDialog = false
            this.$message({
                type: "info",
                message: "Cancel publish. "
            })
        },
        pubPrepare: function () {
            this.count = this.$refs.selectedProofs.$refs.input.files.length
            this.pubData.details.Seller = this.$store.state.account
            this.setDataID()
            this.setProofIDs()
        },
        setDataID: function () {
            let _this = this
            let ipfs = require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'})
            let data = this.$refs.selectedData.$refs.input.files[0]
            this.pubData.details.MetaDataExtension = data.name.slice(data.name.indexOf("."))
            let reader = new FileReader()
            reader.readAsArrayBuffer(data)
            reader.onload = function (evt) {
                ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                    _this.IDs.metaDataID = result[0].hash
                    _this.count--
                }).catch(function (err) {
                    console.log("Node: add.metaData.failed. ", err)
                    _this.$alert(err, "Error: Add meta data failed. ", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                })
            }
        },
        setProofIDs: function () {
            this.IDs.proofDataIDs = []
            let _this = this
            let ipfs = require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'})
            let proofs = this.$refs.selectedProofs.$refs.input.files
            for (let i=0;i<proofs.length;i++) {
                this.pubData.details.ProofDataExtensions.push( proofs[i].name.slice(proofs[i].name.indexOf(".")) )
                let reader = new FileReader()
                reader.readAsArrayBuffer(proofs[i])
                reader.onload = function (evt) {
                    ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                        _this.IDs.proofDataIDs.push(result[0].hash)
                        _this.count--
                    }).catch(function (err) {
                        console.log("Node: add.proofsData.failed. ", err)
                        _this.$alert(err, "Error: Add proofs data failed. ", {
                            confirmButtonText: "I've got it.",
                            showClose: false,
                            type: "error"
                        })
                    })
                }
            }
            return true
        },
        setDetailsID: function () {
            let _this = this
            let ipfs = require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'})
            ipfs.add(Buffer.from(JSON.stringify(this.pubData.details), "utf-8")).then(function (result) {
                _this.IDs.detailsID = result[0].hash
                _this.count--
            }).catch(function (err) {
                console.log("Node: add.detailsData.failed. ", err)
                _this.$alert(err, "Error: Add details data failed. ", {
                    confirmButtonText: "I've got it.",
                    showClose: false,
                    type: "error"
                })
            })
            return true
        },
        pub: function () {
            let _this = this
            let pwd = this.password
            this.password = ""
            astilectron.sendMessage({Name:"publish",Payload: {password: pwd, supportVerify: this.SupportVerify,
                    price: this.pubData.Price, IDs: this.IDs}}, function (message) {
                if (message.name !== "error") {
                    // optimize?: dl_db.write here, seller will see his publish before contract emit event.
                    console.log("Publish new data success.", message)
                }else {
                    console.log("Node: publish.newData failed. ", message.payload)
                    _this.$alert(message.payload, "Error: Publish data failed: ", {
                        confirmButtonText: "I've got it.",
                        showClose: false,
                        type: "error"
                    })
                }
            })
        }
    },
    watch: {
        count: function () {
            if (this.count === -1) {
                this.setDetailsID()
            }
            if (this.count === -2) {
                this.pub()
                this.count = 0
            }
        }
    }
}
</script>

<style>
.pubForm {
    padding: 0 10% 0 10%;
}
</style>
