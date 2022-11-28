package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model"
	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TemplateApi struct {
}

var templateService = service.ServiceGroupApp.TemplateService

func (api *TemplateApi) CreateTemplate(c *gin.Context) {
	var device model.Template
	_ = c.ShouldBindJSON(&device)
	if err := templateService.CreateTemplate(device); err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTemplate 删除Template
// @Tags Template
// @Summary 删除Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.Template true "删除Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /device/deleteTemplate [delete]
func (deviceApi *TemplateApi) DeleteTemplate(c *gin.Context) {
	var device model.Template
	err := c.ShouldBindJSON(&device)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.DeleteTemplate(device); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTemplateByIds 批量删除Template
// @Tags Template
// @Summary 批量删除Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /device/deleteTemplateByIds [delete]
func (deviceApi *TemplateApi) DeleteTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.DeleteTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTemplate 更新Template
// @Tags Template
// @Summary 更新Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.Template true "更新Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /device/updateTemplate [put]
func (deviceApi *TemplateApi) UpdateTemplate(c *gin.Context) {
	var device model.Template
	err := c.ShouldBindJSON(&device)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.UpdateTemplate(device); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTemplate 用id查询Template
// @Tags Template
// @Summary 用id查询Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query device.Template true "用id查询Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /device/findTemplate [get]
func (deviceApi *TemplateApi) FindTemplate(c *gin.Context) {
	var device model.Template
	err := c.ShouldBindQuery(&device)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redevice, err := templateService.GetTemplate(device.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redevice": redevice}, c)
	}
}

// GetTemplateList 分页获取Template列表
// @Tags Template
// @Summary 分页获取Template列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query deviceReq.TemplateSearch true "分页获取Template列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/getTemplateList [get]
func (deviceApi *TemplateApi) GetTemplateList(c *gin.Context) {
	var pageInfo deviceReq.TemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := templateService.GetTemplateInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
