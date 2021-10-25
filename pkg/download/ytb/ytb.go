package ytb

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

type ytbDownloadRepo struct {
	Client youtube.Client
}

// NewDownloadRepository creates new mysqlCommandRepo
func NewDownloadRepository(client youtube.Client) *ytbDownloadRepo {
	return &ytbDownloadRepo{
		Client: client,
	}
}

func (dl *ytbDownloadRepo) DownloadVideo(videoID string) error {
	video, err := dl.Client.GetVideo(videoID)
	if err != nil {
		return err
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := dl.Client.GetStream(video, &formats[0])
	if err != nil {
		return err
	}

	file, err := os.Create("go/downloads/video.mp4")
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
