package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	v1 "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/cmd/v1"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/rabbit"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	mysqlConnector, err := mysql.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	router.Use(pkg.ApiMiddleware(mysqlConnector))

	if err := database.Connect(); err != nil {
		log.Panic(err)
	}

	v1.ApplyRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	rabbit.InitProducer()
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
