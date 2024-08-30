package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection

func InitRabbitMQ(url string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	"hello",
	// 	false,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// failOnError(err, "Failed to declare a queue")
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
