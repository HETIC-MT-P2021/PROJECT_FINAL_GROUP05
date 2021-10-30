package database

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

// commandRepository functions for command repository
type commandRepository interface {
	CreateCommand(command models.Command) error
	GetCommands() ([]models.Command, error)
	UpdateCommand(commandID string, command models.Command) error
	UpdateIsActiveFieldCommand(commandID string, isActive bool) error
}