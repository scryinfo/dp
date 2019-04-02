<template>
    <section>
        <el-button size="mini" type="primary" @click="creditDialog = true">Open dialog</el-button>
        <el-dialog :visible.sync="creditDialog" title="Credit to verifiers:">
            <el-dialog :visible.sync="creditDialog2" title="Input password for this account:" append-to-body>
                <p>1234</p>
                <div slot="footer">
                    <el-button @click="creditDialog2 = false">Cancel</el-button>
                    <el-button type="primary" @click="funcTest">Submit</el-button>
                </div>
            </el-dialog>
            <div>Verifier1:
                <el-slider v-model="verifier1Credit" max="5" v-if="verifier1Revert" show-input></el-slider>
                <span v-if="!verifier1Revert">Not support verify or verifier not revert.</span>
            </div>
            <div>Verifier2:
                <el-slider v-model="verifier2Credit" max="5" v-if="verifier2Revert" show-input></el-slider>
                <span v-if="!verifier2Revert">Not support verify or verifier not revert.</span>
            </div>
            <div slot="footer">
                <el-button size="mini" type="primary" @click="creditDialog2 = true">Open inner dialog</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
export default {
    name: "test.vue",
    data () {
        return {
            creditDialog: false,
            creditDialog2: false,
            verifier1Revert: true,
            verifier1Credit: 3,
            verifier2Revert: false,
            verifier2Credit: 3
        }
    },
    methods: {
        funcTest:function () {
            console.log(this.verifier1Revert, this.verifier1Credit)
            console.log(this.verifier2Revert, this.verifier2Credit)
        }
    }
}
</script>

<style scoped>

</style>