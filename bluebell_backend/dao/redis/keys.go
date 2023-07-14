package redis

// redis key 注意使用命名空间的方式，方便查询和拆分
const (
	KeyPostInfoHashPrefix = "bluebell-plus:post:"
	KeyPostTimeZSet       = "bluebell-plus:post:time"  // zset;帖子及发帖时间定义
	KeyPostScoreZSet      = "bluebell-plus:post:score" // zset;帖子及投票分数定义
	//KeyPostVotedUpSetPrefix   = "bluebell-plus:post:voted:down:"
	//KeyPostVotedDownSetPrefix = "bluebell-plus:post:voted:up:"
	KeyPostVotedZSetPrefix    = "bluebell-plus:post:voted:" // zSet;记录用户及投票类型;参数是post_id
	KeyCommunityPostSetPrefix = "bluebell-plus:community:"  // set保存每个分区下帖子的id
)
