package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitClient struct {
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
	msgs            <-chan amqp.Delivery
	queue           *amqp.Queue
}

// New creates a new consumer state instance, and automatically
// attempts to connect to the server.
func RabbitClientNew(param *config.RabbitMq, quequeName string) *RabbitClient {
	name := quequeName
	addr := config.RabbitMqDns(param)
	session := RabbitClient{
		name: name,
		done: make(chan bool),
	}
	go session.handleReconnect(addr)
	session.RabbitReady = make(chan bool)
	return &session
}

// handleReconnect will wait for a connection error on
// notifyConnClose, and then continuously attempt to reconnect.
func (session *RabbitClient) handleReconnect(addr string) {
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

func (session *RabbitClient) handleReInit(conn *amqp.Connection) bool {
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
func (session *RabbitClient) Queueinit(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	err = ch.Confirm(false)
	if err != nil {
		return err
	}
	queue, err := ch.QueueDeclare(
		"rpc_reply",
		false, // Durable
		false, // Delete when unused
		true,  // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

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

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (session *RabbitClient) Push(data []byte) (err error) {
	corrId := randomString(32)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgs, err := session.channel.Consume(
		session.queue.Name, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		return err
	}
	session.msgs = msgs
	fmt.Println(session.name)
	err = session.channel.PublishWithContext(ctx,
		"",           // exchange
		session.name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       session.queue.Name,
			Body:          data,
		})
	if err != nil {
		mainGlobal.GVA_LOG.Error("Channel not Ready", zap.Error(err))
		return
	}

	go func() {
		for d := range session.msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	// fmt.Println("Published, Await for reply...")

	// for d := range session.msgs {
	// 	if corrId == d.CorrelationId {
	// 		fmt.Println(d.CorrelationId)
	// 		fmt.Println(string([]byte(d.Body)))
	// 		break
	// 	}
	// }
	// fmt.Println("Received end.")
	return
}

// Close will cleanly shutdown the channel and connection.
func (session *RabbitClient) Close() error {
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


