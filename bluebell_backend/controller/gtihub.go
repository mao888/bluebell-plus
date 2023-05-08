package controller

import (
	"bluebell_backend/logic"
	"bluebell_backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GithubTrendingHandler 获取Github热榜项目
func GithubTrendingHandler(c *gin.Context) {
	p := &models.ParamGithubTrending{}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GithubTrendingHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取数据
	data, err := logic.GetGithubTrending(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
