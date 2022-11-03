package rabbitmq

import (
	"log"

	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

type Consumer struct {
	session *RabbitSession
}

func ConsumerNew(session *RabbitSession) *Consumer {
	consumer := Consumer{
		session: session,
	}
	return &consumer
}

func (c *Consumer) Stream() {
	tmp := make(chan bool, 1)
	for {
		ready := <-c.session.RabbitReady
		if ready {
			mainGlobal.GVA_LOG.Info("Session is Ready.")
			msgs, err := c.session.channel.Consume(
				c.session.name,
				"",    // Consumer
				false, // Auto-Ack
				false, // Exclusive
				false, // No-local
				false, // No-Wait
				nil,   // Args
			)
			if err != nil {
				mainGlobal.GVA_LOG.Error("Channel not Ready", zap.Error(err))
				continue
			}

			go func(<-chan bool) {
				mainGlobal.GVA_LOG.Info("Consumer is Up")
				for {
					select {
					case d := <-msgs:
						log.Printf(" [x] %s", d.Body)
						d.Ack(false)
					case <-tmp:
						mainGlobal.GVA_LOG.Error("Consumer is Down!")
						return
					}
				}
			}(tmp)
		} else {
			mainGlobal.GVA_LOG.Error("Session not Ready.")
			tmp <- false
		}
	}
}
