package command

import (
	"database/sql"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
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
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
	}
	
	repo := mysql.NewCommandRepository(dbConn)
	commands, err := repo.GetCommands()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(http.StatusFound, commands)
}

// UpdateCommand from database
// @Summary update command from database
// @Tags commands
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"GET /commands"
// @Router /commands/{id} [PUT]
func UpdateCommand(c *gin.Context) {
	commandID := c.Param("id")

	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
	}

	var req models.Command
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	repo := mysql.NewCommandRepository(dbConn)
	err := repo.UpdateCommand(commandID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(http.StatusOK, req)
}

// UpdateIsActiveCommand from database
// @Summary update is_active field on a command from database
// @Tags commands
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"GET /commands"
// @Router /commands/{id} [PATCH]
func UpdateIsActiveCommand(c *gin.Context) {
	commandID := c.Param("id")

	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
	}

	var req models.Command
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	repo := mysql.NewCommandRepository(dbConn)
	err := repo.UpdateIsActiveFieldCommand(commandID, req.IsActive)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"is_active": req.IsActive,
	})
}
