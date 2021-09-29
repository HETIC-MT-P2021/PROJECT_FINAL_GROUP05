package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	log.Println("⚡️ Queue is running!")

	session, err := discord.CarlosBot()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Carlos has now started properly. Press Ctrl+C to shutdown.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
	session.Close()
}
