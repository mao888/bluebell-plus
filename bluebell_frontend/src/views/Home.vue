<template>
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
        <div class="btn-publish">
          <div class="word-of-day" @click="getWordOfDay" :title="wordOfDay">{{ wordOfDay }}</div>
          <div class="publish" @click="goPublish">发表</div>
        </div>
      </div>
      <ul class="c-l-list">
        <li class="c-l-item" v-for="post in postList" :key="post.post_id">
          <div class="post">
            <a class="vote">
              <span class="iconfont icon-up" @click="vote(post.post_id, 1)"></span>
            </a>
            <span class="text">{{ post.vote_num }}</span>
            <a class="vote">
              <span class="iconfont icon-down" @click="vote(post.id, -1)"></span>
            </a>
          </div>
          <div class="l-container" @click="goDetail(post.id)">
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
</template>

<script>
import SideBar from '../components/SideBar.vue';
// @ is an alias to /src
import TimeMeter from '../components/TimeMeter.vue';
import GithubProjectCard from './components/GithubProjectCard.vue';
export default {
  name: "Home",
  components: { TimeMeter, SideBar, GithubProjectCard },
  data() {
    return {
      order: "time",
      page: 1,
      postList: [],
      wordOfDay: ''
    };
  },
  created() {
    this.getWordOfDay();
  },
  methods: {
    selectOrder(order) {
      this.order = order;
      this.getPostList()
    },
    goPublish() {
      this.$router.push({ name: "Publish" });
    },
    goDetail(id) {
      this.$router.push({ name: "Content", params: { id: id } });
    },
    getPostList() {
      this.$axios({
        method: "get",
        url: "/posts2",
        params: {
          page: this.page,
          order: this.order,
        }
      })
        .then(response => {
          console.log(response.data, 222);
          if (response.code == 1000) {
            this.postList = response.data;
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
          } else {
            console.log(response.msg);
          }
        })
        .catch(error => {
          console.log(error);
        });
    },
    async getWordOfDay() {
      let response = await this.$axios.get(`https://v.api.aa1.cn/api/yiyan/index.php`);
      const reg = '<p>(.*)</p>';
      this.wordOfDay = response.match(reg)[1];
    }
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

      .btn-publish {
        height: 32px;
        width:78%;
        display: flex;
        position:relative;
        border-radius: 4px;

        .word-of-day {
          width:87%;
          line-height: 32px;
          font-size: 14px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          background-image: linear-gradient(to right, orange, purple);
          -webkit-background-clip: text;
          color: transparent;
          text-align:center;
          cursor: pointer;
          margin-left:1rem;
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
          position:absolute;
          right:0;
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
