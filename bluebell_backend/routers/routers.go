package routers

import (
	"bluebell_backend/controller"
	_ "bluebell_backend/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell_backend/logger"
	"bluebell_backend/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-contrib/pprof"
)

/**
 * @Author huchao
 * @Description //TODO 设置路由
 * @Date 21:58 2022/2/10
 **/
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)		// 设置成发布模式
	}
	r := gin.New()												// 每两秒钟添加一个令牌  全局限流
	//r.Use(logger.GinLogger(), logger.GinRecovery(true),middlewares.RateLimitMiddleware(2*time.Second , 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//r := gin.Default()

	r.LoadHTMLFiles("templates/index.html")	// 加载html
	r.Static("/static", "./static")	// 加载静态文件
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)				// 注册业务路由
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	v1.GET("/posts", controller.PostListHandler)		// 分页展示帖子列表
	v1.GET("/posts2", controller.PostList2Handler) // 根据时间或者分数排序分页展示帖子列表
	v1.GET("/community", controller.CommunityHandler)	// 获取分类社区列表
	v1.GET("/community/:id", controller.CommunityDetailHandler)	// 根据ID查找社区详情
	v1.GET("/post/:id", controller.PostDetailHandler) // 查询帖子详情

	v1.Use(middlewares.JWTAuthMiddleware())	// 应用JWT认证中间件
	{
		//v1.GET("/community", controller.CommunityHandler)	// 获取分类社区列表
		//v1.GET("/community/:id", controller.CommunityDetailHandler)	// 根据ID查找社区详情

		v1.POST("/post", controller.CreatePostHandler)	 // 创建帖子
		//v1.GET("/post/:id", controller.PostDetailHandler) // 查询帖子详情
		//v1.GET("/posts", controller.PostListHandler)		// 分页展示帖子列表
		//
		//v1.GET("/posts2", controller.PostList2Handler) // 根据时间或者分数排序分页展示帖子列表

		v1.POST("/vote", controller.VoteHandler)		   // 投票

		v1.POST("/comment", controller.CommentHandler)
		v1.GET("/comment", controller.CommentListHandler)

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	pprof.Register(r)	// 注册pprof相关路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
