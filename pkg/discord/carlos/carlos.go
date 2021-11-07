package carlos

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord/carlos/template"
	"github.com/bwmarrin/discordgo"
)

type CarlosBot struct {
	Session 		*discordgo.Session
	ServerID 		string
	ChannelID 	string
	MessageID 	string
}

func NewDiscordRepository() *CarlosBot {
	return &CarlosBot{}
}

func (bot *CarlosBot) InitBotWithHandler() error {
	err := bot.makeBotConnexion()
	if err != nil {
		return err
	}

	bot.Session.AddHandler(MessageCreate)

	fmt.Println("Carlos has now started properly with handler. Press Ctrl+C to shutdown.")

	bot.StopBot()
	return nil
}

func (bot *CarlosBot) InitBotWithoutHandler() error {
	err := bot.makeBotConnexion()
	if err != nil {
		return err
	}

	return nil
}

func (bot *CarlosBot) makeBotConnexion() error {
	session, err := discordgo.New("Bot " + os.Getenv("REACT_APP_DISCORD_TOKEN"))
	if err != nil {
		return err
	}

	bot.Session = session

	// In this example, we only care about receiving message events.
	session.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = session.Open()
	if err != nil {
		return err
	}

	return nil
}

func (bot *CarlosBot) StopBot() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	bot.Session.Close()
}

func (bot *CarlosBot) UpdateCarlosIsProcessingMessage(fileName string) (*discordgo.Message, error) {
	err := bot.Session.ChannelMessageDelete(bot.ChannelID, bot.MessageID)
	if err != nil {
		return &discordgo.Message{}, err
	}
	path := fmt.Sprintf("%sdownloads/", os.Getenv("BASE_DIR_IN_CONTAINER"))
	file, err := os.Open(fmt.Sprintf("%sedit-%s", path, fileName))
	if err != nil {
		return &discordgo.Message{}, err
	}
  defer file.Close()

	messageForDiscord := template.NewMessageSendWrapper()
	messageForDiscord.SetFile(fileName, file)

	return bot.Session.ChannelMessageSendComplex(bot.ChannelID, messageForDiscord.MessageSend)
}