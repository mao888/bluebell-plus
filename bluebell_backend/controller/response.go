package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Author huchao
 * @Description //TODO 封装响应
 * @Date 22:11 2022/2/10
 * @Param
 * @return
 **/

type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`	// omitempty当data为空时,不展示这个字段
}

func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
