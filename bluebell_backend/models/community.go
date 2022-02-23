package models

import "time"

/**
 * @Author huchao
 * @Description //TODO Community结构体
 * @Date 16:39 2022/2/12
 **/
type Community struct {
	CommunityID   uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
}

/**
 * @Author huchao
 * @Description //TODO 社区详情model
 * @Date 17:14 2022/2/12
 **/
type CommunityDetail struct {
	CommunityID   uint64    `json:"community_id" db:"community_id"`
	CommunityName string    `json:"community_name" db:"community_name"`
	Introduction  string    `json:"introduction,omitempty" db:"introduction"`	// omitempty 当Introduction为空时不展示
	CreateTime    time.Time `json:"create_time" db:"create_time"`
}
