package database

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

// UserRepository functions for user repository
type UserRepository interface {
	GetUserFromUsername(string) (*models.User, error)
	CreateAccount(*models.User) error
}