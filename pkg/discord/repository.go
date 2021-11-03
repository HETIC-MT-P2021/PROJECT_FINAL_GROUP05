package discord

import "github.com/bwmarrin/discordgo"

type DiscordRepository interface {
	InitBotWithHandler() error
	InitBotWithoutHandler() error
	StopBot()
	UpdateCarlosIsProcessingMessage(fileName string) (*discordgo.Message, error)
}