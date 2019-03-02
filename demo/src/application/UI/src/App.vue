<template>
	<div id="app">
		<router-view></router-view>
		<div class="astiloader" id="astiloader">
			<div class="astiloader-background"></div>
			<div class="astiloader-table">
				<div class="astiloader-content">
					<i class="fa fa-spinner fa-spin fa-3x fa-fw"></i>
				</div>
			</div>
		</div>
		<div class="astimodaler" id="astimodaler">
			<div class="astimodaler-background"></div>
			<div class="astimodaler-table">
				<div class="astimodaler-wrapper">
					<div class="astimodaler-body">
						<i class="fa fa-close astimodaler-close" @click="c"></i>
						<div id="astimodaler-content"></div>
					</div>
				</div>
			</div>
		</div>
		<div class="astinotifier" id="astinotifier"></div>
	</div>
</template>

<script>
import {asticode} from './asticode.js'
export default {
	name: 'app',
	data () {
		return {

		}
	},
	methods: {
		c: function () {
			asticode.modaler.close()
		}
	}
}

</script>

<style lang="scss">
body {
	margin: 0;
	padding: 0;
	/*background: url(assets/bg1.jpg) center !important;
		background-size: cover;*/
	// background: #1F2D3D;
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
.astiloader {
	color: #fff;
	display: none;
	height: 100%;
	left: 0;
	position: absolute;
	top: 0;
	width: 100%;
	z-index: 1;
}
.astiloader-background {
	background-color: #000;
	height: 100%;
	left: 0;
	opacity: 0.7;
	position: absolute;
	top: 0;
	width: 100%;
}
.astiloader-table {
	display: table;
	height: 100%;
	left: 0;
	position: absolute;
	top: 0;
	width: 100%;
}
.astiloader-content {
	display: table-cell;
	height: 100%;
	text-align: center;
	vertical-align: middle;
	width: 100%;
}

.astimodaler {
	color: #fff;
	display: none;
	height: 100%;
	left: 0;
	position: fixed;
	top: 0;
	width: 100%;
	z-index: 1;
}
.astimodaler-background {
	background-color: #000;
	height: 100%;
	left: 0;
	opacity: 0.7;
	position: absolute;
	top: 0;
	width: 100%;
}
.astimodaler-table {
	display: table;
	height: 100%;
	left: 0;
	position: absolute;
	top: 0;
	width: 100%;
}
.astimodaler-wrapper {
	display: table-cell;
	height: 100%;
	padding: 2em;
	text-align: center;
	vertical-align: middle;
	width: 100%;
}
.astimodaler-body {
	background-color: #f1f1f1;
	border-radius: 5px;
	color: #333;
	padding: 2em;
	position: relative;
}
.astimodaler-close {
	cursor: pointer;
	position: absolute;
	right: 0.65em;
	top: 0.65em;
}
.astinotifier {
	right: 0;
	position: fixed;
	top: 0;
	width: 100%;
	z-index: 1;
}
@media (min-width: 768px) {
	.astinotifier {
		max-width: 500px;
	}
}
</style>
