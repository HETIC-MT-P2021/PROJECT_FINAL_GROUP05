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
		cmd := command.Payload().(*implementation.VideoNoneCommand)
		err := cHandler.Repo.MakeVideoClip(cmd.Message)
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

