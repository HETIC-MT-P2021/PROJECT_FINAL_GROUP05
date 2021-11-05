package carlos

import (
	"log"
	"regexp"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord/carlos/template"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/parsings"
	"github.com/bwmarrin/discordgo"
)

const (
	regexFindType = `^(audio|video|image)`
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	typeRegex := regexp.MustCompile(regexFindType)
	isMediaType := typeRegex.Match([]byte(m.Content))
	if !isMediaType {
		return
	}

	log.Println(m.GuildID)

	server, err := database.ServerRepo.GetServer(m.GuildID)
	if err != nil {
		log.Println(err)
		return
	}

	messageForDiscord := template.NewMessageSendWrapper()
	if len(server.ServerName) <= 0 {
		messageForDiscord.SetForbiddenAccessMessage()
		s.ChannelMessageSendComplex(m.ChannelID, messageForDiscord.MessageSend)
		return
	}
	
	messageForDiscord.SetCarlosIsProcessingMessage()

	discordMessage, err := s.ChannelMessageSendComplex(m.ChannelID, messageForDiscord.MessageSend)
	if err != nil {
		log.Println(err)
		return
	}
	discordInfo := models.JobCheckerQueueMessage{
		ServerID: m.GuildID,
		ChannelID: discordMessage.ChannelID,
		MessageID: discordMessage.ID,
	}

	messageForRabbit := parsings.ParseCommand(m.Content)
	messageForRabbit.DiscordInfo = discordInfo
	producers.DownloadProducer(messageForRabbit)
}
