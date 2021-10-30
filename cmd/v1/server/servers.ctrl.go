package server

import (
	"database/sql"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/gin-gonic/gin"
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
	}
	
	repo := mysql.NewServerRepository(dbConn)
	servers, err := repo.GetServers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
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
	}
	serverID := c.Param("id")
	repo := mysql.NewServerRepository(dbConn)
	commands, err := repo.GetServer(serverID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(http.StatusFound, commands)
}