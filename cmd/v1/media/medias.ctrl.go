package media

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/gin-gonic/gin"
)

// UpdateMedia from database
// @Summary update media from database
// @Tags medias
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"GET /medias"
// @Router /medias/{id} [PUT]
func UpdateMedia(c *gin.Context) {
	mediaID := c.Param("id")
	mediaIDConverted, err := strconv.Atoi(mediaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
	}

	var req models.Media
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	repo := mysql.NewMediaRepository(dbConn)
	err = repo.UpdateMedia(mediaIDConverted, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(http.StatusOK, req)
}

// UpdateIsArchivedMedia from database
// @Summary update is_archived field on a media from database
// @Tags medias
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"GET /medias/id"
// @Router /medias/{id} [PATCH]
func UpdateIsArchivedMedia(c *gin.Context) {
	mediaID := c.Param("id")

	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve Database connection",
		})
	}

	var req models.Media
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	repo := mysql.NewMediaRepository(dbConn)
	err := repo.UpdateIsArchivedFieldMedia(mediaID, req.IsArchived)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"is_archived": req.IsArchived,
	})
}
