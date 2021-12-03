package handlers

import (
	"errors"
	"strings"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/processing/video/ffmpeg"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
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
		cmd := command.Payload().(*implementation.VideoGIFCommand)
		err := cHandler.Repo.MakeGIF(cmd.Message)
		if err != nil {
			return err
		}

		fileName := strings.Replace(cmd.Message.FileName, ".mp4", ".gif", -1)
		message, err := cmd.DiscordBot.UpdateCarlosIsProcessingMessage(fileName)
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

