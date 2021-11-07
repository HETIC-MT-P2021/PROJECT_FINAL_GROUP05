package media

import (
	"log"
	"net/http"
	"strconv"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
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
// @Param id path int true "Media ID"
// @Success 200 {object} models.Media
// @Failure 500 {object} utils.HTTPError "Failed to update media"
// @Router /medias/{id} [PUT]
func UpdateMedia(c *gin.Context) {
	mediaID := c.Param("id")
	mediaIDConverted, err := strconv.Atoi(mediaID)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Media is not valid")
		log.Println(err)
		return
	}

	var req models.Media
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update media")
		log.Println(err)
		return
	}
	
	req.UpdatedAt = utils.NewDateNow()
	err = database.MediaRepo.UpdateMedia(mediaIDConverted, req)
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
// @Param id path int true "Media ID"
// @Success 200 {object} models.Media
// @Failure 500 {object} utils.HTTPError "Failed to update media"
// @Router /medias/{id} [PATCH]
func UpdateIsArchivedMedia(c *gin.Context) {
	mediaID := c.Param("id")
	mediaIDConverted, err := strconv.Atoi(mediaID)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Media is not valid")
		log.Println(err)
		return
	}

	var req models.Media
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update media")
		log.Println(err)
		return
	}

	req.UpdatedAt = utils.NewDateNow()
	err = database.MediaRepo.UpdateIsArchivedFieldMedia(mediaIDConverted, req.IsArchived)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to update media")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusOK, req)
}

// CreateMedia from database
// @Summary create media from database
// @Tags medias
// @Accept  json
// @Produce  json
// @Success 201 {object} utils.HTTPStatus	"Created"
// @Failure 500 {object} utils.HTTPError "Failed to create new media"
// @Router /medias [POST]
func CreateMedia(c *gin.Context) {
	var req models.Media
	if err := c.BindJSON(&req); err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new media")
		log.Println(err)
		return
	}

	now := utils.NewDateNow()
	req.CreatedAt = now
	req.UpdatedAt = now

	err := database.MediaRepo.CreateMedia(req)
	if err != nil {
		utils.DisplayErrorMessage(c, StatusInternalError, "Failed to create new media")
		log.Println(err)
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}
