<template>
	<div class="side-bar-container">
		<div class="title">
			<el-avatar src="https://z3.ax1x.com/2021/08/31/hdu3a6.jpg"></el-avatar>
			<router-link type="primary" class="leave-message" to="/message">
				<i class="el-icon-chat-dot-round"></i>
				博主留言
			</router-link>
			<el-divider><i class="el-icon-user-solid"></i></el-divider>
		</div>
		<p align="center">
			Visitor count<br>
			<img src="https://profile-counter.glitch.me/mao888/count.svg" />
		</p>
		<el-carousel indicator-position="outside">
			<el-carousel-item v-for="item in imgOptions" :key="item.id">
				<img style="object-fit: cover;width:100%;height: 100%;;" :src="item.path" alt="图片" width="300px" />
			</el-carousel-item>
		</el-carousel>
		<div class="block">
			<h2 class="demonstration">对于博主的评级</h2>
			<el-rate v-model="value2" :colors="colors"> </el-rate>
		</div>
		<div class="message">
			<ul>
				<li>
					<el-tag type="success">
						<i class="el-icon-grape"></i>
						公众号: Gopher毛
					</el-tag>
				</li>
				<li>
					<el-tag type="warning">
						<i class="el-icon-ice-tea"></i>
						qq邮箱: 2557523039@qq.com
					</el-tag>
				</li>
				<li>
					<el-tag>
						<i class="el-icon-watermelon"></i>
						QQ交流群1: 1007576722
					</el-tag>
				</li>
				<li>
					<el-tag>
						<i class="el-icon-watermelon"></i>
						QQ交流群2: 579480724
					</el-tag>
				</li>
				<li>
					<el-tag type="danger">
						<i class="el-icon-cherry"></i>
						电话: 18836288306
					</el-tag>
				</li>
			</ul>
		</div>
		<div class="mine">
			<el-card class="box-card">
				<div slot="header" class="clearfix" style="text-align:center;">
					<span>关于我</span>
				</div>
				<!-- v-for="o in 1" :key="o"  -->
				<div class="text item">
					如果没有奇迹，那就去创造！
				</div>
			</el-card>
		</div>
		<el-calendar class="calendar" v-model="value" width="300px" />
	</div>
</template>
  
<script>
export default {
	name: "SideBar",
	data() {
		return {
			imgOptions: [
				{
					id: 1,
					path: "https://z3.ax1x.com/2021/08/31/hUk4S0.jpg",
				},
				{
					id: 2,
					path: "https://z3.ax1x.com/2021/08/31/hUkXf1.jpg",
				},
				{
					id: 3,
					path: "https://z3.ax1x.com/2021/08/31/hUkxl6.jpg",
				},
				{
					id: 4,
					path: "https://z3.ax1x.com/2021/09/26/46AM0P.jpg",
				},
				{
					id: 5,
					path: "https://z3.ax1x.com/2021/08/31/hUA9mD.jpg",
				},
				{
					id: 6,
					path: "https://z3.ax1x.com/2021/09/26/46AT9e.jpg",
				},
				{
					id: 7,
					path: "https://z3.ax1x.com/2021/08/31/hUAktA.jpg",
				},
				{
					id: 8,
					path: "https://z3.ax1x.com/2021/09/26/46ECcj.jpg",
				}
			],
			value: new Date(),
			value2: null,
			colors: ["#99A9BF", "#F7BA2A", "#FF9900"],
		};
	},
};
</script>
  
<style type="scss" scoped>
.side-bar-container {
	display: flex;
	flex-direction: column;
	font-size: 14px;
}

.title {
	text-align: center;
	margin-top: 10px;
}

.title .el-avatar {
	width: 100px;
	height: 100px;
}

.title .leave-message {
	display: block;
	width: fit-content;
	margin: 0 auto;
	font-size: 16px;
	background-image: linear-gradient(to right, orange, purple);
	-webkit-background-clip: text;
	color: transparent;
	text-align: center;
}

.title .leave-message i {
	font-size: 1.1rem;
}

.title /deep/ .el-avatar img {
	width: 100%;
	height: 100%;
}

.el-carousel {
	width: 100%;
	margin: 0px 0px 10px 0px;
}

.el-carousel img {
	/* width:250px; */
	height: 180px;
}

::v-deep .el-carousel__container {
	height: 200px;
}

.block {
	width: 210px;
	margin: 0 auto;
	text-align: center;
	background-color: rgb(220 169 31 / 29%);
	box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.message {
	display: block;
}

.message ul {
	margin: 20px 0px 20px 0px;
	padding: 0px;
}

.message ul li {
	width: 210px;
	height: 32px;
	list-style: none;
	margin-top: 10px;
	margin-left: 70px;
	margin: 0 auto;
	display: block;
	text-align: center;

	/* box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1) */
}

.message ul li .el-tag {
	width: 210px;
	height: 32px;
}

.mine {
	width: 210px;
	margin: 0 auto;

	.box-card {}
}

.mine span {
	text-align: center;
	font-weight: 600;
	color: #58666e;
}

.mine .box-card {
	margin-bottom: 1rem;
}

.text {
	text-align: center;
}

.el-calendar {
	margin-top: 10px;
}

/deep/ .el-card {
	border: 1px solid #EBEEF5;
	background-color: #fafbfc;
	color: #909399;
	transition: .3s;
}

/deep/ .el-card__header {
	padding: 18px 20px;
	border-bottom: 1px solid #e6a23c;
	box-sizing: border-box;
	/* background: red; */
}

/deep/ .el-calendar__header {
	display: flex;
	justify-content: space-between;
	padding: 12px 20px;
	border-bottom: 1px solid #3a8ee6;
	/* background: #fdf6ec; */
}

/deep/ .el-calendar__body {
	padding: 12px 20px 35px;
	/* background: #fdf6ec; */
}

.calendar {
	margin-bottom: 0.5rem;
}
</style>