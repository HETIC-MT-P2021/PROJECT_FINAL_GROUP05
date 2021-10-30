package fixtures

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

// NewCommands Creates 2 commands
func NewCommands(db *sql.DB) error {
	audioCommand := models.Command{
		Title: "Extrait audio",
		Command: "-audio -s X -d Y -i 'input'",
		IsActive: false,
		ServerID: "1",
	}

	videoCommand := models.Command{
		Title: "Extrait vidéo",
		Command: "-video -s X -d Y -i 'input'",
		IsActive: false,
		ServerID: "2",
	}

	repo := mysql.NewCommandRepository(db)
	err := repo.CreateCommand(audioCommand)
	err = repo.CreateCommand(videoCommand)

	return err
}