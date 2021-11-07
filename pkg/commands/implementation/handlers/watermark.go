package handlers

import (
	"errors"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/processing/video/ffmpeg"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
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
		cmd := command.Payload().(*implementation.VideoWMCommand)
		err := cHandler.Repo.MakeWaterMark(cmd.Message)
		if err != nil {
			return err
		}

		message, err := cmd.DiscordBot.UpdateCarlosIsProcessingMessage(cmd.Message.FileName)
		if err != nil {
			return err
		}

		now := utils.NewDateNow()
		media := models.Media{
			DiscordUrl: message.Attachments[0].URL,
			IsArchived: true,
			UserID: message.Author.ID,
			ServerID: cmd.DiscordBot.ServerID,
			CreatedAt: now,
			UpdatedAt: now,
		}
		return database.MediaRepo.CreateMedia(media)
	default:
		return errors.New("bad command type")
	}
}

