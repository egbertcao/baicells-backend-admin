// 设备管理模块，负责设备的备案
package deviceManagement

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/router"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/utils/mongorapper"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/utils/mqttrapper"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/utils/rabbitmq"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type deviceManagementPlugin struct{}

func CreatedeviceManagement() *deviceManagementPlugin {
	sysconfig := parameterInit()
	rabbitConnection := rabbitmq.RabbitNew(&sysconfig.Rabbit, "basic_queue")
	global.RabbitProducer = rabbitmq.ProducerNew(rabbitConnection)
	consumer := rabbitmq.ConsumerNew(rabbitConnection)
	go consumer.Stream()

	global.MongoSession = mongorapper.MongoNew(sysconfig)
	go mqttrapper.MqttNew(sysconfig)

	return &deviceManagementPlugin{}
}

func (*deviceManagementPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitDeviceRouter(group)
	router.RouterGroupApp.InitSecurityRouter(group)
}

func (*deviceManagementPlugin) RouterPath() string {
	return "deviceManagement"
}

func parameterInit() *config.Config {
	var BasicConfig = new(config.Config)
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../server/plugin/deviceManagement/")
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	if err := v.Unmarshal(BasicConfig); err != nil {
		panic(err.Error())
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config.yaml has been changed")
		if err := v.Unmarshal(BasicConfig); err != nil {
			panic(err.Error())
		}
	})
	return BasicConfig
}
