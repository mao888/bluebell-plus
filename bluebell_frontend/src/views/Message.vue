<template>
  <div class="message-container">
	<Comment :sourceId="-1"></Comment>
  </div>
</template>

<script>
import Comment from '../components/Comment.vue'
export default {
	name:"Message",
	components:{ Comment }

}
</script>

<style lang="less" scoped>
.message-container{
	width:90%;
	margin:0 auto;
	margin-top:4rem;
	min-height: 600px;
}
</style>