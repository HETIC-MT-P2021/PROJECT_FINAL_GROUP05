package server

import "github.com/gin-gonic/gin"

// ApplyRoutes All routes for servers
func ApplyRoutes(r *gin.RouterGroup) {
	r.GET("/servers", GetServers)
	r.GET("/servers/:id", GetServer)
	r.POST("/servers", CreateServer)
}