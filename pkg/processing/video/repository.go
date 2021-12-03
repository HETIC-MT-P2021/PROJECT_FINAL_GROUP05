package video

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"

type VideoCommandRepository interface {
	MakeVideoClip(processing *models.MediaProcessingQueueMessage) error
	MakeGIF(processing *models.MediaProcessingQueueMessage) error
	MakeWaterMark(processing *models.MediaProcessingQueueMessage) error
}