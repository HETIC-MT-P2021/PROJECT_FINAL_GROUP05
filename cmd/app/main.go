package main

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord/carlos"
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

	carlosBot := carlos.NewDiscordRepository()
	err = carlosBot.InitBotWithHandler()
	if err != nil {
		log.Println(err)
		return
	}
}
