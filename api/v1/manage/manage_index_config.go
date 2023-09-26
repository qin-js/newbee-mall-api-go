package manage

import (
	"gee"
	"strconv"

	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
	manageReq "main.go/model/manage/request"
)

type ManageIndexConfigApi struct {
}

func (m *ManageIndexConfigApi) CreateIndexConfig(c *gee.Context) {
	var req manageReq.MallIndexConfigAddParams
	_ = c.BindJSON(&req)
	if err := mallIndexConfigService.CreateMallIndexConfig(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (m *ManageIndexConfigApi) DeleteIndexConfig(c *gee.Context) {
	var ids request.IdsReq
	_ = c.BindJSON(&ids)
	if err := mallIndexConfigService.DeleteMallIndexConfig(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (m *ManageIndexConfigApi) UpdateIndexConfig(c *gee.Context) {
	var req manageReq.MallIndexConfigUpdateParams
	_ = c.BindJSON(&req)
	if err := mallIndexConfigService.UpdateMallIndexConfig(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (m *ManageIndexConfigApi) FindIndexConfig(c *gee.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err, mallIndexConfig := mallIndexConfigService.GetMallIndexConfig(uint(id)); err != nil {
		global.GVA_LOG.Error("查询失败!"+err.Error(), zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(mallIndexConfig, c)
	}
}

func (m *ManageIndexConfigApi) GetIndexConfigList(c *gee.Context) {
	var pageInfo manageReq.MallIndexConfigSearch
	_ = c.BindQuery(&pageInfo)
	if err, list, total := mallIndexConfigService.GetMallIndexConfigInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
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
