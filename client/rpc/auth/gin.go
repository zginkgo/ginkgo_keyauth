package auth

import (
	"github.com/gin-gonic/gin"
)

// GinAuthHandlerFunc Gin auth Middleware
// 通过 r.Use(<middleware>) 添加中间件
// 		r := gin.New()
// 		r.Use()
func (a *KeyauthAuther) GinAuthHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理Request对象
		// c.Request

		// 处理完成后 需要直接中断返回
		// // 处理Response Header
		// c.Writer.Header()
		// // 处理Response Body
		// c.Writer.Write()
		// return

		// 如果处理ok需要把一些中间结果 然后后面的请求也能方式, 需要把结果保存中 上下文中
		// 原生 c.Request.Context()
		// 如何:  context Get
		// 认证完成后, 我们需要用户名称或者其他信息 传递下去

		// c.Set("username", "xxxx")
		// c.Get("username")

		// 把请求flow到下一站去处理
		//c.Next()
	}
}
