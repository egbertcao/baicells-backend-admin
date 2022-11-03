package backend

import (
	"encoding/json"

	"github.com/eclipse/paho.golang/paho"
	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model"
)

func MessageReceiver(m *paho.Publish) {
	mainGlobal.GVA_LOG.Info(string(m.Payload[:]))
	msg := model.Msg{
		ResponseTopic:   m.Properties.ResponseTopic,
		CorrelationData: string(m.Properties.CorrelationData[:]),
		Data:            m.Payload,
	}
	data_bytes, _ := json.Marshal(msg)
	global.RabbitProducer.Push(data_bytes)
}

func RabbitConsumerReceive(data []byte) {

}
