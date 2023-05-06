<template>
  <div>
    <div class="content">
      <div class="left">
        <SideBar></SideBar>
      </div>
      <div class="center">
        <!-- <h4 class="c-l-title">热门帖子</h4> -->
        <div class="c-l-header">
          <div class="new btn-iconfont" :class="{ active: timeOrder }" @click="selectOrder('time')">
            <i class="iconfont icon-polygonred"></i>New
          </div>
          <div class="top btn-iconfont" :class="{ active: scoreOrder }" @click="selectOrder('score')">
            <i class="iconfont icon-top"></i>Score
          </div>
          <div class="search">
            <i class="search-icon el-icon-search" @click="searchPost"></i>
            <input type="text" class="s-input" v-model="keyword" @keydown.enter="searchPost" placeholder="文章搜索" />
          </div>
          <div class="publish-btn" @click="goPublish">发表</div>
        </div>
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
              <!-- <div class="user-btn">
              <span class="btn-item">
                <i class="iconfont icon-comment"></i>
                <span>{{post.comments}} comments</span>
              </span>
            </div> -->
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
        <div class="run-time-container">
          <TimeMeter></TimeMeter>
        </div>
        <div class="github-project-card-container">
          <GithubProjectCard language="all"></GithubProjectCard>
        </div>
        <div class="github-golang-project-card-container">
          <GithubProjectCard language="golang" title="Golang热门项目排行榜"></GithubProjectCard>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SideBar from '../components/SideBar.vue';
// @ is an alias to /src
import TimeMeter from '../components/TimeMeter.vue';
import GithubProjectCard from './components/GithubProjectCard.vue';
import Vue from 'vue';
export default {
  name: "Home",
  components: { TimeMeter, SideBar, GithubProjectCard },
  data() {
    return {
      order: "time",
      postList: [],
      pageNumber: 1,
      pageSize: 5,
      pageTotal: {},
      keyword: '',
      isSearch: false
    };
  },
  methods: {
    selectOrder(order) {
      this.order = order;
      this.getPostList();
    },
    handleCurrentChange(val) {
      this.pageNumber = val;
      if (!this.isSearch) {
        this.getPostList();
      } else {
        this.searchPost();
      }
    },
    handleSizeChange(val) {
      this.pageSize = val;
      if (!this.isSearch) {
        this.getPostList();
      } else {
        this.searchPost();
      }
    },
    goDetail(id) {
      this.$router.push({ name: "Content", params: { id: id } });
    },
    getPostList() {
      this.$axios({
        method: "get",
        url: "/posts2",
        params: {
          page: this.pageNumber,
          size: this.pageSize,
          order: this.order,
        }
      })
        .then(response => {
          console.log(response.data, 222);
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
    goPublish() {
      this.$router.push({ name: "Publish" });
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
            this.getPostList();
          } else if (response.code == 1009) {
            Vue.prototype.$message.error('请勿重复投票')
          } else if (response.code == 1010) {
            Vue.prototype.$message.error('已过投票时间')
          } else {
            console.log(response.msg);
            Vue.prototype.$message.error('请先登录')
          }
        })
        .catch(error => {
          console.log(error);
        });
    },
    async searchPost() {
      if (!this.keyword) {
        this.isSearch = false;
        this.getPostList();
        return;
      }
      this.isSearch = true;
      const response = await this.$axios({
        method: "get",
        url: "/search",
        params: {
          page: this.pageNumber,
          size: this.pageSize,
          search: this.keyword
        }
      });
      if (response.code === 1000) {
        console.log(response.data);
        this.postList = response.data.list;
        this.pageTotal = response.data.page;
      } else {
        console.log(response.message);
      }
    },
  },
  mounted: function () {
    this.getPostList();
  },
  computed: {
    timeOrder() {
      return this.order == "time";
    },
    scoreOrder() {
      return this.order == "score";
    }
  }
};
</script>

<style scoped lang="less">
.content {
  max-width: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin: 48px auto 0;
  padding: 20px 24px;
  background: #6190E8;
  /* fallback for old browsers */
  background: -webkit-linear-gradient(to right, #6190E8, #A7BFE8);
  /* Chrome 10-25, Safari 5.1-6 */
  background: linear-gradient(to right, #6190E8, #A7BFE8);
  /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */


  .left {
    width: 312px;
    height: fit-content;
    margin-top: 28px;
    background: #fff;
    border-radius: 4px;
  }

  .center {
    width: 640px;
    padding-bottom: 10px;
    margin: 0 24px;

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

      .search {
        flex-grow: 0.5;
        margin: 0 auto;
        max-width: 650px;
        position: relative;
        display: flex;
        display: -webkit-flex;
        align-items: center;
        border-radius: 4px;

        .search-icon {
          width: 18px;
          height: 18px;
          line-height: 18px;
          background-size: cover;
          display: inline-block;
          position: absolute;
          cursor: pointer;
          border-radius: 4px;
          right: 1rem;
          padding: 5px;

          &:hover {
            background: silver;
          }
        }

        .s-input {
          flex-grow: 1;
          -webkit-appearance: none;
          appearance: none;
          background-color: #f6f7f8;
          border-radius: 4px;
          border: 1px solid #edeff1;
          box-shadow: none;
          color: #c1c1c1;
          display: block;
          height: 36px;
          outline: none;
          width: 100%;
          text-indent: 1rem;
        }
      }

      .publish-btn {
        width: fit-content;
        height: fit-content;
        border: 1px solid transparent;
        border-radius: 4px;
        box-sizing: border-box;
        text-align: center;
        letter-spacing: 1px;
        text-decoration: none;
        font-size: 12px;
        font-weight: 700;
        letter-spacing: 0.5px;
        line-height: 24px;
        text-transform: uppercase;
        padding: 3px 16px;
        fill: #0079d3;
        display: inline-block;
        cursor: pointer;
        border-color: #0079d3;
        background: #0079d3;
        color:#fff;

        &:nth-child(2) {
          margin: 0 10px;
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

    .pagination-block {
      background: #fff;
      padding: 8px;
    }
  }

  .right {
    width: 312px;
    margin-top: 28px;

    .run-time-container {
      margin-bottom: 1rem;
    }

    .github-project-card-container {
      margin: 1rem 0;
    }
  }
}
</style>
