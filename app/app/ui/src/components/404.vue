<!-- Scry Info.  All rights reserved.-->
<!-- license that can be found in the license file.-->
<template>
    <p class="page-container">
        <span class="page-text">404 :&nbsp;&nbsp;&nbsp;WEBPAGE NOT FOUND.<br/><br/>
            Redirect to login page after {{ count }} seconds.</span>
    </p>
</template>

<script>
export default {
    name: "404.vue",
    data () {
        return {
            count: 5
        }
    },
    created() {
        let clock = window.setInterval(() => {
            this.count--;
            if (this.count < 0) {
                window.clearInterval(clock);
                this.count = 5;
                this.$router.push({ path: "/"});
            }
        },1000);
    }
}
</script>

<style scoped>
.page-container {
    font-size: 40px;
    color: rgb(192, 204, 218);
    height: calc(100vh - 40px);
    display: flex;
    align-items: center;
    text-align: center;
}
.page-text {
    width: 100%;
}
</style>
