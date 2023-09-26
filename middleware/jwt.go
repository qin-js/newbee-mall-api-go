package middleware

import (
	"fmt"
	"gee"
	"time"

	"main.go/model/common/response"
	"main.go/service"
)

var manageAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserTokenService
var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService

func AdminJWTAuth() gee.HandlerFunc {
	return func(c *gee.Context) {
		token := c.Req.Header.Get("token")
		fmt.Println("jwt here")
		if token == "" {
			response.FailWithDetailed(nil, "未登录或非法访问", c)
			c.Fail(100, "未登录或非法访问")
			return
		}
		err, mallAdminUserToken := manageAdminUserTokenService.ExistAdminToken(token)
		if err != nil {
			response.FailWithDetailed(nil, "未登录或非法访问", c)
			c.Fail(100, "未登录或非法访问")
			return
		}
		if time.Now().After(mallAdminUserToken.ExpireTime) {
			response.FailWithDetailed(nil, "授权已过期", c)
			err = manageAdminUserTokenService.DeleteMallAdminUserToken(token)
			if err != nil {
				return
			}
			c.Fail(100, "授权已过期")
			return
		}
		c.Next()
	}

}

func UserJWTAuth() gee.HandlerFunc {
	return func(c *gee.Context) {
		fmt.Println("jwt here")
		token := c.Req.Header.Get("token")
		if token == "" {
			fmt.Println("no token")
			response.UnLogin(nil, c)
			c.Fail(100, "未登录或非法访问")
			return
		}
		err, mallUserToken := mallUserTokenService.ExistUserToken(token)
		if err != nil {
			fmt.Println("token doesn't exist")
			response.UnLogin(nil, c)
			c.Fail(100, "未登录或非法访问")
			return
		}
		if time.Now().After(mallUserToken.ExpireTime) {
			fmt.Println("token guoqi")
			response.FailWithDetailed(nil, "授权已过期", c)
			err = mallUserTokenService.DeleteMallUserToken(token)
			if err != nil {
				return
			}
			c.Fail(100, "未登录或非法访问")
			return
		}
		c.Next()
	}

}
