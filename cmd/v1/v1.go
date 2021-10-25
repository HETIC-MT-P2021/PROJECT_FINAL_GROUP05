package v1

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/cmd/v1/command"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/cmd/v1/fixture"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	app := r.Group("/api/v1/")
	{
		command.ApplyRoutes(app)
		fixture.ApplyRoutes(app)
	}
}
