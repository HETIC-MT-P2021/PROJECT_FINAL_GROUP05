package fixtures

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

// NewServers Creates 2 servers
func NewServers(db *sql.DB) error {
	serverOne := models.Server{
		ServerID: "694885041905598545",
		ServerName: "Wyllis",
		CreatedAt: "2021-01-01",
	}

	serverTwo := models.Server{
		ServerID: "fake",
		ServerName: "Fake",
		CreatedAt: "2021-02-02",
	}

	err := database.ServerRepo.CreateServer(serverOne)
	err = database.ServerRepo.CreateServer(serverTwo)

	return err
}