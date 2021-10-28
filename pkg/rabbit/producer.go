package rabbit

import (
	"encoding/json"
	"log"
	"time"
)

func InitProducer() {
	err := ConnectToRabbitMQ()
	if err != nil {
		log.Println(err)
		return
	}

	rabbitImpl := NewRabbitRepository(RabbitChannel, RabbitQueue)

	message := QueueMessage{
		Type:      "audio",
		MediaLink: "youtube.com",
		Options: Options{
			StartInSeconds:    "10",
			DurationInSeconds: "10",
		},
	}
	buffer, err := json.Marshal(message)
	if err != nil {
		return
	}

	time.Sleep(10 * time.Second) // 10s to be sure rabbit is ready

	rabbitImpl.Publish(string(buffer))
}
