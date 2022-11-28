package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InstanceApi struct {
}

var instanceService = service.ServiceGroupApp.InstanceService

func (api *InstanceApi) CallInstance(c *gin.Context) {
	var device model.Instance
	_ = c.ShouldBindJSON(&device)
	if err := instanceService.CallMethod(device); err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
