package rabbitmq

import (
	"context"
	"log"
	"time"

	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitServer struct {
	name string
	//logger          *log.Logger
	connection      *amqp.Connection
	channel         *amqp.Channel
	done            chan bool
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	RabbitReady     chan bool
	isReady         bool
	queue           *amqp.Queue
	msgs            <-chan amqp.Delivery
	onMsg           func(amqp.Delivery)
}

// New creates a new consumer state instance, and automatically
// attempts to connect to the server.
func RabbitServerNew(param *config.RabbitMq, quequeName string, onmsg func(amqp.Delivery)) *RabbitServer {
	name := quequeName
	addr := config.RabbitMqDns(param)
	session := RabbitServer{
		name:  name,
		done:  make(chan bool),
		onMsg: onmsg,
	}
	go session.handleReconnect(addr)
	session.RabbitReady = make(chan bool)
	go session.Stream()
	return &session
}

// handleReconnect will wait for a connection error on
// notifyConnClose, and then continuously attempt to reconnect.
func (session *RabbitServer) handleReconnect(addr string) {
	for {
		session.isReady = false
		mainGlobal.GVA_LOG.Info("Attempting to connect")
		conn, err := amqp.Dial(addr)
		if err != nil {
			mainGlobal.GVA_LOG.Error("Failed to connect. Retrying...", zap.Error(err))
			select {
			case <-session.done:
				return
			case <-time.After(ReconnectDelay):
			}
			continue
		}
		session.notifyConnClose = make(chan *amqp.Error)
		conn.NotifyClose(session.notifyConnClose)
		mainGlobal.GVA_LOG.Info("Connected!")
		if done := session.handleReInit(conn); done {
			break
		}
	}
}

func (session *RabbitServer) handleReInit(conn *amqp.Connection) bool {
	for {
		session.isReady = false
		err := session.Queueinit(conn)
		if err != nil {
			mainGlobal.GVA_LOG.Info("Failed to initialize channel. Retrying...")
			select {
			case <-session.done:
				return true
			case <-time.After(ReInitDelay):
			}
			continue
		}

		select {
		case <-session.done:
			return true
		case <-session.notifyConnClose:
			mainGlobal.GVA_LOG.Error("Connection closed. Reconnecting...")
			session.RabbitReady <- false
			return false
		case <-session.notifyChanClose:
			session.RabbitReady <- false
			mainGlobal.GVA_LOG.Error("Channel closed. Re-running init...")
		}
	}
}

// init will initialize channel & declare queue
func (session *RabbitServer) Queueinit(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	// err = ch.Confirm(false)
	// if err != nil {
	// 	return err
	// }
	queue, err := ch.QueueDeclare(
		session.name,
		false, // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",    // Consumer
		false, // Auto-Ack
		false, // Exclusive
		false, // No-local
		false, // No-Wait
		nil,   // Args
	)
	if err != nil {
		mainGlobal.GVA_LOG.Error("Channel not Ready", zap.Error(err))
		return err
	}

	session.msgs = msgs
	session.queue = &queue
	session.channel = ch
	session.notifyChanClose = make(chan *amqp.Error)
	session.notifyConfirm = make(chan amqp.Confirmation, 1)
	session.channel.NotifyClose(session.notifyChanClose)
	session.channel.NotifyPublish(session.notifyConfirm)
	session.isReady = true
	session.RabbitReady <- true
	mainGlobal.GVA_LOG.Info("Setup!")
	return nil
}

func (c *RabbitServer) Stream() {
	tmp := make(chan bool, 1)
	for {
		ready := <-c.RabbitReady
		if ready {
			mainGlobal.GVA_LOG.Info("Session is Ready.")

			go func(<-chan bool) {
				mainGlobal.GVA_LOG.Info("Consumer is Up")
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				for {
					select {
					case d := <-c.msgs:
						log.Printf(" [x] %s", d.Body)

						c.onMsg(d)
						//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
						// defer cancel()
						err := c.channel.PublishWithContext(
							ctx,
							"",        // exchange
							d.ReplyTo, // routing key
							false,     // mandatory
							false,     // immediate
							amqp.Publishing{
								ContentType:   "text/plain",
								CorrelationId: d.CorrelationId,
								Body:          []byte("123456"),
							})
						if err != nil {
							mainGlobal.GVA_LOG.Error("Channel not Ready", zap.Error(err))
							return
						}

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

// Close will cleanly shutdown the channel and connection.
func (session *RabbitServer) Close() error {
	if !session.isReady {
		return ErrAlreadyClosed
	}
	err := session.channel.Close()
	if err != nil {
		return err
	}
	err = session.connection.Close()
	if err != nil {
		return err
	}
	close(session.done)
	session.isReady = false
	return nil
}
