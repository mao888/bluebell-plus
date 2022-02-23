package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	ContextUserIDKey = "userID"
)

var (
	ErrorUserNotLogin = errors.New("当前用户未登录")
)

/**
 * @Author huchao
 * @Description //TODO 获取当前登录用户ID
 * @Date 10:41 2022/2/11
 **/
// getCurrentUserID 获取当前登录用户ID
// @Summary 获取当前登录用户ID
// @Description 获取当前登录用户ID
// @Tags 公共接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query userID  path    int     true        "_userID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

/**
 * @Author huchao
 * @Description //TODO 分页参数
 * @Date 23:41 2022/2/12
 **/
// getPageInfo 分页参数
// @Summary 分页参数
// @Description 分页参数
// @Tags 公共接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
func getPageInfo(c *gin.Context) (int64, int64)  {
	pageStr := c.Query("page")
	SizeStr := c.Query("size")

	var (
		page int64		// 第几页 页数
		size int64	    // 每页几条数据
		err error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(SizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}