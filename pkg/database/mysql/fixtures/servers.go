package fixtures

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

// NewServers Creates 2 servers
func NewServers(db *sql.DB) error {
	serverOne := models.Server{
		ID: "1",
		ServerName: "Wyllis",
		CreatedAt: "2021-01-01",
	}

	serverTwo := models.Server{
		ID: "2",
		ServerName: "Jibé",
		CreatedAt: "2021-02-02",
	}

	repo := mysql.NewServerRepository(db)
	err := repo.CreateServer(serverOne)
	err = repo.CreateServer(serverTwo)

	return err
}