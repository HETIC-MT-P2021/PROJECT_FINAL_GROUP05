package router

import (
	"net/http"

	v1 "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/cmd/v1"
	"github.com/gin-gonic/gin"
)

func Init() (error, *gin.Engine) {
	gin.ForceConsoleColor()
	router := gin.Default()
	//mysqlConnector, err := mysql.Connect()
	//if err != nil {
	//	return err, nil
	//}
	//router.Use(utils.ApiMiddleware(mysqlConnector))

	v1.ApplyRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
	})

	return nil, router
}
