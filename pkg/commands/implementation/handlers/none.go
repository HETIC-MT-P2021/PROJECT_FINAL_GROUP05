package handlers

import (
	"errors"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/processing/video/ffmpeg"
)

// VideoNoneCommandHandler allows make processing on a video
type VideoNoneCommandHandler struct{
	Repo *ffmpeg.FfmpegCommandRepo
}

// NewVideoNoneCommandHandler Creates an instance
func NewVideoNoneCommandHandler(repo *ffmpeg.FfmpegCommandRepo) *VideoNoneCommandHandler {
	return &VideoNoneCommandHandler{
		Repo: repo,
	}
}

// Handle process a video
func (cHandler VideoNoneCommandHandler) Handle(command logic.Command) error {
	switch command.Payload().(type) {
	case *implementation.VideoNoneCommand:
		message := command.Payload().(*implementation.VideoNoneCommand).Message
		return cHandler.Repo.MakeVideoClip(message)
	default:
		return errors.New("bad command type")
	}
}

