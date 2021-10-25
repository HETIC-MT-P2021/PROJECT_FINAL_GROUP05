package download

// DownloadRepository functions for youtube videos
type DownloadRepository interface {
	DownloadVideo(videoID string) error
}