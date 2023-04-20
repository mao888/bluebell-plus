<template>
	<div>
		<el-card>
			<div slot="header" class="clearfix">
				<div class="web-status">🚀 网站状况</div>
			</div>
			<div class="run-time"></div>
		</el-card>
	</div>
</template>

<script>
export default {
	name: 'TimeMeter',
	data() {
		return {
			runTime: new Date("2023-4-14")
		}
	},
	mounted() {
		this.init();
	},
	methods: {
		init() {
			let free = new this.Free();
			// 设置网站运行时间
			free.dateUtils.runTime({
				el: '.run-time',
				timeStamp: this.runTime.getTime(),
				desc: '运行时间：'
			});
		}
	}
}
</script>

<style lang="less" scoped>
.web-status {
	font-weight: 600;
	text-align: center;
}

.run-time {
	background-image: linear-gradient(to right, orange, purple);
	-webkit-background-clip: text;
	color: transparent;
	text-align: center;
}
</style>