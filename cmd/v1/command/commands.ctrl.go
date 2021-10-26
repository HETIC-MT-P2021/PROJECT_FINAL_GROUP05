package command

import (
	"database/sql"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/media_download/video/ytb"
	"github.com/gin-gonic/gin"
	"github.com/kkdai/youtube/v2"
)

// GetCommands returns array of commands from database
// @Summary Get all commands from database
// @Description Get an array of command struct
// @Tags commands
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"GET /commands"
// @Router /commands [get]
func GetCommands(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
	}
	
	repo := mysql.NewCommandsRepository(dbConn)
	commands, err := repo.GetCommands()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(http.StatusFound, commands)
}

func Download(c *gin.Context) {
	downloadService := ytb.NewDownloadRepository(youtube.Client{})
	err := downloadService.Download("9QI8PCylni4");
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "nice dl",
	})
}