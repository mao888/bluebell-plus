<template>
	<div>
		<el-row>
			<el-col :span="24">
				<div class="grid-content bg-purple-dark footer">
					<ul class="friends-link">
						<li>
							<a href="https://github.com/mao888">胡超Github</a>
						</li>
						<li>
							<a href="https://space.bilibili.com/289144374?spm_id_from=333.1007.0.0">哔哩哔哩</a>
						</li>
						<li>
							<a href="https://wk-blog.vip">☆往事随風☆</a>
						</li>
						<li>
							<a href="https://blog.csdn.net/qq_45696377?spm=1000.2115.3001.5343">CSDN</a>
						</li>

					</ul>
					<h2 class="compony">&copy;·HuChao版权所有</h2>
				</div>
			</el-col>
		</el-row>
	</div>
</template>

<script>
export default {
	name: "Footer",
};
</script>

<style>
* {
	list-style: none;
	text-decoration: none;
}

.footer {
	min-height: 100px;
	margin-top: 20px;
	border-top: 5px solid #74dfff;
	background-color: rgba(43, 42, 42, 1);
}

.friends-link {
	width: 50%;
	height: 50px;
	margin: 0 auto;
	line-height: 50px;
	display: flex;
	justify-content: space-around;
}

.friends-link li {
	width: 25%;
	height: 100%;
	color: #fff;
}

.friends-link li a {
	display: block;
	width: 100%;
	height: 100%;
	color: #fff;
	text-align: center;
}

.friends-link li a:hover {
	color: orange;
	text-decoration: underline;
}

.footer .compony {
	width: 50%;
	height: 50px;
	margin: 0 auto;
	color: #fff;
	line-height: 50px;
	text-align: center;
}</style>