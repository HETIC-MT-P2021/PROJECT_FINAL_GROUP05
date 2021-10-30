package producers

import (
	"encoding/json"
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

func JobCheckerProducer(message models.JobCheckerQueueMessage) {
	err, channel, queue := rabbit.ConnectToRabbitMQ(rabbit.JOB_CHECKER_QUEUE_NAME)
	if err != nil {
		log.Println(err)
		return
	}

	rabbitImpl := rabbit.NewRabbitRepository(channel, queue)
	buffer, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	rabbitImpl.Publish(string(buffer))
}