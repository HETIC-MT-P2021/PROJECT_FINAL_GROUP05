package ytb

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
	uuid "github.com/satori/go.uuid"
)

const DOCKER_CONTAINER_PATH = "go/downloads/"
const VIDEO_QUALITY = "medium"

type ytbDownloadRepo struct {
	Client youtube.Client
}

// NewDownloadRepository creates new mysqlCommandRepo
func NewDownloadRepository(client youtube.Client) *ytbDownloadRepo {
	return &ytbDownloadRepo{
		Client: client,
	}
}

// DownloadVideo and store into "downloads" folder
func (dl *ytbDownloadRepo) DownloadVideo(videoID string) error {
	video, err := dl.Client.GetVideo(videoID)
	if err != nil {
		return err
	}

	formats := video.Formats.Quality(VIDEO_QUALITY).WithAudioChannels()
	stream, _, err := dl.Client.GetStream(video, &formats[0])
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s%s.mp4", DOCKER_CONTAINER_PATH, uuid.NewV4().String()))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}
	return nil
}
