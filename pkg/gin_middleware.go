package pkg

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// ApiMiddleware will add the db connection to the context
func ApiMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Set("databaseConn", db)
			c.Next()
	}
}