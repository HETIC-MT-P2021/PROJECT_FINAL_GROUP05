package discord

type DiscordRepository interface {
	InitBotWithHandler() error
	InitBotWithoutHandler() error
	StopBot()
	UpdateCarlosIsProcessingMessage() error
}