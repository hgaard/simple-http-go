package main

import (
	"fmt"
	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"simple-http-server/logger"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	seqURL := logger.GetEnv("SEQ_URL", "http://localhost:5341")
	fmt.Printf("Logging to SEQ_URL '%s'\n", seqURL)
	log.AddHook(logruseq.NewSeqHook(seqURL))
}

func main() {
	rabbit := logger.GetEnv("RABBIT_CONNECTION", "amqp://guest:guest@localhost:5672/")
	log.Info("Rabbit connection to RABBIT_CONNECTION %s", rabbit)

	conn, err := amqp.Dial(rabbit)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
