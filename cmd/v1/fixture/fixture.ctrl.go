package fixture

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql/fixtures"
	"github.com/gin-gonic/gin"
)

// CreateFixtures Insert servers, commands, medias in database
// @Summary Create servers, commands in db
// @Tags fixtures
// @Accept  json
// @Produce  json
// @Success 201 {object} pkg.HTTPStatus "Created"
// @Failure 500 {object} pkg.HTTPStatus "Not Created"
// @Router /fixtures [POST]
func CreateFixtures(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
		return
	}

	if err := fixtures.NewServers(dbConn); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "Not created",
		})
		return
	}

	if err := fixtures.NewMedias(dbConn); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "Not created",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}


// CreateServersFixtures Insert servers database
// @Summary Create servers in db
// @Tags fixtures
// @Accept  json
// @Produce  json
// @Success 201 {object} pkg.HTTPStatus "Created"
// @Failure 500 {object} pkg.HTTPStatus "Not Created"
// @Router /fixtures/servers [POST]
func CreateServersFixtures(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
	}

	if err := fixtures.NewServers(dbConn); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "Not created",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}

// CreateCommandsFixtures Insert commands database
// @Summary Create commands in db
// @Tags fixtures
// @Accept  json
// @Produce  json
// @Success 201 {object} pkg.HTTPStatus "Created"
// @Failure 500 {object} pkg.HTTPStatus "Not Created"
// @Router /fixtures/commands [POST]
func CreateCommandsFixtures(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
	}

	if err := fixtures.NewCommands(dbConn); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "Not created",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}

// CreateMediasFixtures Insert medias database
// @Summary Create medias in db
// @Tags fixtures
// @Accept  json
// @Produce  json
// @Success 201 {object} pkg.HTTPStatus "Created"
// @Failure 500 {object} pkg.HTTPStatus "Not Created"
// @Router /fixtures/medias [POST]
func CreateMediasFixtures(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
		return
	}

	if err := fixtures.NewMedias(dbConn); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "Not created",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}
