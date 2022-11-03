package rabbitmq

import (
	"time"

	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Producer struct {
	session *RabbitSession
}

func ProducerNew(session *RabbitSession) *Producer {
	producer := Producer{
		session: session,
	}
	return &producer
}

func (producer *Producer) Push(data []byte) error {
	for {
		err := producer.session.channel.Publish(
			"",                    // Exchange
			producer.session.name, // Routing key
			false,                 // Mandatory
			false,                 // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        data,
			},
		)
		if err != nil {
			mainGlobal.GVA_LOG.Error("Push failed. Retrying...", zap.Error(err))
			select {
			case <-producer.session.done:
				return errShutdown
			case <-time.After(resendDelay):
			}
			continue
		}
		select {
		case confirm := <-producer.session.notifyConfirm:
			if confirm.Ack {
				return nil
			}
		case <-time.After(resendDelay):
		}
		mainGlobal.GVA_LOG.Error("Push didn't confirm. Retrying...")
	}
}
