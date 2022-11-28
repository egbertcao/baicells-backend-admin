package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/api"
	"github.com/gin-gonic/gin"
)

type TemplateRouter struct{}

func (s *TemplateRouter) InitTemplateRouter(Router *gin.RouterGroup) {
	deviceRouter := Router.Group("template").Use(middleware.OperationRecord())
	deviceRouterWithoutRecord := Router.Group("template")
	var deviceApi = api.ApiGroupApp.TemplateApi
	{
		deviceRouter.POST("createTemplate", deviceApi.CreateTemplate)             // 新建template
		deviceRouter.DELETE("deleteTemplate", deviceApi.DeleteTemplate)           // 删除template
		deviceRouter.DELETE("deleteTemplateByIds", deviceApi.DeleteTemplateByIds) // 批量删除template
		deviceRouter.PUT("updateTemplate", deviceApi.UpdateTemplate)              // 更新template
	}
	{
		deviceRouterWithoutRecord.GET("findTemplate", deviceApi.FindTemplate)       // 根据ID获取template
		deviceRouterWithoutRecord.GET("getTemplateList", deviceApi.GetTemplateList) // 获取template列表
	}
}
