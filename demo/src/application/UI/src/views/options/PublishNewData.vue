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
                SupportVerify: false,
            },
            send: {
                metaDataID: "",
                proofDataIDs: [],
                detailsID: "",
                price: 0,
                supportVerify: false,  // this.send.SupportVerify = this.pubData.SupportVerify
                password: ""
            }
        }
    },
    methods: {
        pubPwd: function () {
            this.$prompt(this.$store.state.account, "Input password for this account:", {
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel"
            }).then( ({ value }) => {
                // login.verify
                this.send.price = this.pubData.Price
                this.send.supportVerify = this.pubData.SupportVerify
                this.send.password = value
                this.getIDs()   // return this.getIDs()
            }).then( () => {    // param: ok
                // if (ok) {
                //     // if it is necessary to judge it?
                // }
                this.pub()
            }).catch((err) => {
                this.$message({
                    type: "info",
                    message: "Cancel publish. " + err
                })
            })
        },
        pub: function () {
            let _this = this
            let pub = {
                Title: this.pubData.details.Title,
                Price: this.pubData.Price,
                Keys: this.pubData.details.Keys,
                Description: this.pubData.details.Description,
                SupportVerify: this.pubData.SupportVerify
            }
            astilectron.sendMessage({Name:"publish",Payload:this.send}, function (message) {
                if (message.name !== "error") {
                    dl_db.write(pub, message.payload)
                    dl_db.init(_this)
                    console.log("Info: Publish new data success.")
                }else {
                    console.log("Node: publish.newData failed. ", message)
                    alert("Publish data failed: ", message.payload)
                }
            })
            // this.resetData(), especially pwd and some arrays.
        },
        getIDs: function () {
            const ipfsAPI = require("ipfs-api")
            const ipfs = ipfsAPI("/ip4/127.0.0.1/tcp/5001")     // send message to go and listen response.
            let ok1 = this.uploadData(ipfs, this)
            let ok2 = this.uploadProofs(ipfs, this)
            let ok3 = this.uploadDetails(ipfs, this)
            return (ok1 && ok2 && ok3)
        },
        uploadData: function (ipfs, _this) {
            let data = this.$refs.selectedData.$refs.input.files[0]
            this.pubData.details.metaDataExtension = data.name.split(".").pop()
            let reader = new FileReader()
            let ok = true
            reader.readAsArrayBuffer(data)
            reader.onload = function (evt) {
                ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                    _this.send.metaDataID = result[0].hash
                }).catch(function (err) {
                    ok = false
                    console.log("Node: add.metaData.failed. ", err)
                    alert("Add meta data failed. ", err)
                })
            }
            ok = (this.pubData.details.metaDataExtension !== "") && ok
            return ok
        },
        uploadProofs: function (ipfs, _this) {
            let proofs = this.$refs.selectedProofs.$refs.input.files
            let ok = true
            for (let i=0;i<proofs.length;i++) {
                this.pubData.details.proofDataExtensions.push( proofs[i].name.split(".").pop() )
                let reader = new FileReader()
                reader.readAsArrayBuffer(proofs[i])
                reader.onload = function (evt) {
                    ipfs.add(Buffer.from(evt.target.result, "utf-8")).then(function (result) {
                        _this.send.proofDataIDs.push( result[0].hash )
                    }).catch(function (err) {
                        ok = false
                        console.log("Node: add.proofsData.failed. ", err)
                        alert("Add proofs data failed. ", err)
                    })
                }
            }
            return ok
        },
        uploadDetails: function (ipfs, _this) {
            let ok = true
            ipfs.add(Buffer.from(JSON.stringify(this.pubData.details), "utf-8")).then(function (result) {
                _this.send.detailsID = result[0].hash
            }).catch(function (err) {
                ok = false
                console.log("Node: add.detailsData.failed. ", err)
                alert("Add details data failed. ", err)
            })
            return ok
        }
    }
}
</script>

<style scoped>
.pubForm {
    padding: 0 20% 0 10%;
}
</style>
