package router

import (
	"net/http"

	v1 "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/cmd/v1"
	_ "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/docs/api"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() (error, *gin.Engine) {
	gin.ForceConsoleColor()
	router := gin.Default()

	mysqlConnector, err := mysql.Connect()
	if err != nil {
		return err, nil
	}
	database.DB = mysqlConnector
	database.CommandRepo = mysql.NewCommandRepository(mysqlConnector)
	database.MediaRepo = mysql.NewMediaRepository(mysqlConnector)
	database.ServerRepo = mysql.NewServerRepository(mysqlConnector, database.CommandRepo, database.MediaRepo)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1.ApplyRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
	})

	return nil, router
}