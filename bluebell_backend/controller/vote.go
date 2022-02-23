package controller

import (
	"bluebell_backend/logic"
	"bluebell_backend/models"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type VoteData struct {
	//UserID int 从请求中获取当前的用户
	PostID    string  `json:"post_id,string"`	 // 帖子id
	Direction int     `json:"direction,string"`  // 赞成票(1)还是反对票(-1)取消投票(0)
}

// UnmarshalJSON 为VoteData类型实现自定义的UnmarshalJSON方法
func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string  `json:"post_id"`
		Direction int `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}

/**
 * @Author huchao
 * @Description //TODO 投票
 * @Date 10:53 2022/2/14
 **/
// VoteHandler 投票
// @Summary 投票
// @Description 投票
// @Tags 投票业务接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /vote [POST]
func VoteHandler(c *gin.Context) {
	// 参数校验,给哪个文章投什么票
	vote := new(models.VoteDataForm)
	if err := c.ShouldBindJSON(&vote); err != nil {
		errs, ok := err.(validator.ValidationErrors)	// 类型断言
		if !ok {
			ResponseError(c,CodeInvalidParams)
			return
		}
		errdata := removeTopStruct(errs.Translate(trans))	// 翻译并去除掉错误提示中的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParams, errdata)
		return
	}
	// 获取当前请求用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, vote); err != nil {
		zap.L().Error("logic.VoteForPost() failed",zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
