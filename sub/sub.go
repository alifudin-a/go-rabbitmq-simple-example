package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const url = "amqp://root:root@localhost:5672/"

func main() {
	// var exchange = "exchange-1"
	// var routingKey = "a.b.c"

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("cannot (re)dial: %v: %q", err, url)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"queue-1", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalln(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalln(err)
	}

	message := <-msgs
	fmt.Println("message: ", string(message.Body))
}
