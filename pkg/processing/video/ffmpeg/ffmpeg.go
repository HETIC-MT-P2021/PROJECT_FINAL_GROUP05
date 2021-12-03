package ffmpeg

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type FfmpegCommandRepo struct {}

// NewVideoCommandRepository creates new ffmpegCommandRepo
func NewVideoCommandRepository() *FfmpegCommandRepo {
	return &FfmpegCommandRepo{}
}

func (cmd *FfmpegCommandRepo) MakeVideoClip(processing *models.MediaProcessingQueueMessage) error {
	path := fmt.Sprintf("%sdownloads/", os.Getenv("BASE_DIR_IN_CONTAINER"))
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

func (cmd *FfmpegCommandRepo) MakeGIF(processing *models.MediaProcessingQueueMessage) error {
	path := fmt.Sprintf("%sdownloads/", os.Getenv("BASE_DIR_IN_CONTAINER"))
	if len(path) <= 0 {
		return errors.New("Path not found")
	}

	fileName := strings.Replace(processing.FileName, ".mp4", ".gif", -1)

	return ffmpeg.Input(
			fmt.Sprintf("%s%s", path, processing.FileName),
			ffmpeg.KwArgs{"ss": processing.Options.StartInSeconds}).
    Output(
			fmt.Sprintf("%sedit-%s", path, fileName), 
			ffmpeg.KwArgs{
				"t": processing.Options.DurationInSeconds,
				"s": "320x240",
				"pix_fmt": "rgb24",
				"r": "3",
			}).
    OverWriteOutput().Run()
}

func (cmd *FfmpegCommandRepo) MakeWaterMark(processing *models.MediaProcessingQueueMessage) error {
	/*downloadPath := fmt.Sprintf("%sdownloads/", os.Getenv("BASE_DIR_IN_CONTAINER"))
	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		return errors.New("File not found")
	}

	assetsPath := fmt.Sprintf("%sassets/", os.Getenv("BASE_DIR_IN_CONTAINER"))
	if _, err := os.Stat(assetsPath); os.IsNotExist(err) {
		return errors.New("File not found")
	}

	
	overlay := ffmpeg.Input(fmt.Sprintf("%scarlos_logo.png", assetsPath)).
		Filter("scale", ffmpeg.Args{"64:-1"})
	return ffmpeg.Filter(
    []*ffmpeg.Stream{
        ffmpeg.Input(fmt.Sprintf("%s%s", downloadPath, processing.FileName)),
        overlay,
    }, "overlay", ffmpeg.Args{"10:10"}, ffmpeg.KwArgs{"enable": "gte(t,1)"}).
    Output(fmt.Sprintf("%sedit-%s", downloadPath, processing.FileName)).
		OverWriteOutput().Run()*/
	return nil
}