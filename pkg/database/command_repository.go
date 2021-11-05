package database

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

// CommandRepository functions for command repository
type CommandRepository interface {
	CreateCommand(command models.Command) error
	CreateDefaultCommandsInServer(serverID string) error
	GetCommands() ([]models.Command, error)
	UpdateCommand(commandID string, command models.Command) error
	UpdateIsActiveFieldCommand(commandID string, isActive bool) error
}