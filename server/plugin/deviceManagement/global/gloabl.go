package global

import (
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/utils/mongorapper"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/utils/rabbitmq"
)

var (
	MongoSession     *mongorapper.MongoSession
	RabbitProducer   *rabbitmq.Producer
	MqttConnection   *autopaho.ConnectionManager
	RabbitConnection *rabbitmq.RabbitSession
)
