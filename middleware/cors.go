package middleware

import (
	"fmt"
	"gee"
	"net/http"
)

// 处理跨域请求,支持options访问
func Cors() gee.HandlerFunc {
	return func(c *gee.Context) {
		method := c.Req.Method
		origin := c.Req.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id,X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		fmt.Println("cors here")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			fmt.Println("options here")
			c.Fail(http.StatusNoContent, "options")
			// c.FailWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
