package handlers

import (
	"errors"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/processing/video/ffmpeg"
)

// VideoWMCommandHandler allows make processing on a video
type VideoWMCommandHandler struct{
	Repo *ffmpeg.FfmpegCommandRepo
}

// VideoWMCommandHandler Creates an instance
func NewVideoWMCommandHandler(repo *ffmpeg.FfmpegCommandRepo) *VideoWMCommandHandler {
	return &VideoWMCommandHandler{
		Repo: repo,
	}
}

// Handle process a video
func (cHandler VideoWMCommandHandler) Handle(command logic.Command) error {
	switch command.Payload().(type) {
	case *implementation.VideoWMCommand:
		message := command.Payload().(*implementation.VideoWMCommand).Message
		return cHandler.Repo.MakeWaterMark(message)
	default:
		return errors.New("bad command type")
	}
}

