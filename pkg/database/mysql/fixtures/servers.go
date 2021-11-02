package fixtures

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

// NewServers Creates 2 servers
func NewServers(db *sql.DB) error {
	serverOne := models.Server{
		ServerID: "b48a2954-9a51-4f31-84a8-3bc16e6a241f",
		ServerName: "Wyllis",
		CreatedAt: "2021-01-01",
	}

	serverTwo := models.Server{
		ServerID: "x48a2954-9a51-4f31-84a8-3bc16e6a241f",
		ServerName: "Jibé",
		CreatedAt: "2021-02-02",
	}

	commandRepo := mysql.NewCommandRepository(db)
	mediaRepo := mysql.NewMediaRepository(db)
	serverRepo := mysql.NewServerRepository(db, commandRepo, mediaRepo)
	err := serverRepo.CreateServer(serverOne)
	err = serverRepo.CreateServer(serverTwo)

	return err
}