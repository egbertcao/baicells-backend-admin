package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/api"
	"github.com/gin-gonic/gin"
)

type InstanceRouter struct{}

func (s *InstanceRouter) InitinstanceRouter(Router *gin.RouterGroup) {
	instanceRouter := Router.Group("instance").Use(middleware.OperationRecord())
	var instanceApi = api.ApiGroupApp.InstanceApi
	{
		instanceRouter.POST("createinstance", instanceApi.CallInstance)
	}

}
