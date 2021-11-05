package ytb

import (
	"regexp"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
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

// Download and store into "downloads" folder
func (dl *ytbDownloadRepo) Download(youtubeURL string) (error, string) {
	var formats youtube.FormatList
	videoID := dl.getVideoID(youtubeURL)

	video, err := dl.Client.GetVideo(videoID)
	if err != nil {
		return err, ""
	}

	formats = video.Formats.Quality(utils.VIDEO_QUALITY).WithAudioChannels()

	if len(formats) < 1 {
		formats = video.Formats.WithAudioChannels()
		if len(formats) < 1 {
			return err, ""
		}
	}

	stream, _, err := dl.Client.GetStream(video, &formats[0])
	if err != nil {
		return err, ""
	}

	return utils.CreateMedia("mp4", stream)
}

const (
	DefaultYoutubeURL = `^https:\/\/www.youtube\.com\/watch\?v=`
	SharedYoutubeURL  = `^https:\/\/youtu\.be\/`
)

// getVideoID looking for DefaultYoutubeURL or SharedYoutubeURL on youtube url
func (dl *ytbDownloadRepo) getVideoID(youtubeURL string) string {
	defaultUrl := regexp.MustCompile(DefaultYoutubeURL)
	sharedUrl := regexp.MustCompile(SharedYoutubeURL)

	videoID := defaultUrl.ReplaceAll([]byte(youtubeURL), []byte(""))
	videoID = sharedUrl.ReplaceAll(videoID, []byte(""))
	return string(videoID)
}
