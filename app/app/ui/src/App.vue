<!--Scry Info.  All rights reserved.-->
<!--license that can be found in the license file.-->

<template>
	<div id="app">
		<router-view></router-view>
	</div>
</template>

<script>
import {connect} from "./utils/connect.js";
export default {
	name: "app",
    created () {
        connect.WSConnect(this);

        if (sessionStorage.getItem("store")) {
            this.$store.replaceState(Object.assign({}, this.$store.state,JSON.parse(sessionStorage.getItem("store"))));
            sessionStorage.removeItem('store');
        }

        window.addEventListener("beforeunload", function (){
            sessionStorage.setItem("store",JSON.stringify(this.$store.state))
        })
	}
}
</script>

<style>
body {
	margin: 0;
	padding: 0;
	font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
	font-size: 16px;
	-webkit-font-smoothing: antialiased;
}
#app {
	position: absolute;
	top: 0;
	bottom: 0;
	width: 100%;
}
</style>
