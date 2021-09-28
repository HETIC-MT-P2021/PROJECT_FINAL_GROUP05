package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

// Publish new message on rabbitmq
func (rabbit *RabbitRepositoryImpl) Publish(message string) error {
	err := rabbit.Chan.Publish(
		"",
		rabbit.Queue.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	return err
}

// Consume Receives message and make process
func (rabbit *RabbitRepositoryImpl) Consume() {
	msgs, err := rabbit.Chan.Consume(
		rabbit.Queue.Name, // queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(err)
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// make process
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}