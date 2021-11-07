package handlers

import (
	"errors"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/processing/video/ffmpeg"
)

// VideoGIFCommandHandler allows make processing on a video
type VideoGIFCommandHandler struct{
	Repo *ffmpeg.FfmpegCommandRepo
}

// NewVideoGIFCommandHandler Creates an instance
func NewVideoGIFCommandHandler(repo *ffmpeg.FfmpegCommandRepo) *VideoGIFCommandHandler {
	return &VideoGIFCommandHandler{
		Repo: repo,
	}
}

// Handle process a video
func (cHandler VideoGIFCommandHandler) Handle(command logic.Command) error {
	switch command.Payload().(type) {
	case *implementation.VideoGIFCommand:
		message := command.Payload().(*implementation.VideoGIFCommand).Message
		return cHandler.Repo.MakeGIF(message)
	default:
		return errors.New("bad command type")
	}
}

