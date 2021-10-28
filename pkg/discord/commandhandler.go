package discord

import (
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/parsings"
	"log"
)

var (
	message parsings.QueueMessage
)

func CommandHandler(userID string, channelID string, content string) (string, bool) {
	log.Println(userID)
	log.Println(channelID)
	log.Println(content)

	message = parsings.ParseCommand(content)
	fmt.Println(message)

	result := fmt.Sprintf("%#v", message)
	return result, true
}
