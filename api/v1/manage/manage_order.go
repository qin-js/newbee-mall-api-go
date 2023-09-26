package manage

import (
	"gee"

	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
)

type ManageOrderApi struct {
}

// CheckDoneOrder 发货
func (m *ManageOrderApi) CheckDoneOrder(c *gee.Context) {
	var IDS request.IdsReq
	_ = c.BindJSON(&IDS)
	if err := mallOrderService.CheckDone(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// CheckOutOrder 出库
func (m *ManageOrderApi) CheckOutOrder(c *gee.Context) {
	var IDS request.IdsReq
	_ = c.BindJSON(&IDS)
	if err := mallOrderService.CheckOut(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// CloseOrder 出库
func (m *ManageOrderApi) CloseOrder(c *gee.Context) {
	var IDS request.IdsReq
	_ = c.BindJSON(&IDS)
	if err := mallOrderService.CloseOrder(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMallOrder 用id查询MallOrder
func (m *ManageOrderApi) FindMallOrder(c *gee.Context) {
	id := c.Param("orderId")
	if err, newBeeMallOrderDetailVO := mallOrderService.GetMallOrder(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(newBeeMallOrderDetailVO, c)
	}
}

// GetMallOrderList 分页获取MallOrder列表
func (m *ManageOrderApi) GetMallOrderList(c *gee.Context) {
	var pageInfo request.PageInfo
	_ = c.BindQuery(&pageInfo)
	orderNo := c.Query("orderNo")
	orderStatus := c.Query("orderStatus")
	if err, list, total := mallOrderService.GetMallOrderInfoList(pageInfo, orderNo, orderStatus); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
