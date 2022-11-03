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

type AuthenticationApi struct {
}

var authenticationService = service.ServiceGroupApp.AuthenticationSerivice

func (api *AuthenticationApi) CreateAuthentication(c *gin.Context) {
	var authentication model.MqttAuthentication
	_ = c.ShouldBindJSON(&authentication)
	if err := authenticationService.CreateAuthentication(authentication); err != nil {
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
func (authenticationApi *AuthenticationApi) DeleteAuthentication(c *gin.Context) {
	var authentication model.MqttAuthentication
	err := c.ShouldBindJSON(&authentication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authenticationService.DeleteAuthentication(authentication); err != nil {
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
func (authenticationApi *AuthenticationApi) DeleteAuthenticationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authenticationService.DeleteAuthenticationByIds(IDS); err != nil {
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
func (authenticationApi *AuthenticationApi) UpdateAuthentication(c *gin.Context) {
	var authentication model.MqttAuthentication
	err := c.ShouldBindJSON(&authentication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authenticationService.UpdateAuthentication(authentication); err != nil {
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
func (authenticationApi *AuthenticationApi) FindAuthentication(c *gin.Context) {
	var authentication model.MqttAuthentication
	err := c.ShouldBindQuery(&authentication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redevice, err := authenticationService.GetAuthentication(authentication.ID); err != nil {
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
func (authenticationApi *AuthenticationApi) GetAuthenticationList(c *gin.Context) {
	var pageInfo deviceReq.MqttAuthenticationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := authenticationService.GetAuthenticationInfoList(pageInfo); err != nil {
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
