package mysql

import (
	"bluebell_backend/models"
	"database/sql"

	"go.uber.org/zap"
)

/**
 * @Author huchao
 * @Description //TODO 查询分类社区列表
 * @Date 16:42 2022/2/12
 **/
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {	// 查询为空
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return
}

func GetCommunityNameByID(idStr string) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := `select community_id, community_name
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

/**
 * @Author huchao
 * @Description //TODO 根据ID查询分类社区详情
 * @Date 17:08 2022/2/12
 **/
func GetCommunityByID(id uint64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows {	// 查询为空
		err = ErrorInvalidID	// 无效的ID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
	}
	return community,err
}
