package controller

import (
	"bluebell_backend/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 社区

/**
 * @Author huchao
 * @Description //TODO 查找社区列表
 * @Date 16:28 2022/2/12
 **/
// CommunityHandler 社区列表
// @Summary 社区列表
// @Description 社区列表
// @Tags 社区业务接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.Community false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区(community_id,community_name)以列表的形式返回
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)	// 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, communityList)
}

/**
 * @Author huchao
 * @Description //TODO 根据ID查找到社区分类的详情
 * @Date 17:01 2022/2/12
 **/
// CommunityDetailHandler 社区详情
// @Summary 社区详情
// @Description 社区详情
// @Tags 社区业务接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query communityId     path    int     true        "id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {
	// 1、获取社区ID
	communityIdStr := c.Param("id")	// 获取URL参数
	communityId, err := strconv.ParseUint(communityIdStr,10,64) // id字符串格式转换
	if err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}

	// 2、根据ID获取社区详情
	communityList, err := logic.GetCommunityDetailByID(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityByID() failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeSuccess, err.Error())
		return
	}
	ResponseSuccess(c, communityList)
}
