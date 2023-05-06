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
