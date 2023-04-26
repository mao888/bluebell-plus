package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
)

// GetCommunityList 查询分类社区列表
func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

// GetCommunityDetailByID 根据ID查询分类社区详情
func GetCommunityDetailByID(id uint64) (*models.CommunityDetailRes, error) {
	return mysql.GetCommunityByID(id)
}
