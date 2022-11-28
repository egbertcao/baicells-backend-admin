package backend

import (
	"encoding/json"
	"fmt"

	"github.com/eclipse/paho.golang/paho"
	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/backend/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/backend/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/utils/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

var LoginService = service.ServiceGroupApp.LoginService

func RabbitConsumerReceive(d amqp.Delivery) {
	mainGlobal.GVA_LOG.Debug(string(d.Body[:]))
	mainGlobal.GVA_LOG.Debug(d.CorrelationId)
	mainGlobal.GVA_LOG.Debug(d.ReplyTo)
	var msg model.Msg
	var request model.Login
	json.Unmarshal(d.Body, &msg)
	json.Unmarshal(msg.Data, &request)
	switch request.Command {
	case "device_login_req":
		LoginService.LoginRequest(request)
	}
}

func MqttReceiver(m *paho.Publish) {
	fmt.Println(m.String())
	mainGlobal.GVA_LOG.Debug(string(m.Payload[:]))
	msg := model.Msg{
		ResponseTopic:   m.Properties.ResponseTopic,
		CorrelationData: string(m.Properties.CorrelationData[:]),
		Data:            m.Payload,
	}
	data_bytes, _ := json.Marshal(msg)
	go rabbitmq.RabbitRPC(data_bytes)
}

func HttpDespose() {

}
