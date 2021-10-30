package carlos

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord/carlos/template"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/parsings"
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	messageForDiscord := template.NewMessageSendWrapper()
	messageForDiscord.SetCarlosIsProcessingMessage()

	discordMessage, err := s.ChannelMessageSendComplex(m.ChannelID, messageForDiscord.MessageSend)
	if err != nil {
		log.Println(err)
	}
	discordInfo := models.JobCheckerQueueMessage{
		ChannelID: discordMessage.ChannelID,
		MessageID: discordMessage.ID,
	}

	messageForRabbit := parsings.ParseCommand(m.Content)
	messageForRabbit.DiscordInfo = discordInfo
	producers.DownloadProducer(messageForRabbit)
}
