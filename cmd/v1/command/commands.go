package command

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes All routes for articles
func ApplyRoutes(r *gin.RouterGroup) {
	r.GET("/commands", GetCommands)
	r.PUT("/commands/:id", UpdateCommand)
}