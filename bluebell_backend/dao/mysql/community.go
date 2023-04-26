package mysql

import (
	"bluebell_backend/models"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

// GetCommunityList 查询分类社区列表
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows { // 查询为空
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return communityList, nil
}

func GetCommunityNameByID(idStr string) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := `select community_id, community_name
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(ErrorInvalidID)

		}
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		return nil, errors.New(ErrorQueryFailed)
	}
	return
}

// GetCommunityByID 根据ID查询分类社区详情
func GetCommunityByID(id uint64) (*models.CommunityDetailRes, error) {
	community := new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err := db.Get(community, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows { // 查询为空
			return nil, errors.New(ErrorInvalidID) // 无效的ID return
		}
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		return nil, errors.New(ErrorQueryFailed)
	}
	return &models.CommunityDetailRes{
		CommunityID:   community.CommunityID,
		CommunityName: community.CommunityName,
		Introduction:  community.Introduction,
		CreateTime:    community.CreateTime.Format("2006-01-02 15:04:05"),
	}, err
}
