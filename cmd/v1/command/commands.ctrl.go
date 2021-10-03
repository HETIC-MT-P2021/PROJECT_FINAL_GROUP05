package command

import (
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/commands"
	"github.com/gin-gonic/gin"
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
	repo := commands.NewCommandsRepository(database.DbConn)
	commands, err := repo.GetCommands()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(http.StatusFound, commands)
}