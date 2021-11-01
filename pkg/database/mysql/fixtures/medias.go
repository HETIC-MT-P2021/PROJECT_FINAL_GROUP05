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
		ServerID: "b48a2954-9a51-4f31-84a8-3bc16e6a241f",
	}

	mediaTwo := models.Media{
		ID: 1,
		DiscordUrl: "https://c.tenor.com/I6kN-6X7nhAAAAAj/loading-buffering.gif",
		IsArchived: false,
		UserID: "2",
		ServerID: "c59b3065-0b62-5g42-95b9-4cd27f7b352g",
	}

	repo := mysql.NewMediaRepository(db)
	err := repo.CreateMedia(mediaOne)
	err = repo.CreateMedia(mediaTwo)

	return err
}