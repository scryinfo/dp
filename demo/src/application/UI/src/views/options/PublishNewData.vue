<template>
    <section>
        <el-form class="pubForm" :model="pubData" label-width="15%">
            <el-form-item label="Title"><el-input v-model="pubData.Title" clearable></el-input></el-form-item>
            <el-form-item label="Price">
                <el-input v-model.number="pubData.Price" placeholder="Unit is DDD" clearable></el-input></el-form-item>
            <el-form-item label="Keys"><el-input placeholder="Separate each tag with a comma or semicolon" :rows=2
                                                 v-model="pubData.Keys" type="textarea" clearable></el-input></el-form-item>
            <el-form-item label="Description">
                <el-input v-model="pubData.Description" type="textarea" :rows=3 clearable></el-input></el-form-item>
            <el-form-item label="Data"><el-input ref="selectedData" @change="getData" type="file"></el-input></el-form-item>
            <el-form-item label="Proofs">
                <el-input ref="selectedProofs" @change="getProofs" type="file" multiple></el-input></el-form-item>
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
                Title: "",
                Price: 0,
                Keys: "",
                Description: "",
                Data: "",
                Proofs: [],
                Owner: this.$store.state.account,
                Password: ""
            }
        }
    },
    methods: {
        getData: function () {
            this.pubData.Data = ""
            let data = this.$refs.selectedData.$refs.input.files[0]
            let reader = new FileReader()
            let _this = this
            reader.readAsDataURL(data)
            reader.onload = function (evt) {
                _this.pubData.Data = evt.target.result
            }
        },
        getProofs: function () {
            this.pubData.Proofs = []
            let proofs = this.$refs.selectedProofs.$refs.input.files
            let _this = this
            for (let i=0;i<ppf.length;i++) {
                let reader = new FileReader()
                reader.readAsDataURL(proofs[i])
                reader.onload = function (evt) {
                    _this.pubData.Proofs.push(evt.target.result)
                }
            }
        },
        pubPwd: function () {
            this.$prompt(this.$store.state.account, "Input password for this account:", {
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel"
            }).then(({ value }) => {
                this.pubData.Password = value
                this.pub()
            }).catch(() => {
                this.$message({
                    type: "info",
                    message: "Cancel publish."
                })
            })
        },
        pub: function () {
            let _this = this
            let pub = {
                Title: this.pubData.Title,
                Price: this.pubData.Price,
                Keys: this.pubData.Keys,
                Description: this.pubData.Description,
                Owner: this.pubData.Owner
            }
            astilectron.sendMessage({Name:"publish",Payload:this.pubData}, function (message) {
                if (message.name !== "error") {
                    dl_db.write(pub, message.payload)
                    dl_db.init(_this)
                    console.log("Info: Publish new data success.")
                }else {
                    console.log("Node: publish.newData failed. ", message)
                    alert("Publish data failed: ", message.payload)
                }
            })
            // this.resetData(), especially pwd.
        }
    }
}
</script>

<style scoped>
.pubForm {
    padding: 0 20% 0 10%;
}
</style>
