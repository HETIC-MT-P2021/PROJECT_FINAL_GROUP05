package fixtures

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
)

// NewCommands Creates 2 commands
func NewCommands(db *sql.DB) error {
	now := utils.NewDateNow()
	audioCommand := models.Command{
		Title: "Extrait audio",
		Command: "-audio -s X -d Y -i 'input'",
		IsActive: false,
		ServerID: "1",
		CreatedAt: now,
		UpdatedAt: now,
	}

	videoCommand := models.Command{
		Title: "Extrait vidéo",
		Command: "-video -s X -d Y -i 'input'",
		IsActive: false,
		ServerID: "2",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := database.CommandRepo.CreateCommand(audioCommand)
	err = database.CommandRepo.CreateCommand(videoCommand)

	return err
}