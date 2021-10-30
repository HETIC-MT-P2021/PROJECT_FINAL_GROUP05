package main

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/consumers"
)

func main() {
	err := commands.InitCommandBus()
	if err != nil {
		log.Println(err)
		return
	}

	consumers.MediaProcessingConsumer()
}