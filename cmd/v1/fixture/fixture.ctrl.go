package fixture

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql/fixtures"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
	"github.com/gin-gonic/gin"
)

const StatusInternalError = http.StatusInternalServerError

// CreateFixtures Insert servers, commands, medias in database
// @Summary Create servers, commands in db
// @Tags fixtures
// @Accept  json
// @Produce  json
// @Success 201 {object} utils.HTTPStatus "Created"
// @Failure 500 {object} utils.HTTPStatus "Not Created"
// @Router /fixtures [POST]
func CreateFixtures(c *gin.Context) {
	if err := fixtures.NewServers(database.DB); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new servers")
		log.Println(err)
		return
	}

	if err := fixtures.NewMedias(database.DB); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new medias")
		log.Println(err)
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
// @Success 201 {object} utils.HTTPStatus "Created"
// @Failure 500 {object} utils.HTTPStatus "Not Created"
// @Router /fixtures/servers [POST]
func CreateServersFixtures(c *gin.Context) {
	if err := fixtures.NewServers(database.DB); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new servers")
		log.Println(err)
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
// @Success 201 {object} utils.HTTPStatus "Created"
// @Failure 500 {object} utils.HTTPStatus "Not Created"
// @Router /fixtures/commands [POST]
func CreateCommandsFixtures(c *gin.Context) {
	if err := fixtures.NewCommands(database.DB); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new commands")
		log.Println(err)
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
// @Success 201 {object} utils.HTTPStatus "Created"
// @Failure 500 {object} utils.HTTPStatus "Not Created"
// @Router /fixtures/medias [POST]
func CreateMediasFixtures(c *gin.Context) {
	if err := fixtures.NewMedias(database.DB); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new medias")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}
