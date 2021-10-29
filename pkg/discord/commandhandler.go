package discord

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/parsings"
	"github.com/bwmarrin/discordgo"
)

var (
	message models.DownloadQueueMessage
)

const LOADER_URL = "https://c.tenor.com/I6kN-6X7nhAAAAAj/loading-buffering.gif"

func CommandHandler(s *discordgo.Session, userID string, channelID string, content string) error {
	message = parsings.ParseCommand(content)

	msg := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: LOADER_URL,
			},
			Title: "Carlos is processing 2",
			Description: 
				"Media :  \n" +
				"State : In progress",
		},
	}

	discordMessage, err := s.ChannelMessageSendComplex(channelID, msg)
	if err != nil {
		return err
	}
	discordInfo := models.JobCheckerQueueMessage{
		ChannelID: discordMessage.ChannelID,
		MessageID: discordMessage.ID,
	}
	message.DiscordInfo = discordInfo
	producers.DownloadProducer(message)

	return err
}
