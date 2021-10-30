package fixture

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes All routes for fixtures
func ApplyRoutes(r *gin.RouterGroup) {
	r.POST("/fixtures", CreateFixtures)
}
