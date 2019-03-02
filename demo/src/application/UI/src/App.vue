<template>
	<div id="app">
		<router-view></router-view>
	</div>
</template>

<script>
export default {
	name: 'app',
	data () {
		return {

		}
	}
}

</script>

<style lang="scss">
body {
	margin: 0;
	padding: 0;
	background: #1F2D3D;
	font-family: Helvetica Neue, Helvetica, PingFang SC, Hiragino Sans GB, Microsoft YaHei, SimSun, sans-serif;
	font-size: 16px;
	-webkit-font-smoothing: antialiased;
}

#app {
	position: absolute;
	top: 0;
	bottom: 0;
	width: 100%;
}

.el-submenu [class^=fa] {
	vertical-align: baseline;
	margin-right: 10px;
}

.el-menu-item [class^=fa] {
	vertical-align: baseline;
	margin-right: 10px;
}
</style>
