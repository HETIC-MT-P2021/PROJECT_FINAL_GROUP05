package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetServers returns array of server from database
// @Summary Get all servers from database
// @Description Get an array of server struct
// @Tags servers
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"GET /servers"
// @Router /servers [GET]
func GetServers(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
		return
	}
	
	commandRepo := mysql.NewCommandRepository(dbConn)
	serverRepo := mysql.NewServerRepository(dbConn, commandRepo)
	servers, err := serverRepo.GetServers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
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
// @Success 200 {string} string	"GET /servers"
// @Router /servers/{id} [GET]
func GetServer(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
		return
	}
	serverID := c.Param("id")
	commandRepo := mysql.NewCommandRepository(dbConn)
	serverRepo := mysql.NewServerRepository(dbConn, commandRepo)
	server, err := serverRepo.GetServer(serverID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	
	c.JSON(http.StatusFound, utils.GetServerResponse(server))
}

// CreateServer insert new server in database
// @Summary Insert a server with default commands in database
// @Tags servers
// @Accept  json
// @Produce  json
// @Success 201 {string} string	"POST /servers"
// @Router /servers [POST]
func CreateServer(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
		return
	}

	var req models.Server
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	req.ID = uuid.NewV4().String()
	req.CreatedAt = time.Unix(time.Now().Unix(), 0).Format("2000-01-01")

	commandRepo := mysql.NewCommandRepository(dbConn)
	serverRepo := mysql.NewServerRepository(dbConn, commandRepo)
	err := serverRepo.CreateServer(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	
	c.JSON(http.StatusCreated, req)
}

// GetServerMedias returns array of medias from database
// @Summary Get all medias of a server from database
// @Tags servers
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"GET /servers"
// @Router /servers/{id}/medias [GET]
func GetServerMedias(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
		return
	}
	
	commandRepo := mysql.NewCommandRepository(dbConn)
	serverRepo := mysql.NewServerRepository(dbConn, commandRepo)

	serverID := c.Param("id")
	medias, err := serverRepo.GetServerMedias(serverID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	
	c.JSON(http.StatusFound, utils.GetServerMediasResponse(medias))
}