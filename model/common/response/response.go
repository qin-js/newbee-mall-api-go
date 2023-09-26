package response

import (
	"fmt"
	"gee"
	"net/http"
)

type Response struct {
	ResultCode int         `json:"resultCode"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"message"`
}

const (
	ERROR   = 500
	SUCCESS = 200
	UNLOGIN = 416
)

func Result(code int, data interface{}, msg string, c *gee.Context) {
	// 开始时间
	fmt.Println("%v, %v", code, msg)
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gee.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gee.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gee.Context) {
	Result(SUCCESS, data, "SUCCESS", c)
}

func OkWithDetailed(data interface{}, message string, c *gee.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gee.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gee.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gee.Context) {
	Result(ERROR, data, message, c)
}

func UnLogin(data interface{}, c *gee.Context) {
	Result(UNLOGIN, data, "未登录！", c)
}
