package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/config"
	"github.com/gin-gonic/gin"
)

var GlobalConfig = new(config.Email)
var Group *gin.RouterGroup
