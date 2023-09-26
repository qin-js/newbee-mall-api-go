package mall

import (
	"fmt"
	"gee"

	v1 "main.go/api/v1"
)

type MallCarouselIndexRouter struct {
}

func (m *MallCarouselIndexRouter) InitMallCarouselIndexRouter(Router *gee.RouterGroup) {
	mallCarouselRouter := Router.Group("/v1")
	var mallCarouselApi = v1.ApiGroupApp.MallApiGroup.MallIndexApi
	{
		mallCarouselRouter.GET("/index-infos", func(c *gee.Context) {
			fmt.Println("test")
			mallCarouselApi.MallIndexInfo(c) // 获取首页数据
		})
	}
}
