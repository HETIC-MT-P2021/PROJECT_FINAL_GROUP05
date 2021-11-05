package fixture

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes All routes for fixtures
func ApplyRoutes(r *gin.RouterGroup) {
	r.POST("/fixtures", CreateFixtures)
	r.POST("/fixtures/servers", CreateServersFixtures)
	r.POST("/fixtures/commands", CreateCommandsFixtures)
	r.POST("/fixtures/medias", CreateMediasFixtures)
}
