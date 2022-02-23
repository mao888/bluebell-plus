/**
    @author:Huchao
    @data:2022/2/17
    @note: 专门用来放接口文档用到的model
**/
package controller

import "bluebell_backend/models"

// 因为我们的接口文档返回的数据格式是一致的，但是具体的data类型不一致
type _ResponsePostList struct {
	Code  	MyCode  				 `json:"code"`	   // 业务响应状态码
	Message string					 `json:"message"`  // 提示信息
	Data    []*models.ApiPostDetail	 `json:"data"`	   // 数据
}