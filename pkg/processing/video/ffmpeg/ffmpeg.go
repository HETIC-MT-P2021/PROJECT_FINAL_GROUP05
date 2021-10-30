package ffmpeg

import (
	"errors"
	"fmt"
	"os"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type FfmpegCommandRepo struct {}

// NewVideoCommandRepository creates new ffmpegCommandRepo
func NewVideoCommandRepository() *FfmpegCommandRepo {
	return &FfmpegCommandRepo{}
}

func (cmd *FfmpegCommandRepo) MakeVideoClip(processing *models.MediaProcessingQueueMessage) error {
	path := os.Getenv("DOCKER_CONTAINER_PATH")
	if len(path) <= 0 {
		return errors.New("Path not found")
	}
	
	return ffmpeg.Input(
			fmt.Sprintf("%s%s", path, processing.FileName),
			ffmpeg.KwArgs{"ss": processing.Options.StartInSeconds}).
    Output(
			fmt.Sprintf("%sedit-%s", path, processing.FileName), 
			ffmpeg.KwArgs{"t": processing.Options.DurationInSeconds}).
		OverWriteOutput().Run()
}