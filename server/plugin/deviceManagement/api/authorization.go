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

type AuthorizationApi struct {
}

var authorizationService = service.ServiceGroupApp.AuthorizationSerivice

func (api *AuthorizationApi) CreateAuthorization(c *gin.Context) {
	var authorization model.MqttAuthorization
	_ = c.ShouldBindJSON(&authorization)
	if err := authorizationService.CreateAuthorization(authorization); err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDevice 删除Device
// @Tags Device
// @Summary 删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.Device true "删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /device/deleteDevice [delete]
func (authorizationApi *AuthorizationApi) DeleteAuthorization(c *gin.Context) {
	var authorization model.MqttAuthorization
	err := c.ShouldBindJSON(&authorization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorizationService.DeleteAuthorization(authorization); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeviceByIds 批量删除Device
// @Tags Device
// @Summary 批量删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /device/deleteDeviceByIds [delete]
func (authorizationApi *AuthorizationApi) DeleteAuthorizationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorizationService.DeleteAuthorizationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDevice 更新Device
// @Tags Device
// @Summary 更新Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.Device true "更新Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /device/updateDevice [put]
func (authorizationApi *AuthorizationApi) UpdateAuthorization(c *gin.Context) {
	var authorization model.MqttAuthorization
	err := c.ShouldBindJSON(&authorization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorizationService.UpdateAuthorization(authorization); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDevice 用id查询Device
// @Tags Device
// @Summary 用id查询Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query device.Device true "用id查询Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /device/findDevice [get]
func (authorizationApi *AuthorizationApi) FindAuthorization(c *gin.Context) {
	var authorization model.MqttAuthorization
	err := c.ShouldBindQuery(&authorization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redevice, err := authorizationService.GetAuthorization(authorization.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redevice": redevice}, c)
	}
}

// GetDeviceList 分页获取Device列表
// @Tags Device
// @Summary 分页获取Device列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query deviceReq.DeviceSearch true "分页获取Device列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/getDeviceList [get]
func (authorizationApi *AuthorizationApi) GetAuthorizationList(c *gin.Context) {
	var pageInfo deviceReq.MqttAuthorizationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := authorizationService.GetAuthorizationInfoList(pageInfo); err != nil {
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
