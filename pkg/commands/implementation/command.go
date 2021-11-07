package implementation

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord/carlos"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

// VideoNoneCommand to cut a video without filter
type VideoNoneCommand struct {
	Message *models.MediaProcessingQueueMessage
	DiscordBot *carlos.CarlosBot
}

// VideoGIFCommand for gif processing
type VideoGIFCommand struct {
	Message *models.MediaProcessingQueueMessage
	DiscordBot *carlos.CarlosBot
}

// VideoWMCommand for watermark processing
type VideoWMCommand struct {
	Message *models.MediaProcessingQueueMessage
	DiscordBot *carlos.CarlosBot
}
