package main

import (
	"flag"
	"fmt"
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

func bot() {

	// Create a new Discord session using the provided bot token. Currently using Carlos Bot's token.
	dg, err := discordgo.New("Bot " + "ODkyMzk3MTI2MzAxMTU1MzMw.YVMTlg.sMgLiQEQH_ncQMymxCIkpOusLfU")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" || m.Content == "carlos" {
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/694885042400264266/892463464591790190/carlos.jpg")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" || m.Content == "contrecarlos" {
		s.ChannelMessageSend(m.ChannelID, "https://media.discordapp.net/attachments/694885042400264266/892464389855256586/solrac.png")
	}
}
