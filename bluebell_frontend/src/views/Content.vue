<template>
  <div class="content">
    <div class="left">
      <div class="container">
        <div class="post">
          <a class="vote">
            <span class="iconfont icon-up" @click="vote(post.post_id, 1)"></span>
          </a>
          <span class="text">{{ post.vote_num }}</span>
          <a class="vote">
            <span class="iconfont icon-down" @click="vote(post.post_id, -1)"></span>
          </a>
        </div>
        <div class="l-container">
          <h4 class="con-title">{{ post.title }}</h4>
          <div class="con-info markdown-body" v-html="post.content"></div>
        </div>
      </div>
      <!-- 评论区 -->
      <Comment :sourceId="this.$route.params.id"></Comment>
    </div>
    <div class="right">
      <div class="topic-info">
        <h5 class="t-header"></h5>
        <div class="t-info">
          <a class="avatar"></a>
          <span class="topic-name">b/{{ post.community.community_name }}</span>
        </div>
        <p class="t-desc">{{ post.community.introduction }}</p>
        <p class="t-create-time">{{ post.community.create_time }}</p>
        <div class="date">{{ create_time }}</div>
        <button class="topic-btn" @click="goCommunityDetail(post.community.community_id)">JOIN</button>
      </div>
    </div>
  </div>
</template>
<script>
import Comment from '../components/Comment.vue';
import Vue from 'vue';
export default {
  name: "Content",
  components: { Comment },
  data() {
    return {
      post: {},
    }
  },
  methods: {
    getPostDetail() {
      this.$axios({
        method: "get",
        url: "/post/" + this.$route.params.id,
      })
        .then(response => {
          console.log(1, response.data);
          if (response.code == 1000) {
            let MarkdownIt = require('markdown-it');
            let md = new MarkdownIt();
            this.post = response.data;
            this.post.content = md.render(this.post.content);
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
            this.getPostDetail();
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
    goCommunityDetail(community_id) {
      this.$router.push({
        name: 'Community',
        params: {
          id: community_id
        }
      });
    },
  },
  mounted() {
    this.getPostDetail();
  }
};
</script>

<style lang="less" scoped>
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

  .left {
    flex-grow: 1;
    max-width: 740px;
    border-radius: 4px;
    word-break: break-word;
    // background: #ffffff;
    border: #edeff1;
    flex: 1;
    margin: 32px;
    margin-right: 12px;
    padding-bottom: 30px;
    position: relative;

    .container {
      width: 100%;
      height: auto;
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
        margin-left: 40px;
        background-color: rgba(255, 255, 255, 0.8);

        .con-title {
          color: #000000;
          font-size: 18px;
          font-weight: 500;
          line-height: 22px;
          text-decoration: none;
          word-break: break-word;
        }

        .con-info {
          margin: 25px 0;
          line-height: 2;
          code{
            overflow-x: scroll !important;
            ::-webkit-scrollbar {
              width: 10px;
              height: 10px;
            }
          }
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
          font-size: 12px;
          display: flex;
          display: -webkit-flex;

          .btn-item {
            display: flex;
            display: -webkit-flex;
            align-items: center;
            margin-right: 10px;

            .iconfont {
              margin-right: 4px;
            }
          }
        }
      }
    }

    .comment {
      width: 100%;
      height: auto;
      position: relative;

      .c-left {
        .line {
          border-right: 2px solid #edeff1;
          // width: 20px;
          height: 100%;
          position: absolute;
          left: 20px;
        }

        .c-arrow {
          display: flex;
          display: -webkit-flex;
          position: absolute;
          z-index: 2;
          flex-direction: column;
          left: 12px;
          background: #ffffff;
          padding-bottom: 5px;
        }
      }

      .c-right {
        margin-left: 40px;
        padding-right: 10px;

        .c-user-info {
          margin-bottom: 10px;

          .name {
            color: #1c1c1c;
            font-size: 12px;
            font-weight: 400;
            line-height: 16px;
          }

          .num {
            padding-left: 4px;
            font-size: 12px;
            font-weight: 400;
            line-height: 16px;
            color: #7c7c7c;
          }
        }

        .c-content {
          font-family: Noto Sans, Arial, sans-serif;
          font-size: 14px;
          font-weight: 400;
          line-height: 21px;
          word-break: break-word;
          color: rgb(26, 26, 27);
        }
      }
    }
  }

  .right {
    flex-grow: 0;
    width: 312px;
    margin-top: 32px;

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

      .date {
        font-family: Noto Sans, Arial, sans-serif;
        font-size: 14px;
        font-weight: 400;
        line-height: 18px;
        margin-top: 20px;
        padding: 0 12px;
      }

      .topic-btn {
        width: 286px;
        height: 34px;
        line-height: 34px;
        color: #ffffff;
        margin: 12px auto 0 auto;
        background: #003f6d;
        border-radius: 4px;
        box-sizing: border-box;
        margin-left: 13px;
      }
    }
  }
}
</style>