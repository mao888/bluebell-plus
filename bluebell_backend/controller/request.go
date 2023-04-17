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

// getCurrentUserID 获取当前登录用户ID
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

// getPageInfo 分页参数
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	SizeStr := c.Query("size")

	var (
		page int64 // 第几页 页数
		size int64 // 每页几条数据
		err  error
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
