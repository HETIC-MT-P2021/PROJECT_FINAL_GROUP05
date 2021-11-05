package consumers

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit"
)

func DownloadConsumer() {
	err, channel, queue := rabbit.ConnectToRabbitMQ(rabbit.DOWLOAD_QUEUE_NAME)
	if err != nil {
		log.Println(err)
		return
	}

	action := NewCallMediaProcessingProducer()

	rabbitImpl := rabbit.NewRabbitRepository(channel, queue)
	rabbitImpl.Consume(action)
}
