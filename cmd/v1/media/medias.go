package media

import "github.com/gin-gonic/gin"

// ApplyRoutes All routes for servers
func ApplyRoutes(r *gin.RouterGroup) {
	r.PUT("/medias/:id", UpdateMedia)
	r.PATCH("/medias/:id", UpdateIsArchivedMedia)
}