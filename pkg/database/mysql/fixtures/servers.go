package fixtures

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

// NewServers Creates 2 servers
func NewServers(db *sql.DB) error {
	serverOne := models.Server{
		ID: "b48a2954-9a51-4f31-84a8-3bc16e6a241f",
		ServerName: "Wyllis",
		CreatedAt: "2021-01-01",
	}

	serverTwo := models.Server{
		ID: "c59b3065-0b62-5g42-95b9-4cd27f7b352g",
		ServerName: "Jibé",
		CreatedAt: "2021-02-02",
	}

	commandRepo := mysql.NewCommandRepository(db)
	serverRepo := mysql.NewServerRepository(db, commandRepo)
	err := serverRepo.CreateServer(serverOne)
	err = serverRepo.CreateServer(serverTwo)

	return err
}