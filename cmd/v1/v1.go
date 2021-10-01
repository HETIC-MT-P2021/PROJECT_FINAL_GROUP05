package v1

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/cmd/v1/commands"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	app := r.Group("/api/v1/")
	{
		commands.ApplyRoutes(app)
	}
}