<!--Scry Info.  All rights reserved.-->
<!--license that can be found in the license file.-->

<template>
	<div id="app">
		<router-view></router-view>
	</div>
</template>

<script lang="ts">
  import Vue from 'vue'
  import Component from 'vue-class-component';
  import connects from './utils/connect'
  @Component({})
  export default class App extends Vue {
    created () {
        connects.WSConnect(this);
      if (sessionStorage.getItem("store")) {
        let store = sessionStorage.getItem("store");
        let s = store!==null?store:"";
        this.$store.replaceState(Object.assign({}, this.$store.state,JSON.parse(s)));
        sessionStorage.removeItem('store');
      }

      window.addEventListener("beforeunload", function (s:any){
        s as App;
        sessionStorage.setItem("store",JSON.stringify(s.$store.state))
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
