package fixtures

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

// NewMedias Creates 2 medias
func NewMedias(db *sql.DB) error {
	mediaOne := models.Media{
		ID: 1,
		DiscordUrl: "https://c.tenor.com/I6kN-6X7nhAAAAAj/loading-buffering.gif",
		IsArchived: false,
		UserID: "1",
		ServerID: "694885041905598545",
	}

	mediaTwo := models.Media{
		ID: 1,
		DiscordUrl: "https://c.tenor.com/I6kN-6X7nhAAAAAj/loading-buffering.gif",
		IsArchived: false,
		UserID: "2",
		ServerID: "694885041905598545",
	}

	repo := mysql.NewMediaRepository(db)
	err := repo.CreateMedia(mediaOne)
	err = repo.CreateMedia(mediaTwo)

	return err
}