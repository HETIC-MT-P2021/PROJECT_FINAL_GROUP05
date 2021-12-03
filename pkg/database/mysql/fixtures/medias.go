package fixtures

import (
	"database/sql"
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
)

// NewMedias Creates 2 medias
func NewMedias(db *sql.DB) error {
	now := utils.NewDateNow()
	log.Println("now")
	log.Println(now)
	mediaOne := models.Media{
		DiscordUrl: "https://c.tenor.com/I6kN-6X7nhAAAAAj/loading-buffering.gif",
		IsArchived: false,
		UserID: "1",
		ServerID: "694885041905598545",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mediaTwo := models.Media{
		DiscordUrl: "https://c.tenor.com/I6kN-6X7nhAAAAAj/loading-buffering.gif",
		IsArchived: false,
		UserID: "2",
		ServerID: "694885041905598545",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := database.MediaRepo.CreateMedia(mediaOne)
	err = database.MediaRepo.CreateMedia(mediaTwo)

	return err
}