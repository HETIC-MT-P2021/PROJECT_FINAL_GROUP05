package discord

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

func DiscordCommandHandler(userID string, channelID string, content string) (string, bool) {
	message := models.DownloadQueueMessage{
		Type: "audio",
		MediaLink: "youtube.com",
		Options: 	models.Options{
			StartInSeconds: "20",
			DurationInSeconds: "35",
		},
	}
	producers.DownloadProducer(message)
	
	switch content {
	case "ping":
		result := Ping(content)
		return result, true
	case "pong":
		result := Ping(content)
		return result, true
	case "carlos":
		result := Ping(content)
		return result, true
	case "contrecarlos":
		result := Ping(content)
		return result, true
	default:
		result := ""
		return result, false
	}
}
