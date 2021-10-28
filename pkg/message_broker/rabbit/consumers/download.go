package consumers

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

func DownloadConsumer() {
	err, chanel, queue := rabbit.ConnectToRabbitMQ(rabbit.DOWLOAD_QUEUE_NAME)
	if err != nil {
		log.Println(err)
		return
	}

	message := models.MediaProcessingQueueMessage{
		FileName: "13.mp4",
	}

	action := NewCallMediaProcessingProducer(message)

	rabbitImpl := rabbit.NewRabbitRepository(chanel, queue)
	rabbitImpl.Consume(action)
}