package config

import (
	"fmt"

	"github.com/streadway/amqp"
)

func ConnectToRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	fmt.Println("RabbitMQ in Golang: Getting Started")

	connection, err := amqp.Dial("amqp://user:pass@localhost:5672/")
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to connect to RabbitMQ: %w", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open a channel: %w", err)
	}

	return connection, channel, nil
}
