package video

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

type CommandRepository interface {
	MakeVideoClip(processing *models.MediaProcessingQueueMessage) error
}