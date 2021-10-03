package fixture

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/fixtures"
	"github.com/gin-gonic/gin"
)

// CreateCommands Insert commands in database
// @Summary Create commands in db
// @Tags fixtures
// @Accept  json
// @Produce  json
// @Success 201 {object} pkg.HTTPStatus "Created"
// @Failure 500 {object} pkg.HTTPStatus "Not Created"
// @Router /fixtures/commands [post]
func CreateCommands(c *gin.Context) {
	if err := fixtures.NewCommands(); err != nil {
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
