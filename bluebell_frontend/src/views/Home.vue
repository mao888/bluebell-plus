<template>
  <div class="content">
    <div class="left">
      <!-- <h4 class="c-l-title">热门帖子</h4> -->
      <div class="c-l-header">
        <div class="new btn-iconfont"
        :class="{ active: timeOrder }"
        @click="selectOrder('time')"
        >
          <i class="iconfont icon-polygonred"></i>New
        </div>
        <div class="top btn-iconfont"
         :class="{ active: scoreOrder }"
         @click="selectOrder('score')"
        >
          <i class="iconfont icon-top"></i>Score
        </div>
        <button class="btn-publish" @click="goPublish">发表</button>
      </div>
      <ul class="c-l-list">
        <li class="c-l-item"  v-for="post in postList" :key="post.id">
          <div class="post">
            <a class="vote">
              <span class="iconfont icon-up"
              @click="vote(post.id, '1')"
              ></span>
            </a>
            <span class="text">{{post.vote_num}}</span>
            <a class="vote">
              <span class="iconfont icon-down"
              @click="vote(post.id, '-1')"
              ></span>
            </a>
          </div>
          <div class="l-container" @click="goDetail(post.id)">
            <h4 class="con-title">{{post.title}}</h4>
            <div class="con-memo">
              <p>{{post.content}}</p>
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
      <div class="communities">
        <h2 class="r-c-title">今日火热频道排行榜</h2>
        <ul class="r-c-content">
          <li class="r-c-item">
            <span class="index">1</span>
            <i class="icon"></i>
            b/coding
          </li>
          <li class="r-c-item">
            <span class="index">2</span>
            <i class="icon"></i>
            b/tree_hole
          </li>
          <li class="r-c-item">
            <span class="index">3</span>
            <i class="icon"></i>
            b/job
          </li>
        </ul>
        <button class="view-all">查看所有</button>
      </div>
      <div class="r-trending">
        <h2 class="r-t-title">持续热门频道</h2>
        <ul class="rank">
          <li class="r-t-cell">
            <div class="r-t-cell-info">
              <div class="avatar"></div>
              <div class="info">
                <span class="info-title">b/Book</span>
                <p class="info-num">7.1k members</p>
              </div>
            </div>
            <button class="join-btn">JOIN</button>
          </li>
          <li class="r-t-cell">
            <div class="r-t-cell-info">
              <div class="avatar"></div>
              <div class="info">
                <span class="info-title">b/coding</span>
                <p class="info-num">3.2k members</p>
              </div>
            </div>
            <button class="join-btn">JOIN</button>
          </li>
          <li class="r-t-cell">
            <div class="r-t-cell-info">
              <div class="avatar"></div>
              <div class="info">
                <span class="info-title">b/job</span>
                <p class="info-num">2.5k members</p>
              </div>
            </div>
            <button class="join-btn">JOIN</button>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "Home",
  components: {},
  data() {
    return {
      order: "time",
      page: 1,
      postList: []
    };
  },
  methods: {
    selectOrder(order){
      this.order = order;
      this.getPostList()
    },
    goPublish(){
      this.$router.push({ name: "Publish" });
    },
    goDetail(id){
      this.$router.push({ name: "Content", params: { id: id }});
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
    vote(post_id, direction){
      this.$axios({
        method: "post",
        url: "/vote",
        data: JSON.stringify({
          post_id: post_id,
          direction: direction,
        })
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
    }
  },
  mounted: function() {
    this.getPostList();
  },
  computed:{
    timeOrder(){
      return this.order == "time";
    },
    scoreOrder(){
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
  .left {
    width: 640px;
    padding-bottom: 10px;
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
        width: 64px;
        height: 32px;
        line-height: 32px;
        background-color: #54b351;
        color: #ffffff;
        border: 1px solid transparent;
        border-radius: 4px;
        box-sizing: border-box;
        text-align: center;
        margin-left: auto;
        cursor: pointer;
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
            background: url("https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1585999647247&di=7e9061211c23e3ed9f0c4375bb3822dc&imgtype=0&src=http%3A%2F%2Fi1.hdslb.com%2Fbfs%2Farchive%2F04d8cda08e170f4a58c18c45a93c539375c22162.jpg")
              no-repeat;
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
    margin-left: 24px;
    margin-top: 28px;
    .communities {
      background-color: #ffffff;
      color: #1a1a1b;
      border: 1px solid #ccc;
      border-radius: 4px;
      overflow: visible;
      word-wrap: break-word;
      margin-bottom: 20px;
      .r-c-title {
        background-image: linear-gradient(
          0deg,
          rgba(0, 0, 0, 0.3) 0,
          transparent
        );
        background-color: #0079d3;
        height: 80px;
        width: 100%;
        color: #fff;
        font-size: 20px;
        line-height: 120px;
        padding-left: 10px;
        box-sizing: border-box;
        text-align: center;
      }
      .r-c-content {
        .r-c-item {
          align-items: center;
          display: flex;
          display: -webkit-flex;
          height: 48px;
          padding: 0 12px;
          border-bottom: thin solid #edeff1;
          font-size: 14px;
          .index {
            width: 20px;
            color: #1c1c1c;
            font-size: 14px;
            font-weight: 500;
            line-height: 18px;
          }
          .icon {
            width: 32px;
            height: 32px;
            background-image: url("../assets/images/avatar.png");
            background-repeat: no-repeat;
            background-size: cover;
            margin-right: 20px;
          }
          &:last-child {
            border-bottom: none;
          }
        }
      }
      .view-all {
        background-color: #0079d3;
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
        padding: 3px 0;
        width: 280px;
        color: #fff;
        margin: 20px 0 20px 16px;
      }
    }
    .r-trending {
      padding-top: 16px;
      width: 312px;
      background-color: #ffffff;
      color: #1a1a1b;
      fill: #1a1a1b;
      border: 1px solid #cccccc;
      border-radius: 4px;
      overflow: visible;
      word-wrap: break-word;
      .r-t-title {
        font-size: 10px;
        font-weight: 700;
        letter-spacing: 0.5px;
        line-height: 12px;
        text-transform: uppercase;
        background-color: #ffffff;
        border-radius: 3px 3px 0 0;
        color: #1a1a1b;
        display: -ms-flexbox;
        display: flex;
        fill: #1a1a1b;
        padding: 0 12px 12px;
      }
      .rank {
        padding: 12px;
        .r-t-cell {
          display: flex;
          display: -webkit-flex;
          align-items: center;
          justify-content: space-between;
          margin-bottom: 16px;
          .r-t-cell-info {
            display: flex;
          }
          .avatar {
            width: 32px;
            height: 32px;
            background: url("../assets/images/avatar.png") no-repeat;
            background-size: cover;
            margin-right: 10px;
          }
          .info {
            margin-right: 10px;
            .info-title {
              font-size: 12px;
              font-weight: 500;
              line-height: 16px;
              text-overflow: ellipsis;
              width: 144px;
            }
            .info-num {
              font-size: 12px;
              font-weight: 400;
              line-height: 16px;
              padding-bottom: 4px;
            }
          }
          .join-btn {
            width: 106px;
            height: 32px;
            line-height: 32px;
            background-color: #0079d3;
            color: #ffffff;
            border: 1px solid transparent;
            border-radius: 4px;
            box-sizing: border-box;
            text-align: center;
          }
        }
      }
    }
  }
}
</style>
