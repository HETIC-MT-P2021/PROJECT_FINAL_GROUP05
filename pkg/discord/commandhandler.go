package discord

import (
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/parsings"
	"log"
)

var (
	message models.DownloadQueueMessage
)

func CommandHandler(userID string, channelID string, content string) (string, bool) {

	log.Println(userID)
	log.Println(channelID)
	log.Println(content)

	message = parsings.ParseCommand(content)

	producers.DownloadProducer(message)
	log.Printf("%#v", message)

	result := fmt.Sprintf("%#v", message)
	return result, true
}
