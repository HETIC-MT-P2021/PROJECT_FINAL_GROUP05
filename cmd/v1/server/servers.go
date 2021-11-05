package server

import "github.com/gin-gonic/gin"

// ApplyRoutes All routes for servers
func ApplyRoutes(r *gin.RouterGroup) {
	r.GET("/servers", GetServers)
	r.GET("/servers/:id", GetServer)
	r.POST("/servers", CreateServer)
	r.GET("/servers/:id/medias", GetServerMedias)
	r.GET("/servers/:id/commands", GetServerCommands)
}