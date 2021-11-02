package database

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

// ServerRepository functions for server repository
type ServerRepository interface {
	CreateServer(models.Server) error
	GetServers() ([]models.Server, error)
	GetServer(serverID string) (*models.ServerCommandsAndMedias, error)
	GetServerMedias(serverID string) ([]models.Media, error)
	GetServerCommands(serverID string) ([]models.Command, error)
}