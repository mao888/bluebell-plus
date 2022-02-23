package controller

import (
	"bluebell_backend/logic"
	"bluebell_backend/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/**
 * @Author huchao
 * @Description //TODO 创建帖子
 * @Date 17:40 2022/2/12
 **/
// CreatePostHandler 创建帖子
// @Summary 创建帖子
// @Description 创建帖子
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.Post false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post [POST]
func CreatePostHandler(c *gin.Context) {
	// 1、获取参数及校验参数
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {   // validator --> binding tag
		zap.L().Debug("c.ShouldBindJSON(post) err",zap.Any("err",err))
		zap.L().Error("create post with invalid parm")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 参数校验

	// 获取作者ID，当前请求的UserID(从c取到当前发请求的用户ID)
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorId = userID
	// 2、创建帖子
	err = logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, nil)
}

// PostListHandler 帖子列表
//func PostListHandler(c *gin.Context) {
//	order, _ := c.GetQuery("order")
//	pageStr, ok := c.GetQuery("page")
//	if !ok {
//		pageStr = "1"
//	}
//	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
//	if err != nil {
//		pageNum = 1
//	}
//	posts := redis.GetPost(order, pageNum)
//	fmt.Println(len(posts))
//	ResponseSuccess(c, posts)
//}

/**
 * @Author huchao
 * @Description //TODO 分页获取帖子列表
 * @Date 22:55 2022/2/12
 **/
// PostListHandler 分页获取帖子列表
// @Summary 分页获取帖子列表
// @Description 分页获取帖子列表
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts [GET]
func PostListHandler(c *gin.Context) {
	// 获取分页参数
	page,size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page,size)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

/**
 * @Author huchao
 * @Description //TODO 升级版帖子列表接口：按创建时间排序 或者 按照 分数排序
 * @Date 21:34 2022/2/15
 **/
// 根据前端传来的参数动态的获取帖子列表
// 按创建时间排序 或者 按照 分数排序
// 1、获取请求的query string 参数
// 2、去redis查询id列表
// 3、根据id去数据库查询帖子详细信息
// PostList2Handler 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func PostList2Handler(c *gin.Context)  {
	// GET请求参数(query string)： /api/v1/posts2?page=1&size=10&order=time
	// 获取分页参数
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
		Order: models.OrderTime,	// magic string
	}
	//c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("PostList2Handler with invalid params",zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 获取数据
	data, err := logic.GetPostListNew(p)	// 更新：合二为一
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

/**
 * @Author huchao
 * @Description //TODO 根据Id查询帖子详情
 * @Date 17:44 2022/2/12
 **/
// PostDetailHandler 根据Id查询帖子详情
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query postId  path    int     true        "id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post/:id [get]
func PostDetailHandler(c *gin.Context) {
	// 1、获取参数(从URL中获取帖子的id)
	postIdStr := c.Param("id")
	postId,err := strconv.ParseInt(postIdStr,10,64)
	if err != nil {
		zap.L().Error("get post detail with invalid param",zap.Error(err))
		ResponseError(c,CodeInvalidParams)
	}

	// 2、根据id取出id帖子数据(查数据库)
	post, err := logic.GetPostById(postId)
	if err != nil {
		zap.L().Error("logic.GetPost(postID) failed", zap.Error(err))
		ResponseError(c,CodeServerBusy)
	}

	// 3、返回响应
	ResponseSuccess(c, post)
}

/**
 * @Author huchao
 * @Description //TODO 根据社区去查询帖子列表
 * @Date 22:44 2022/2/16
 **/
func GetCommunityPostListHandler(c *gin.Context)  {
	// GET请求参数(query string)： /api/v1/posts2?page=1&size=10&order=time
	// 获取分页参数
	p := &models.ParamPostList{
		CommunityID: 0,
		Page:        1,
		Size:        10,
		Order:       models.OrderScore,
	}
	//c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid params",zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取数据
	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}