package mall

import (
	"gee"

	v1 "main.go/api/v1"
)

type MallGoodsCategoryIndexRouter struct {
}

func (m *MallGoodsInfoIndexRouter) InitMallGoodsCategoryIndexRouter(Router *gee.RouterGroup) {
	mallGoodsRouter := Router.Group("/v1")
	var mallGoodsCategoryApi = v1.ApiGroupApp.MallApiGroup.MallGoodsCategoryApi
	{
		mallGoodsRouter.GET("/categories", mallGoodsCategoryApi.GetGoodsCategory) // 获取分类数据
	}
}
