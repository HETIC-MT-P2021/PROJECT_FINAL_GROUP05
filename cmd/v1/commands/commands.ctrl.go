package commands

import (
	"net/http"

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
	c.JSON(http.StatusCreated, gin.H{
		"route": "GET /commands",
	})
}