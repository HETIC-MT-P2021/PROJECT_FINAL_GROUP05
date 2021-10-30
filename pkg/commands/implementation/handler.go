package implementation

import (
	"errors"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/processing/video/ffmpeg"
)

// VideoCommandHandler allows make processing on a video
type VideoCommandHandler struct{
	Repo *ffmpeg.FfmpegCommandRepo
}

// NewVideoCommandHandler Creates an instance
func NewVideoCommandHandler(repo *ffmpeg.FfmpegCommandRepo) *VideoCommandHandler {
	return &VideoCommandHandler{
		Repo: repo,
	}
}

// Handle process a video
func (cHandler VideoCommandHandler) Handle(command logic.Command) error {
	switch command.Payload().(type) {
	case *VideoCommand:
		message := command.Payload().(*VideoCommand).Message
		return cHandler.Repo.MakeVideoClip(message)
	default:
		return errors.New("bad command type")
	}
}

