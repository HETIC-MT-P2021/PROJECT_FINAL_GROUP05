package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/bwmarrin/discordgo"
)

func UpdateMessage(session models.DiscordSession, fileName string) error {
	err := session.Session.ChannelMessageDelete(session.ChannelID, session.MessageID)
	if err != nil {
		return err
	}
	path := os.Getenv("DOCKER_CONTAINER_PATH")
	file, err := os.Open(fmt.Sprintf("%sedit-%s", path, fileName))
	if err != nil {
		return err
	}
  defer file.Close()

	msg := &discordgo.MessageSend{
		File: &discordgo.File{
			Name: fileName,
			Reader: file,
		},
	}

	_, err = session.Session.ChannelMessageSendComplex(session.ChannelID, msg)
	log.Println("step final...")
	return err
}