package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("rabbitmq producer")

	conn, err := amqp.Dial("amqps://username:password@url/")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer conn.Close()

	fmt.Println("connected to the rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello RabbitMq"),
		},
	)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("published message")
}
