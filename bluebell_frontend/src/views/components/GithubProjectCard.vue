<template>
	<el-row>
		<el-col :span="24">
			<el-card :body-style="{ padding: 0, border: 'none' }" class="box-card">
				<h2 class="title">{{ title }}</h2>
				<ul class="github-hot-project-list">
					<li class="github-hot-project-item" v-for="project in projectList" :key="project.owner.id">
						<div class="user-info">
							<div class="avatar">
								<el-avatar size="small" :src="project.owner.avatar_url"></el-avatar>
							</div>
							<div class="project-name">
								<el-link type="primary" :href="project.html_url">{{ project.full_name }}</el-link>
							</div>
						</div>
						<div class="introduction">{{ project.description }}</div>
						<div class="meta">
							<span class="starts">
								<i class="el-icon-star-on"></i>
								<span>{{ handleNumber(project.stargazers_count) }}</span>
							</span>
							<span class="forks">
								<i class="el-icon-share"></i>
								<span>{{ handleNumber(project.forks_count) }}</span>
							</span>
							<span class="languages" v-if="project.language">
								<i class="el-icon-s-help" :style="{ color: handleTagColor(project.language) }"></i>
								<span>{{ project.language }}</span>
							</span>
						</div>
						<el-divider></el-divider>
					</li>
					<div class="next-btn" v-if="projectList.length > 0" @click="getMoreProject">
						<i :class="loading"></i>
						<span>加载更多...</span>
					</div>
					<el-skeleton v-if="projectList.length <= 0" :rows="6" />
				</ul>
			</el-card>
		</el-col>
	</el-row>
</template>

<script>
export default {
	name: 'GithubProjectCard',
	props: {
		title: {
			type: String,
			require: true,
			default: 'Github热门项目排行榜'
		},
		language: {
			type: String,
			require: true,
			default: 'all'
		},
		sortType: {
			type: String,
			require: true,
			default: 'desc'
		},
		pageSize: {
			type: Number,
			require: true,
			default: 3
		}
	},
	data() {
		return {
			pageNumber:1,
			projectList: [],
			tagColors: {
				typeScript: '#3178c6',
				javaScript: '#f1e05a',
				html: '#e34c26',
				css: '#563d7c',
				java: 'orange',
				python: '#3572A5',
				golang: '#00ADD8',
				go: '#00ADD8',
				shell: '#89e051',
				'c++': '#f34b7d',
				other: '#ededed'
			},
			loading: '',
		}
	},
	created() {
		this.getProject();
	},
	methods: {
		handleNumber(number) {
			if (number >= 1000) {
				number = number / 1000;
				return Math.round(number) + "k";
			}
			return number;
		},
		getRequestUrl(language) {
			const urls = {
				all: `https://api.github.com/search/repositories?q=stars:%3E1&sort=stars&order=${this.sortType}&per_page=${this.pageSize}&page=${this.pageNumber}`,
				goLang: `https://api.github.com/search/repositories?q=stars:%3E1+language:go&sort=stars&order=${this.sortType}&per_page=${this.pageSize}&page=${this.pageNumber}`
			};
			let url = urls['all'];
			Object.keys(urls).forEach(key => {
				if (key.toLowerCase() === language.toLowerCase()) {
					url = urls[key];
				}
			});
			return url;
		},
		handleTagColor(language) {
			let color = this.tagColors['other'];
			if (!language) {
				return color;
			}
			Object.keys(this.tagColors).forEach(key => {
				if (key.toLowerCase() === language.toLowerCase()) {
					color = this.tagColors[key];
				}
			});
			return color;
		},
		async getProject() {
			const url = this.getRequestUrl(this.language);
			let response = await this.$axios.get(url);
			this.projectList = response.items;
		},
		async getMoreProject() {
			this.loading = 'el-icon-loading';
			this.pageNumber += 1;
			const url = this.getRequestUrl(this.language);
			let response = await this.$axios.get(url);
			this.projectList = [...new Set(this.projectList.concat(response.items))];
			this.loading = '';
		},
	}
}
</script>

<style lang="less" scoped>
.box-card {
	width: 100%;

	.title {
		background-image: linear-gradient(0deg,
				rgba(0, 0, 0, 0.3) 0,
				transparent);
		background-color: #0079d3;
		height: 80px;
		width: 100%;
		color: #fff;
		font-size: 20px;
		line-height: 80px;
		padding-left: 10px;
		box-sizing: border-box;
		text-align: center;
		border-radius: 4px 4px 0px 0px;
	}

	.github-hot-project-list {
		.github-hot-project-item {
			margin-top: 1rem;
			padding: 0px 5px 0px 5px;

			.user-info {
				display: flex;

				.avatar {
					margin-right: 10px;
				}

				.project-name {
					font-weight: 600;
					display: flex;
					align-items: center;

					a {
						white-space: nowrap;
						overflow: hidden;
						text-overflow: ellipsis;
					}
				}
			}

			.introduction {
				margin: 0.5rem 0;
				font-size: 12px;
			}

			.meta {
				font-size: 12px;
				color: #a7a3a3;

				i {
					margin-right: 2px;
				}

				.forks {
					margin: 0 8px;
				}
			}
		}

		.next-btn {
			background-image: linear-gradient(0deg,
					rgba(0, 0, 0, 0.3) 0,
					transparent);
			background-color: #0079d3;
			color: #fff;
			height: 40px;
			line-height: 40px;
			text-align: center;
			cursor: pointer;
			font-weight: 600;

			i {
				margin-right: 5px;
			}
		}
	}
}
</style>