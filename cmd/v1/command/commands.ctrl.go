package command

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
	"github.com/gin-gonic/gin"
)

const StatusInternalError = http.StatusInternalServerError
const StatusNotFoundError = http.StatusNotFound

// GetCommands returns array of commands from database
// @Summary Get all commands from database
// @Description Get an array of command struct
// @Tags commands
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Command
// @Failure 404 {object} utils.HTTPError "Failed to get commands"
// @Router /commands [GET]
func GetCommands(c *gin.Context) {
	commands, err := database.CommandRepo.GetCommands()
	if err != nil {
		utils.DisplayErrorMessage(c, StatusNotFoundError, "Failed to get commands")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusFound, commands)
}

// UpdateCommand from database
// @Summary update command from database
// @Tags commands
// @Accept  json
// @Produce  json
// @Param id path int true "Command ID"
// @Success 200 {object} models.Command
// @Failure 500 {object} utils.HTTPError "Failed to update command"
// @Router /commands/{id} [PUT]
func UpdateCommand(c *gin.Context) {
	commandID := c.Param("id")

	var req models.Command
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update command")
		log.Println(err)
		return
	}
	
	req.UpdatedAt = utils.NewDateNow()
	err := database.CommandRepo.UpdateCommand(commandID, req)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update command")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusOK, req)
}

// UpdateIsActiveCommand from database
// @Summary update is_active field on a command from database
// @Tags commands
// @Accept  json
// @Produce  json
// @Param id path int true "Command ID"
// @Success 200 {object} models.Command
// @Failure 500 {object} utils.HTTPError "Failed to update command"
// @Router /commands/{id} [PATCH]
func UpdateIsActiveCommand(c *gin.Context) {
	commandID := c.Param("id")

	var req models.Command
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update command")
		log.Println(err)
		return
	}
	
	req.UpdatedAt = utils.NewDateNow()
	err := database.CommandRepo.UpdateIsActiveFieldCommand(commandID, req.IsActive)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update command")
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, req)
}
