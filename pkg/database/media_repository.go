package database

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

// mediaRepository functions for media repository
type mediaRepository interface {
	CreateMedia(command models.Media) error
	GetMedias() ([]models.Media, error)
	UpdateMedia(mediaID int, command models.Media) error
	UpdateIsArchivedFieldMedia(mediaID int, isArchived bool) error
}