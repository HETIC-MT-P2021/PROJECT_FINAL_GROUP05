package consumers

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit"
	"github.com/bwmarrin/discordgo"
)

func JobCheckerConsumer(session *discordgo.Session) {
	err, chanel, queue := rabbit.ConnectToRabbitMQ(rabbit.JOB_CHECKER_QUEUE_NAME)
	if err != nil {
		log.Println(err)
		return
	}

	action := NewJobCheckerAction(session)

	rabbitImpl := rabbit.NewRabbitRepository(chanel, queue)
	rabbitImpl.Consume(action)
}