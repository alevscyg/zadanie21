package rabbitmq

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/alevscyg/zadanie23/internal/handler"
	sqlstorage "github.com/alevscyg/zadanie23/storage/sqlStorage"
	"github.com/rabbitmq/amqp091-go"
)

type Data struct {
	UserId      int
	Title       string
	Description string
	Value       int
}

type ToDo struct {
	Pattern string
	Data    Data
	Id      string
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func RabbitMqStart(RABBITMQ_URL string, RABBITMQ_TODO_QUEUE string, DatabaseURL string) {
	fmt.Println("Consumer app")

	db, err := newDB(DatabaseURL)
	failOnError(err, "Failed to connect to Database")
	defer db.Close()
	store := sqlstorage.New(db)

	conn, err := amqp091.Dial(RABBITMQ_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RABBITMQ_TODO_QUEUE, // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name,              // queue
		RABBITMQ_TODO_QUEUE, // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			toDo := ToDo{}
			err := json.Unmarshal([]byte(d.Body), &toDo)
			failOnError(err, "Failed to publish a message")
			if toDo.Pattern == "createToDo" {
				result, err := handler.CreateToDo(toDo.Data.UserId, toDo.Data.Title, toDo.Data.Description, store)
				failOnError(err, "Failed RequestHandler")
				err = ch.PublishWithContext(ctx,
					"",        // exchange
					d.ReplyTo, // routing key
					false,     // mandatory
					false,     // immediate
					amqp091.Publishing{
						ContentType:   "application/json",
						CorrelationId: d.CorrelationId,
						Body:          []byte(" Id: " + fmt.Sprint(result.Id) + " Userid: " + fmt.Sprint(result.Userid) + " Title: " + result.Title + " Description: " + result.Description),
					})
				failOnError(err, "Failed to publish a message")
			} else if toDo.Pattern == "findToDoById" {
				result, err := handler.FindToDoById(toDo.Data.Value, store)
				failOnError(err, "Failed RequestHandler")
				err = ch.PublishWithContext(ctx,
					"",        // exchange
					d.ReplyTo, // routing key
					false,     // mandatory
					false,     // immediate
					amqp091.Publishing{
						ContentType:   "application/json",
						CorrelationId: d.CorrelationId,
						Body:          []byte(" Id: " + fmt.Sprint(result.Id) + " Userid: " + fmt.Sprint(result.Userid) + " Title: " + result.Title + " Description: " + result.Description),
					})
				failOnError(err, "Failed to publish a message")
			}
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
