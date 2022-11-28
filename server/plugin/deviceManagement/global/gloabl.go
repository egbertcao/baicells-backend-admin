package global

import (
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/utils/mongorapper"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/utils/rabbitmq"
	"github.com/gin-gonic/gin"
)

var (
	MongoSession   *mongorapper.MongoSession
	MqttConnection *autopaho.ConnectionManager
	RabbitClient   *rabbitmq.RabbitClient
	Group          *gin.RouterGroup
)
