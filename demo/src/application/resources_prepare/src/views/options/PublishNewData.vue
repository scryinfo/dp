<template>
    <el-form class="pubForm" :model="pubData" label-width="15%">
        <el-form-item label="Title"><el-input v-model="pubData.Title"></el-input></el-form-item>
        <el-form-item label="Price"><el-input v-model.number="pubData.Price" placeholder="Unit is DDD"></el-input></el-form-item>
        <el-form-item label="Keys"><el-input placeholder="Separate each tag with a comma or semicolon"
                                             v-model="pubData.Keys" type="textarea" :rows=2></el-input></el-form-item>
        <el-form-item label="Description">
            <el-input v-model="pubData.Description" type="textarea" :rows=3></el-input></el-form-item>
        <el-form-item label="Data"><el-input ref="seldata" @change="selFile" type="file"></el-input></el-form-item>
        <el-form-item label="Proofs"><el-input ref="selproof" @change="selFiles" type="file" multiple></el-input></el-form-item>
        <el-form-item>
            <el-button type="primary" @click="pub">Publish</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
import {dl_db} from "../../options"
export default {
    name: "PublishNewData",
    data () {
        return {
            pubData: {
                Title: '',
                Price: 0,
                Keys: '',
                Description: '',
                Data: '',
                Proofs: [],
                Owner: this.$store.state.account
            }
        }
    },
    methods: {
        selFile: function () {
            let dp = this.$refs.seldata
            let dpf = dp.$refs.input.files[0]
            let reader = new FileReader()
            let _this = this
            reader.readAsDataURL(dpf)
            reader.onload = function (evt) {
                _this.pubData.Data = evt.target.result
            }
        },
        selFiles: function () {
            let pp = this.$refs.selproof
            let ppf = pp.$refs.input.files
            let _this = this
            for (let i=0;i<ppf.length;i++) {
                let reader = new FileReader()
                reader.readAsDataURL(ppf[i])
                reader.onload = function (evt) {
                    _this.pubData.Proofs.push(evt.target.result)
                }
            }
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
                    console.log(message)
                    dl_db.write(pub, message.payload)
                    dl_db.init(_this)
                    console.log("Log: Publish new data success.")
                }else {
                    alert("Publish data failed: ", message.payload)
                }
            })
            // this.resetData()
        }
    }
}
</script>

<style scoped>
.pubForm {
    padding: 0 20% 0 10%;
}
</style>
