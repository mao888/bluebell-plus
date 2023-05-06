<template>
	<div class="word-of-day" @click="getWordOfDay" :title="wordOfDay">{{ wordOfDay }}</div>
</template>

<script>
export default {
	name: "DayOfWord",
	data() {
		return {
			wordOfDay: '',
		};
	},
	created() {
		this.getWordOfDay();
	},
	methods: {
		async getWordOfDay() {
			console.log('执行了');
			let response = await this.$axios.get(`https://v.api.aa1.cn/api/yiyan/index.php`);
			const reg = '<p>(.*)</p>';
			this.wordOfDay = response.match(reg)[1];
		},
	}

}
</script>

<style lang="less" scoped>
.word-of-day {
	width: 100%;
	line-height: 32px;
	font-size: 14px;
	background-image: linear-gradient(to right, orange, purple);
	-webkit-background-clip: text;
	color: transparent;
	text-align: center;
	cursor: pointer;
	margin-left: 1rem;
}
</style>