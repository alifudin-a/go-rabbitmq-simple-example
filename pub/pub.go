package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const url = "amqp://root:root@localhost:5672/"

func main() {
	var exchange = "exchange-1"
	var routingKey = "a.b.c"

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("cannot (re)dial: %v: %q", err, url)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

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

	err = ch.QueueBind(q.Name, routingKey, exchange, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text",
			Body:        []byte("halo"),
		})
	if err != nil {
		log.Fatalln(err)
	}
}
