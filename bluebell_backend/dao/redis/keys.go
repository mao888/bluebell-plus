package redis

/*
	Redis Key
*/
// redis key 注意使用命名空间的方式，方便查询和拆分
const (
	KeyPostInfoHashPrefix = "bluebell:post:"
	KeyPostTimeZSet       = "bluebell:post:time"	// zset;帖子及发帖时间定义
	KeyPostScoreZSet      = "bluebell:post:score"	// zset;帖子及投票分数定义
	//KeyPostVotedUpSetPrefix   = "bluebell:post:voted:down:"
	//KeyPostVotedDownSetPrefix = "bluebell:post:voted:up:"
	KeyPostVotedZSetPrefix = "bluebell:post:voted:"	// zset;记录用户及投票类型;参数是post_id

	KeyCommunityPostSetPrefix = "bluebell:community:"	// set保存每个分区下帖子的id
)
