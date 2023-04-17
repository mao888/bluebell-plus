package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

// 令牌桶其实和漏桶的原理类似，令牌桶按固定的速率往桶里放入令牌，并且只要能从桶里取出令牌就能通过，令牌桶支持突发流量的快速处理。
// 对于从桶里取不到令牌的场景，我们可以选择等待也可以直接拒绝并返回。
// 对于令牌桶的Go语言实现，大家可以参照github.com/juju/ratelimit库。这个库支持多种令牌桶模式，并且使用起来也比较简单。

// 对于该限流中间件的注册位置，我们可以按照不同的限流策略将其注册到不同的位置，例如：
// 如果要对全站限流就可以注册成全局的中间件。
// 如果是某一组路由需要限流，那么就只需将该限流中间件注册到对应的路由组即可。

// RateLimitMiddleware 创建指定填充速率和容量大小的令牌桶
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}
