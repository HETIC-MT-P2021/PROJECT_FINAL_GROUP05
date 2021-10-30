package server

import "github.com/gin-gonic/gin"

// ApplyRoutes All routes for servers
func ApplyRoutes(r *gin.RouterGroup) {
	r.GET("/servers", GetServers)
	//r.PUT("/commands/:id", UpdateCommand)
	//r.PATCH("/commands/:id", UpdateIsActiveCommand)
}