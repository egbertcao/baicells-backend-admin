package rabbitmq

import (
	"context"
	"log"
	"time"

	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func failOnError(err error, msg string) {
	if err != nil {
		mainGlobal.GVA_LOG.Error("Channel not Ready", zap.Error(err))
	}
}

var gch *amqp.Channel
var gqueue amqp.Queue
var gmsgs <-chan amqp.Delivery

func RabbitRpcInit() {
	gconn, err := amqp.Dial("amqp://admin:admin@172.16.4.226:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	gch, err = gconn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()

	gqueue, err = gch.QueueDeclare(
		"rpc_reply", // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // noWait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	gmsgs, err = gch.Consume(
		gqueue.Name, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")
}

func RabbitRPC(data []byte) (res int, err error) {

	corrId := randomString(32)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = gch.PublishWithContext(ctx,
		"",            // exchange
		"basic_queue", // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       gqueue.Name,
			Body:          data,
		})
	failOnError(err, "Failed to publish a message")

	for d := range gmsgs {
		if corrId == d.CorrelationId {
			log.Printf("Received a message: %s", d.Body)
			failOnError(err, "Failed to convert body to integer")
			break
		}
	}

	return
}
