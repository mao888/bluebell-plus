<template>
	<div class="content">
		<div class="center">
			<ul class="c-l-list">
				<li class="c-l-item" v-for="post in postList" :key="post.post_id">
					<div class="post">
						<a class="vote">
							<span class="iconfont icon-up" @click="vote(post.post_id, 1)"></span>
						</a>
						<span class="text">{{ post.vote_num }}</span>
						<a class="vote">
							<span class="iconfont icon-down" @click="vote(post.post_id, -1)"></span>
						</a>
					</div>
					<div class="l-container" @click="goDetail(post.post_id)">
						<h4 class="con-title">{{ post.title }}</h4>
						<div class="con-memo">
							<p>{{ post.content }}</p>
						</div>
					</div>
				</li>
				<div class="pagination-block">
					<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="1"
						:page-sizes="[5, 10, 20, 30]" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
						:total="pageTotal.total">
					</el-pagination>
				</div>
			</ul>
		</div>
		<div class="right">
			<div class="topic-info">
				<h5 class="t-header"></h5>
				<div class="t-info">
					<a class="avatar"></a>
					<span class="topic-name">b/{{ community.community_name }}</span>
					<span class="publish" @click="goPublish">发表</span>
				</div>
				<p class="t-desc">{{ community.introduction }}</p>
				<p class="t-create-time">{{ community.create_time }}</p>
				<!-- <button class="topic-btn" @click="goCommunityDetail(community.community_id)">JOIN</button> -->
			</div>
		</div>
	</div>
</template>
  
<script>
import Vue from 'vue';
export default {
	name: "Community",
	data() {
		return {
			postList: [],
			pageTotal: {},
			community: {},
			pageNumber: 1,
			pageSize: 5,
		}
	},
	methods: {
		handleCurrentChange(val) {
			this.pageNumber = val;
			this.getCommunityPostList();
		},
		handleSizeChange(val) {
			this.pageSize = val;
			this.getCommunityPostList();
		},
		getCommunityDetail() {
			this.$axios({
				method: "get",
				url: "/community/" + this.$route.params.id,
			})
				.then(response => {
					console.log(response.data);
					if (response.code == 1000) {
						this.community = response.data;
					} else {
						console.log(response.msg);
					}
				})
				.catch(error => {
					console.log(error);
				});
		},
		getCommunityPostList() {
			this.$axios({
				method: "get",
				url: "/posts2",
				params: {
					community_id: this.$route.params.id,
					page: this.pageNumber,
					size: this.pageSize,
					order: 'score'
				}
			})
				.then(response => {
					console.log(1, response.data);
					if (response.code == 1000) {
						this.postList = response.data.list;
						this.pageTotal = response.data.page;
					} else {
						console.log(response.msg);
					}
				})
				.catch(error => {
					console.log(error);
				});
		},
		vote(post_id, direction) {
			this.$axios({
				method: "post",
				url: "/vote",
				data: {
					post_id: post_id,
					direction: direction,
				}
			})
				.then(response => {
					if (response.code == 1000) {
						console.log("vote success");
						this.getCommunityPostList();
					} else if (response.code == 1009) {
						Vue.prototype.$message.error('请勿重复投票')
					} else {
						console.log(response.msg);
					}
				})
				.catch(error => {
					console.log(error);
				});
		},
		goDetail(id) {
			this.$router.push({ name: "Content", params: { id: id } });
		},
		goPublish() {
			this.$router.push({ name: "Publish" });
		},
	},
	mounted() {
		this.getCommunityDetail();
		this.getCommunityPostList();
	}
};
</script>

<style scoped lang="less">
.content {
	max-width: 100%;
	min-height: 600px;
	box-sizing: border-box;
	display: flex;
	flex-direction: row;
	justify-content: center;
	margin: 0 auto;
	padding: 20px 24px;
	margin-top: 48px;

	.center {
		width: 640px;
		padding-bottom: 10px;
		margin: 0 24px;

		.pagination-block {
			background: #fff;
			padding: 8px;
		}

		.c-l-title {
			font-size: 14px;
			font-weight: 500;
			line-height: 18px;
			color: #1a1a1b;
			text-transform: unset;
			padding-bottom: 10px;
		}

		.c-l-header {
			align-items: center;
			background-color: #ffffff;
			border: 1px solid #ccc;
			border-radius: 4px;
			box-sizing: border-box;
			display: -ms-flexbox;
			display: flex;
			-ms-flex-flow: row nowrap;
			flex-flow: row nowrap;
			height: 56px;
			-ms-flex-pack: start;
			justify-content: flex-start;
			margin-bottom: 16px;
			padding: 0 12px;

			.iconfont {
				margin-right: 4px;
			}

			.btn-iconfont {
				display: flex;
				display: -webkit-flex;
			}

			.active {
				background: #f6f7f8;
				color: #0079d3;
				fill: #0079d3;
				border-radius: 20px;
				height: 32px;
				line-height: 32px;
				margin-right: 8px;
				padding: 0 10px;
			}

			.new {
				font-size: 14px;
				margin-right: 18px;
			}

			.top {
				font-size: 14px;
			}

			.btn-publish {
				height: 32px;
				width: 78%;
				display: flex;
				position: relative;
				border-radius: 4px;

				.word-of-day {
					width: 87%;
					line-height: 32px;
					font-size: 14px;
					overflow: hidden;
					text-overflow: ellipsis;
					white-space: nowrap;
					background-image: linear-gradient(to right, orange, purple);
					-webkit-background-clip: text;
					color: transparent;
					text-align: center;
					cursor: pointer;
					margin-left: 1rem;
				}

				.publish {
					width: 64px;
					height: 100%;
					line-height: 32px;
					background-color: #54b351;
					color: #ffffff;
					border: 1px solid transparent;
					border-radius: 4px;
					box-sizing: border-box;
					text-align: center;
					margin-left: auto;
					cursor: pointer;
					position: absolute;
					right: 0;
				}
			}

			.sort {
				margin-left: 300px;
				display: flex;
				color: #0079d3;
				display: -webkit-flex;
				align-items: center;

				.sort-triangle {
					width: 0;
					height: 0;
					border-top: 5px solid #0079d3;
					border-right: 5px solid transparent;
					border-bottom: 5px solid transparent;
					border-left: 5px solid transparent;
					margin-top: 5px;
					margin-left: 10px;
				}
			}
		}

		.c-l-list {
			.c-l-item {
				list-style: none;
				border-radius: 4px;
				padding-left: 40px;
				cursor: pointer;
				border: 1px solid #ccc;
				margin-bottom: 10px;
				background-color: rgba(255, 255, 255, 0.8);
				color: #878a8c;
				position: relative;

				.post {
					align-items: center;
					box-sizing: border-box;
					display: -ms-flexbox;
					display: flex;
					-ms-flex-direction: column;
					flex-direction: column;
					height: 100%;
					left: 0;
					padding: 8px 4px 8px 0;
					position: absolute;
					top: 0;
					width: 40px;
					border-left: 4px solid transparent;
					background: #f8f9fa;

					.iconfont {
						margin-right: 0;
					}

					.down {
						transform: scaleY(-1);
					}

					.text {
						color: #1a1a1b;
						font-size: 12px;
						font-weight: 700;
						line-height: 16px;
						pointer-events: none;
						word-break: normal;
					}
				}

				.l-container {
					padding: 15px;

					.con-title {
						color: #000000;
						font-size: 18px;
						font-weight: 500;
						line-height: 22px;
						text-decoration: none;
						word-break: break-word;
					}

					.con-memo {
						margin-top: 10px;
						margin-bottom: 10px;
						overflow: hidden;
						text-overflow: ellipsis;
						-webkit-line-clamp: 4;
						display: -webkit-box;
						-webkit-box-orient: vertical;
					}

					.con-cover {
						height: 512px;
						width: 100%;
						background: url("https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1585999647247&di=7e9061211c23e3ed9f0c4375bb3822dc&imgtype=0&src=http%3A%2F%2Fi1.hdslb.com%2Fbfs%2Farchive%2F04d8cda08e170f4a58c18c45a93c539375c22162.jpg") no-repeat;
						background-size: cover;
						margin-top: 10px;
						margin-bottom: 10px;
					}

					.user-btn {
						font-size: 14px;
						display: flex;
						display: -webkit-flex;

						.btn-item {
							display: flex;
							display: -webkit-flex;
							margin-right: 10px;

							.iconfont {
								margin-right: 4px;
							}
						}
					}
				}
			}
		}
	}

	.right {
		flex-grow: 0;
		width: 312px;

		.topic-info {
			width: 100%;
			// padding: 12px;
			cursor: pointer;
			background-color: #ffffff;
			color: #1a1a1b;
			border: 1px solid #cccccc;
			border-radius: 4px;
			overflow: visible;
			word-wrap: break-word;
			padding-bottom: 30px;

			.t-header {
				width: 100%;
				height: 34px;
				background: #0079d3;
			}

			.t-info {
				padding: 0 12px;
				display: flex;
				display: -webkit-flex;
				width: 100%;
				height: 54px;
				align-items: center;

				.avatar {
					width: 54px;
					height: 54px;
					background: url("../assets/images/avatar.png") no-repeat;
					background-size: cover;
					margin-right: 10px;
				}

				.topic-name {
					height: 100%;
					line-height: 54px;
					font-size: 16px;
					font-weight: 500;
				}

				.publish {
					border: 1px solid rgb(31, 148, 237);
					margin-left: auto;
					margin-right: 20px;
					padding: 5px 15px;
					border-color: #0079d3;
					color: #0079d3;
					fill: #0079d3;
					border-radius: 4px;
					font-size: 12px;
				}
			}

			.t-desc,
			.t-create-time {
				font-family: Noto Sans, Arial, sans-serif;
				font-size: 14px;
				line-height: 21px;
				font-weight: 400;
				word-wrap: break-word;
				margin-bottom: 8px;
				padding: 0 12px;
			}

		}
	}
}
</style>