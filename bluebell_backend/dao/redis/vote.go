package redis

import (
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	OneWeekInSeconds          = 7 * 24 * 3600        // 一周的秒数
	OneMonthInSeconds         = 4 * OneWeekInSeconds // 一个月的秒数
	VoteScore         float64 = 432                  // 每一票的值432分
	PostPerAge                = 20                   // 每页显示20条帖子
)

/*
投票算法：http://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_reddit.html
本项目使用简化版的投票分数
投一票就加432分 86400/200 -> 200张赞成票就可以给帖子在首页续天  -> 《redis实战》
*/

/* PostVote 为帖子投票
投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况
	1.之前没投过票，现在要投赞成票		 --> 更新分数和投票记录	差值的绝对值：1	+432
	2.之前投过反对票，现在要改为赞成票	 --> 更新分数和投票记录	差值的绝对值：2	+432*2
v=0时，有两种情况
	1.之前投过反对票，现在要取消			 --> 更新分数和投票记录	差值的绝对值：1	+432
	2.之前投过赞成票，现在要取消			 --> 更新分数和投票记录	差值的绝对值：1	-432
v=-1时，有两种情况
	1.之前没投过票，现在要投反对票		 --> 更新分数和投票记录	差值的绝对值：1	-432
	2.之前投过赞成票，现在要改为反对票	 --> 更新分数和投票记录	差值的绝对值：2	-432*2

投票的限制：
每个帖子子发表之日起一个星期之内允许用户投票，超过一个星期就不允许投票了
	1、到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2、到期之后删除那个 KeyPostVotedZSetPrefix
*/

// VoteForPost	为帖子投票
func VoteForPost(userID string, postID string, v float64) (err error) {
	// 1.判断投票限制
	// 去redis取帖子发布时间
	postTime := client.ZScore(KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds { // 超过一个星期就不允许投票了
		// 不允许投票了
		return ErrorVoteTimeExpire
	}
	// 2、更新帖子的分数
	// 2和3 需要放到一个pipeline事务中操作
	// 判断是否已经投过票 查当前用户给当前帖子的投票记录
	key := KeyPostVotedZSetPrefix + postID
	ov := client.ZScore(key, userID).Val()

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if v == ov {
		return ErrVoteRepeated
	}
	var op float64
	if v > ov {
		op = 1
	} else {
		op = -1
	}
	diffAbs := math.Abs(ov - v)                // 计算两次投票的差值
	pipeline := client.TxPipeline()            // 事务操作
	incrementScore := VoteScore * diffAbs * op // 计算分数（新增）
	// ZIncrBy 用于将有序集合中的成员分数增加指定数量
	_, err = pipeline.ZIncrBy(KeyPostScoreZSet, incrementScore, postID).Result() // 更新分数
	if err != nil {
		return err
	}
	// 3、记录用户为该帖子投票的数据
	if v == 0 {
		_, err = client.ZRem(key, postID).Result()
	} else {
		pipeline.ZAdd(key, redis.Z{ // 记录已投票
			Score:  v, // 赞成票还是反对票
			Member: userID,
		})
	}
	// 4、更新帖子的投票数
	pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", int64(op))

	//switch math.Abs(ov) - math.Abs(v) {
	//case 1:
	//	// 取消投票 ov=1/-1 v=0
	//	// 投票数-1
	//	pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", -1)
	//case 0:
	//	// 反转投票 ov=-1/1 v=1/-1
	//	// 投票数不用更新
	//case -1:
	//	// 新增投票 ov=0 v=1/-1
	//	// 投票数+1
	//	pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", 1)
	//default:
	//	// 已经投过票了
	//	return ErrorVoted
	//}
	_, err = pipeline.Exec()
	return err
}

// CreatePost redis存储帖子信息 使用hash存储帖子信息
func CreatePost(postID, userID uint64, title, summary string, CommunityID uint64) (err error) {
	now := float64(time.Now().Unix())
	votedKey := KeyPostVotedZSetPrefix + strconv.Itoa(int(postID))
	communityKey := KeyCommunityPostSetPrefix + strconv.Itoa(int(CommunityID))
	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}

	// 事务操作
	pipeline := client.TxPipeline()
	// 投票 zSet
	pipeline.ZAdd(votedKey, redis.Z{ // 作者默认投赞成票
		Score:  1,
		Member: userID,
	})
	pipeline.Expire(votedKey, time.Second*OneMonthInSeconds*6) // 过期时间：6个月
	// 文章 hash
	pipeline.HMSet(KeyPostInfoHashPrefix+strconv.Itoa(int(postID)), postInfo)
	// 添加到分数 ZSet
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  now + VoteScore,
		Member: postID,
	})
	// 添加到时间 ZSet
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  now,
		Member: postID,
	})
	// 添加到对应版块 把帖子添加到社区 set
	pipeline.SAdd(communityKey, postID)
	_, err = pipeline.Exec()
	return
}

// GetPost 从key中分页取出帖子
func GetPost(order string, page int64) []map[string]string {
	key := KeyPostScoreZSet
	if order == "time" {
		key = KeyPostTimeZSet
	}
	start := (page - 1) * PostPerAge
	end := start + PostPerAge - 1
	ids := client.ZRevRange(key, start, end).Val()
	postList := make([]map[string]string, 0, len(ids))
	for _, id := range ids {
		postData := client.HGetAll(KeyPostInfoHashPrefix + id).Val()
		postData["id"] = id
		postList = append(postList, postData)
	}
	return postList
}

// GetCommunityPost 分社区根据发帖时间或者分数取出分页的帖子
func GetCommunityPost(communityName, orderKey string, page int64) []map[string]string {
	key := orderKey + communityName // 创建缓存键

	if client.Exists(key).Val() < 1 {
		client.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, KeyCommunityPostSetPrefix+communityName, orderKey)
		client.Expire(key, 60*time.Second)
	}
	return GetPost(key, page)
}

// Reddit Hot rank algorithms
// from https://github.com/reddit-archive/reddit/blob/master/r2/r2/lib/db/_sorts.pyx
func Hot(ups, downs int, date time.Time) float64 {
	s := float64(ups - downs)
	order := math.Log10(math.Max(math.Abs(s), 1))
	var sign float64
	if s > 0 {
		sign = 1
	} else if s == 0 {
		sign = 0
	} else {
		sign = -1
	}
	seconds := float64(date.Second() - 1577808000)
	return math.Round(sign*order + seconds/43200)
}
