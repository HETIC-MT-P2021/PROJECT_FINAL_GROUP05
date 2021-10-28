package discord

import (
	"flag"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Token Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func CarlosBot() (session *discordgo.Session, err error) {
	dg, err := discordgo.New("Bot " + os.Getenv("REACT_APP_DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil, err
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(MessageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return nil, err
	}

	//Returns the created session
	return dg, nil
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	Message, status := CommandHandler(m.Author.ID, m.ChannelID, m.Content)
	if status {
		s.ChannelMessageSend(m.ChannelID, Message)
	}
}
