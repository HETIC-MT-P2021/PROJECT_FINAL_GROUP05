package main

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/rabbit"
)

func main() {
	err := rabbit.ConnectToRabbitMQ()
	if err != nil {
		log.Println(err)
		return
	}

	_ = rabbit.NewRabbitRepository(rabbit.RabbitChannel, rabbit.RabbitQueue)

	log.Println("⚡️ app is running!")
	discord.CarlosBot()
}
