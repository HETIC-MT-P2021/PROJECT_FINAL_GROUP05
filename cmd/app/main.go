package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/rabbit"
)

func main() {
	err := rabbit.ConnectToRabbitMQ()
	if err != nil {
		log.Println(err)
		return
	}

	rabbitImpl := rabbit.NewRabbitRepository(rabbit.RabbitChannel, rabbit.RabbitQueue)
	
	message := rabbit.QueueMessage{
		Type: "audio",
		MediaLink: "youtube.com",
		Options: 	rabbit.Options{
			StartInSeconds: "10",
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