package database

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

// serverRepository functions for server repository
type serverRepository interface {
	CreateServer(models.Server) error
	GetServerss() ([]models.Server, error)
}