package routers

import (
	"bluebell_backend/controller"
	_ "bluebell_backend/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell_backend/logger"
	"bluebell_backend/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-contrib/pprof"
)

// SetupRouter 设置路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}
	//初始化 gin Engine  新建一个没有任何默认中间件的路由
	r := gin.New()
	//设置中间件
	r.Use(logger.GinLogger(),
		logger.GinRecovery(true),                           // Recovery 中间件会 recover掉项目可能出现的panic，并使用zap记录相关日志
		middlewares.RateLimitMiddleware(2*time.Second, 40), // 每两秒钟添加十个令牌  全局限流
	)

	r.LoadHTMLFiles("templates/index.html") // 加载html
	r.Static("/static", "./static")         // 加载静态文件
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	// 注册登陆业务
	v1.POST("/login", controller.LoginHandler)               // 登陆业务
	v1.POST("/signup", controller.SignUpHandler)             // 注册业务
	v1.GET("/refresh_token", controller.RefreshTokenHandler) // 刷新accessToken
	// 帖子业务
	v1.GET("/posts", controller.PostListHandler)      // 分页展示帖子列表
	v1.GET("/posts2", controller.PostList2Handler)    // 根据社区id及时间或者分数排序分页展示帖子列表
	v1.GET("/post/:id", controller.PostDetailHandler) // 查询帖子详情
	v1.GET("/search", controller.PostSearchHandler)   // 搜索业务-搜索帖子
	// 社区业务
	v1.GET("/community", controller.CommunityHandler)           // 获取分类社区列表
	v1.GET("/community/:id", controller.CommunityDetailHandler) // 根据ID查找社区详情
	// Github热榜
	v1.GET("/github_trending", controller.GithubTrendingHandler) // Github热榜

	// 中间件
	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件
	{
		v1.POST("/post", controller.CreatePostHandler) // 创建帖子

		v1.POST("/vote", controller.VoteHandler) // 投票

		v1.POST("/comment", controller.CommentHandler)    // 评论
		v1.GET("/comment", controller.CommentListHandler) // 评论列表

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	pprof.Register(r) // 注册pprof相关路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
