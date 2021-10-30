package command

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes All routes for commands
func ApplyRoutes(r *gin.RouterGroup) {
	r.GET("/commands", GetCommands)
	r.PUT("/commands/:id", UpdateCommand)
	r.PATCH("/commands/:id", UpdateIsActiveCommand)
}