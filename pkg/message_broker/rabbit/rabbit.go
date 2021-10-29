package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

type rabbitRepository struct {
	Chan  *amqp.Channel
	Queue amqp.Queue
}

func NewRabbitRepository(channel *amqp.Channel, queue amqp.Queue) *rabbitRepository {
	return &rabbitRepository{
		Chan: channel,
		Queue: queue,
	}
}

// Publish new message on rabbitmq
func (rabbit *rabbitRepository) Publish(message string) error {
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
func (rabbit *rabbitRepository) Consume(action ConsumerAction) {
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

			if action != nil {
				action.SetBody(d.Body)
				err := action.Execute()
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}