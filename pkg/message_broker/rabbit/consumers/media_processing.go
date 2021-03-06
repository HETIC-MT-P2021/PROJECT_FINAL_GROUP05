package consumers

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit"
)

func MediaProcessingConsumer() {
	err, chanel, queue := rabbit.ConnectToRabbitMQ(rabbit.MEDIA_PROCESSING_QUEUE_NAME)
	if err != nil {
		log.Println(err)
		return
	}

	action := NewMediaProcessingAction()

	rabbitImpl := rabbit.NewRabbitRepository(chanel, queue)
	rabbitImpl.Consume(action)
}