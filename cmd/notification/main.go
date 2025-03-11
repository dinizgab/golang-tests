package main 

import (
	"log"

	"github.com/dinizgab/golang-tests/internal/service"
)

const notificationTopic = "notification"

func main() {
	conn, err := service.NewBrokerConnection()
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.Consume(
		notificationTopic,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var forever chan struct{}
	go func() {
		for msg := range q {
			log.Println("Received message")
			log.Println(string(msg.Body))
		}
	}()

	log.Println("Waiting for messages...")
	<-forever
}
