package logger

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"net/http"
	"time"
)

type RequestLogger struct {
	recorder *RequestRecorder
}

func NewRequestLogger(r *RequestRecorder) *RequestLogger {
	rl := &RequestLogger{recorder: r}
	return rl
}

func (rl *RequestLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	log.WithFields(log.Fields{"Path": r.URL.Path, "Query": r.URL.RawQuery}).Info("Executing request pipeline")

	t := time.Now()
	next(rw, r)

	// look at timer
	elapsed := time.Since(t)
	// add to to in-mem storage
	record := Record{r.RequestURI, elapsed}
	rl.recorder.add(record)
	// put on queue
	send(record)
	// log
	log.WithFields(log.Fields{"Elapsed time": elapsed}).Info("Executed request in %v", elapsed)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func send(record Record) {
	rabbit := GetEnv("RABBIT_CONNECTION", "amqp://guest:guest@localhost:5672/")
	log.Info("%s", rabbit)
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

	body, err := json.Marshal(record)

	failOnError(err, "failed to serialise message to json")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
