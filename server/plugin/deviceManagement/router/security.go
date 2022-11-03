package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/api"
	"github.com/gin-gonic/gin"
)

type SecurityRouter struct{}

func (s *SecurityRouter) InitSecurityRouter(Router *gin.RouterGroup) {
	mqttRouter := Router.Group("mqtt").Use(middleware.OperationRecord())
	mqttRouterWithoutRecord := Router.Group("mqtt")
	var authenticationApi = api.ApiGroupApp.AuthenticationApi
	{
		mqttRouter.POST("createAuthentication", authenticationApi.CreateAuthentication)             // 新建Authentication
		mqttRouter.DELETE("deleteAuthentication", authenticationApi.DeleteAuthentication)           // 删除Authentication
		mqttRouter.DELETE("deleteAuthenticationByIds", authenticationApi.DeleteAuthenticationByIds) // 批量删除Authentication
		mqttRouter.PUT("updateAuthentication", authenticationApi.UpdateAuthentication)              // 更新Authentication
	}
	{
		mqttRouterWithoutRecord.GET("findAuthentication", authenticationApi.FindAuthentication)       // 根据ID获取Authentication
		mqttRouterWithoutRecord.GET("getAuthenticationList", authenticationApi.GetAuthenticationList) // 获取Authentication列表
	}

	var authorizationApi = api.ApiGroupApp.AuthorizationApi
	{
		mqttRouter.POST("createAuthorization", authorizationApi.CreateAuthorization)             // 新建Authorization
		mqttRouter.DELETE("deleteAuthorization", authorizationApi.DeleteAuthorization)           // 删除Authorization
		mqttRouter.DELETE("deleteAuthorizationByIds", authorizationApi.DeleteAuthorizationByIds) // 批量删除Authorization
		mqttRouter.PUT("updateAuthorization", authorizationApi.UpdateAuthorization)              // 更新Authorization
	}
	{
		mqttRouterWithoutRecord.GET("findAuthorization", authorizationApi.FindAuthorization)       // 根据ID获取Authorization
		mqttRouterWithoutRecord.GET("getAuthorizationList", authorizationApi.GetAuthorizationList) // 获取Authorization列表
	}
}
