package main

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/router"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
)

func main() {
	err, router := router.Init()
	if err != nil {
		log.Println(err)
		return
	}
	
	utils.InitServer(router)

	message := models.DownloadMessage{
		Type: "audio",
		MediaLink: "youtube.com",
		Options: 	models.Options{
			StartInSeconds: "20",
			DurationInSeconds: "35",
		},
	}
	producers.DownloadProducer(message)

	err = discord.InitCarlosBot()
	if err != nil {
		log.Println(err)
		return
	}
}
