package media

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
	"github.com/gin-gonic/gin"
)

const StatusInternalError = http.StatusInternalServerError

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
		utils.DisplayErrorMessage(c, StatusInternalError, "Media is not valid")
		log.Println(err)
		return
	}

	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to retrieve Database connection")
		return
	}

	var req models.Media
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update media")
		log.Println(err)
		return
	}
	
	repo := mysql.NewMediaRepository(dbConn)
	err = repo.UpdateMedia(mediaIDConverted, req)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update media")
		log.Println(err)
		return
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
	mediaIDConverted, err := strconv.Atoi(mediaID)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Media is not valid")
		log.Println(err)
		return
	}

	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to retrieve Database connection")
		return
	}

	var req models.Media
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update media")
		log.Println(err)
		return
	}
	
	repo := mysql.NewMediaRepository(dbConn)
	err = repo.UpdateIsArchivedFieldMedia(mediaIDConverted, req.IsArchived)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update media")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"is_archived": req.IsArchived,
	})
}

// CreateMedia from database
// @Summary create media from database
// @Tags medias
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"POST /medias"
// @Router /medias [POST]
func CreateMedia(c *gin.Context) {
	dbConn, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to retrieve Database connection")
		return
	}

	var req models.Media
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new media")
		log.Println(err)
		return
	}

	mediaRepo := mysql.NewMediaRepository(dbConn)
	err := mediaRepo.CreateMedia(req)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new media")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusCreated, req)
}
