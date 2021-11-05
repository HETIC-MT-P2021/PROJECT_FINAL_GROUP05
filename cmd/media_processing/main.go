package main

import (
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/consumers"
)

func main() {
	err := commands.InitCommandBus()
	if err != nil {
		log.Println(err)
		return
	}

	mysqlConnector, err := mysql.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	database.DB = mysqlConnector
	database.CommandRepo = mysql.NewCommandRepository(mysqlConnector)
	database.MediaRepo = mysql.NewMediaRepository(mysqlConnector)
	database.ServerRepo = mysql.NewServerRepository(mysqlConnector, database.CommandRepo, database.MediaRepo)

	consumers.MediaProcessingConsumer()
}