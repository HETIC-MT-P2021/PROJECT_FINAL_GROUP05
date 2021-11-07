package fixtures

import (
	"database/sql"
	"log"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
)

// NewServers Creates 2 servers
func NewServers(db *sql.DB) error {
	now := utils.NewDateNow()
	log.Println("now")
	log.Println(now)
	serverOne := models.Server{
		ServerID: "694885041905598545",
		ServerName: "Wyllis",
		CreatedAt: now,
		UpdatedAt: now,
	}

	serverTwo := models.Server{
		ServerID: "fake",
		ServerName: "Fake",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := database.ServerRepo.CreateServer(serverOne)
	err = database.ServerRepo.CreateServer(serverTwo)

	return err
}