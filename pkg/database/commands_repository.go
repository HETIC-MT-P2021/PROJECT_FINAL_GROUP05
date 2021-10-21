package database

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

type commandsRepository interface {
	CreateCommand(command models.Command) error
	GetCommands() ([]models.Command, error)
}