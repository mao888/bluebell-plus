package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamPostList 获取帖子列表query 参数
type ParamPostList struct {
	Search      string `json:"search" form:"search"`               // 关键字搜索
	CommunityID uint64 `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page"`                   // 页码
	Size        int64  `json:"size" form:"size"`                   // 每页数量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

// ParamGithubTrending 获取Github热榜项目query 参数
type ParamGithubTrending struct {
	Language int   `json:"language" form:"language"` // 语言 0：All 1：Go 2：Python 3：JavaScript 4：Java
	Page     int64 `json:"page" form:"page"`         // 页码
	Size     int64 `json:"size" form:"size"`         // 每页数量
}
