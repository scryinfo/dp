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
                    ProofDataExtensions: []
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
                this.count = this.$refs.selectedProofs.$refs.input.files.length
                this.send.price = this.pubData.Price
                this.send.supportVerify = this.pubData.SupportVerify
                this.send.password = pwd.value
                this.setDataID()
                this.setProofIDs()
            }).catch((err) => {
                this.$message({
                    type: "info",
                    message: "Cancel publish. " + err
                })
            })
        },
        setDataID: function () {
            let _this = this
            let ipfs = require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'})
            let data = this.$refs.selectedData.$refs.input.files[0]
            this.pubData.details.MetaDataExtension = data.name.split(".").pop()
            let reader = new FileReader()
            reader.readAsArrayBuffer(data)
            reader.onload = function (evt) {
                ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                    _this.send.metaDataID = result[0].hash
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
            this.send.proofDataIDs = []
            let _this = this
            let ipfs = require("ipfs-http-client")({host: 'localhost', port: '5001', protocol: 'http'})
            let proofs = this.$refs.selectedProofs.$refs.input.files
            for (let i=0;i<proofs.length;i++) {
                this.pubData.details.ProofDataExtensions.push( proofs[i].name.split(".").pop() )
                let reader = new FileReader()
                reader.readAsArrayBuffer(proofs[i])
                reader.onload = function (evt) {
                    ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                        _this.send.proofDataIDs.push(result[0].hash)
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
                _this.send.detailsID = result[0].hash
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
            astilectron.sendMessage({Name:"publish",Payload: this.send}, function (message) {
                if (message.name !== "error") {
                    // dl_db.write here, seller will see his publish before contract emit event.
                    console.log("Publish new data success.", message)
                    // reset datas.
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
    padding: 0 20% 0 10%;
}
</style>
