<template>
    <p class="page-container">
        404 :&nbsp;&nbsp;&nbsp;WALLPAGE NOT FOUND.<br/>
        Redirect to login page after {{count}} seconds.
    </p>
</template>

<script>
export default {
    name: "notFound",
    data () {
        return {
            count: 5
        }
    },
    created() {
        let clock = window.setInterval(() => {
            this.count--
            if (this.count < 0) {
                window.clearInterval(clock)
                this.count = 5
                this.$router.push({ path: "/"})
            }
        },1000)
    }
}
</script>

<style scoped>
.page-container {
    font-size: 40px;
    text-align: center;
    color: rgb(192, 204, 218);
    height: 100px;
    line-height: 100px;
    margin-top: 70px;
}
</style>
