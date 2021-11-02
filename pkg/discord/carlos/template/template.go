package template

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type MessageSendWrapper struct {
	MessageSend *discordgo.MessageSend
}

func NewMessageSendWrapper() *MessageSendWrapper {
	return &MessageSendWrapper{
		MessageSend: &discordgo.MessageSend{},
	}
}

const LOADER_URL = "https://c.tenor.com/I6kN-6X7nhAAAAAj/loading-buffering.gif"

func (message *MessageSendWrapper) SetCarlosIsProcessingMessage() {
	message.SetEmptyEmbed()
	message.SetThumbnail(LOADER_URL)
	message.SetTitle("Carlos is processing")
	message.SetDescription(
		"Media :  \n" +
		"State : In progress")
}

func (message *MessageSendWrapper) SetForbiddenAccessMessage() {
	message.SetEmptyEmbed()
	message.SetTitle("You're not registered to our bot")
}

func (message *MessageSendWrapper) SetEmptyEmbed() {
	message.MessageSend.Embed = &discordgo.MessageEmbed{}
}

func (message *MessageSendWrapper) SetThumbnail(url string) {
	message.MessageSend.Embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
}

func (message *MessageSendWrapper) SetTitle(title string) {
	message.MessageSend.Embed.Title = title
}

func (message *MessageSendWrapper) SetDescription(description string) {
	message.MessageSend.Embed.Description = description
}

func (message *MessageSendWrapper) SetFile(fileName string, file *os.File) {
	message.MessageSend.File = &discordgo.File{
		Name: fileName,
		Reader: file,
	}
}