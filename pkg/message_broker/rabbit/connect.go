package rabbit

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/streadway/amqp"
)

// rabbitMqEnv contains rabbitmq env credentials
type rabbitMqEnv struct {
	RabbitMqHost string `env:"RABBITMQ_HOST"`
	RabbitMqPort string `env:"RABBITMQ_PORT"`
	RabbitMqUser string `env:"RABBITMQ_DEFAULT_USER"`
	RabbitMqPass string `env:"RABBITMQ_DEFAULT_PASS"`
}

const (
	numberOftries       = 10
	timeToWaitInSeconds = 5
	AMQP_PROTOCOL 			= "amqp"
)

// ConnectToRabbitMQ is for connecting to rabbitmq
func ConnectToRabbitMQ(queueName string) (error, *amqp.Channel, amqp.Queue) {
	cfg := rabbitMqEnv{}
	if err := env.Parse(&cfg); err != nil {
		return err, nil, amqp.Queue{}
	}

	urlConn := fmt.Sprintf("%s://%s:%s@%s:%s/",
		AMQP_PROTOCOL,
		cfg.RabbitMqPass,
		cfg.RabbitMqUser,
		cfg.RabbitMqHost,
		cfg.RabbitMqPort)

	var rabbitConnection *amqp.Connection
	var err error
	for index := 0; index < numberOftries; index++ {
		rabbitConnection, err = amqp.Dial(urlConn)
		if err != nil {
			time.Sleep(timeToWaitInSeconds * time.Second)
		} else {
			break
		}
	}

	ch, err := rabbitConnection.Channel()
	if err != nil {
		return err, nil, amqp.Queue{}
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err, nil, amqp.Queue{}
	}

	return nil, ch, q
}