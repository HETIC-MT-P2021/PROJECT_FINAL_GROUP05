package ytb

import (
	"strings"

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

// DownloadVideo and store into "downloads" folder
func (dl *ytbDownloadRepo) Download(youtubeURL string) (error, string) {
	videoID := dl.getVideoID(youtubeURL)
	video, err := dl.Client.GetVideo(videoID)
	if err != nil {
		return err, ""
	}

	formats := video.Formats.Quality(utils.VIDEO_QUALITY).WithAudioChannels()
	stream, _, err := dl.Client.GetStream(video, &formats[0])
	if err != nil {
		return err, ""
	}

	return utils.CreateMedia("mp4", stream) 
}

const YOUTUBE_PATTERN_URL = "https://www.youtube.com/watch?v="

// getVideoID looking for ?v=id on youtube url
func (dl *ytbDownloadRepo) getVideoID(youtubeURL string) string {
	return strings.Replace(youtubeURL, YOUTUBE_PATTERN_URL, "", -1)
}
