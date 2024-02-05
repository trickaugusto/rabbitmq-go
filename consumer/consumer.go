package main

import (
	"fmt"

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

	msgs, err := channel.Consume("hello", "", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
