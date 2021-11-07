package server

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
	"github.com/gin-gonic/gin"
)

const StatusInternalError = http.StatusInternalServerError
const StatusNotFoundError = http.StatusNotFound

// GetServers returns array of server from database
// @Summary Get all servers from database
// @Description Get an array of server struct
// @Tags servers
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Server
// @Failure 404 {object} utils.HTTPError "Failed to get servers"
// @Router /servers [GET]
func GetServers(c *gin.Context) {
	commandRepo := mysql.NewCommandRepository(database.DB)
	mediaRepo := mysql.NewMediaRepository(database.DB)
	serverRepo := mysql.NewServerRepository(database.DB, commandRepo, mediaRepo)
	servers, err := serverRepo.GetServers()
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to get servers")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusFound, servers)
}

// GetServer returns server with medias and commands from database
// @Summary Get server from database
// @Description Get server struct with command and media struct
// @Tags servers
// @Accept  json
// @Produce  json
// @Param id path int true "Server ID"
// @Success 200 {object} models.Server
// @Failure 404 {object} utils.HTTPError "Failed to get server"
// @Router /servers/{id} [GET]
func GetServer(c *gin.Context) {
	serverID := c.Param("id")
	server, err := database.ServerRepo.GetServer(serverID)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to get server")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusFound, utils.GetServerResponse(server))
}

// CreateServer insert new server in database
// @Summary Insert a server with default commands in database
// @Tags servers
// @Accept  json
// @Produce  json
// @Success 201 {object} utils.HTTPStatus	"Created"
// @Failure 500 {object} utils.HTTPError "Failed to create new server"
// @Router /servers [POST]
func CreateServer(c *gin.Context) {
	var req models.Server
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new server")
		log.Println(err)
		return
	}

	now := utils.NewDateNow()
	req.CreatedAt = now
	req.UpdatedAt = now

	err := database.ServerRepo.CreateServer(req)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new server")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}

// GetServerMedias returns array of medias from database
// @Summary Get all medias of a server from database
// @Tags servers
// @Accept  json
// @Produce  json
// @Param id path int true "Server ID"
// @Success 200 {object} []models.Media
// @Failure 404 {object} utils.HTTPError "Failed to get medias"
// @Router /servers/{id}/medias [GET]
func GetServerMedias(c *gin.Context) {
	serverID := c.Param("id")
	medias, err := database.ServerRepo.GetServerMedias(serverID)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to get medias")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusFound, utils.GetServerMediasResponse(medias))
}

// GetServerCommands returns array of commands from database
// @Summary Get all commands of a server from database
// @Tags servers
// @Accept  json
// @Produce  json
// @Param id path int true "Server ID"
// @Success 200 {object} []models.Command
// @Failure 404 {object} utils.HTTPError "Failed to get commands"
// @Router /servers/{id}/commands [GET]
func GetServerCommands(c *gin.Context) {
	serverID := c.Param("id")
	medias, err := database.ServerRepo.GetServerCommands(serverID)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to get commands")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusFound, utils.GetServerCommandsResponse(medias))
}