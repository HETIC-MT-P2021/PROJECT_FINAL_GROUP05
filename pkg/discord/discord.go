package discord

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func InitCarlosBot() error {
	session, err := CarlosBot()
	if err != nil {
		return err
	}

	fmt.Println("Carlos has now started properly. Press Ctrl+C to shutdown.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
	session.Close()
	return nil
}

func CarlosBot() (session *discordgo.Session, err error) {
	fmt.Println(os.Getenv("REACT_APP_DISCORD_TOKEN"))
	// Create a new Discord session using the provided bot token.
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

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	Message, status := DiscordCommandHandler(m.Author.ID, m.ChannelID, m.Content)
	if status {
		s.ChannelMessageSend(m.ChannelID, Message)
	}
}
