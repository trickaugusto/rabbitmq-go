package main

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/trickaugusto/rabbitmq-go/config"
)

func main() {
	connection, channel, err := config.ConnectToRabbitMQ()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	defer channel.Close()

	queue, err := channel.QueueDeclare("hello", false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World"),
	},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Message sent")
	fmt.Println("Successfully published message to RabbitMQ")
}
