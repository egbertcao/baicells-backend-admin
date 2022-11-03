package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/api"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

func (s *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup) {
	deviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	deviceRouterWithoutRecord := Router.Group("device")
	var deviceApi = api.ApiGroupApp.DeviceApi
	{
		deviceRouter.POST("createDevice", deviceApi.CreateDevice)             // 新建Device
		deviceRouter.DELETE("deleteDevice", deviceApi.DeleteDevice)           // 删除Device
		deviceRouter.DELETE("deleteDeviceByIds", deviceApi.DeleteDeviceByIds) // 批量删除Device
		deviceRouter.PUT("updateDevice", deviceApi.UpdateDevice)              // 更新Device
	}
	{
		deviceRouterWithoutRecord.GET("findDevice", deviceApi.FindDevice)       // 根据ID获取Device
		deviceRouterWithoutRecord.GET("getDeviceList", deviceApi.GetDeviceList) // 获取Device列表
	}
}
