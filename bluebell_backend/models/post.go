package models

import (
	"encoding/json"
	"errors"
	"time"
)

// Post 帖子Post结构体 内存对齐概念 字段类型相同的对齐 缩小变量所占内存大小
type Post struct {
	PostID      uint64    `json:"post_id,string" db:"post_id"`
	AuthorId    uint64    `json:"author_id" db:"author_id"`
	CommunityID uint64    `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"-" db:"create_time"`
	UpdateTime  time.Time `json:"-" db:"update_time"`
}

// UnmarshalJSON 为Post类型实现自定义的UnmarshalJSON方法
func (p *Post) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Title       string `json:"title" db:"title"`
		Content     string `json:"content" db:"content"`
		CommunityID int64  `json:"community_id" db:"community_id"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if required.CommunityID == 0 {
		err = errors.New("未指定版块")
	} else {
		p.Title = required.Title
		p.Content = required.Content
		p.CommunityID = uint64(required.CommunityID)
	}
	return
}

// ApiPostDetail 帖子返回的详情结构体
type ApiPostDetail struct {
	*Post                                  // 嵌入帖子结构体
	*CommunityDetailRes `json:"community"` // 嵌入社区信息
	AuthorName          string             `json:"author_name"`
	VoteNum             int64              `json:"vote_num"` // 投票数量
	//CommunityName string `json:"community_name"`
}

type Page struct {
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
}

type ApiPostDetailRes struct {
	Page Page             `json:"page"`
	List []*ApiPostDetail `json:"list"`
}
