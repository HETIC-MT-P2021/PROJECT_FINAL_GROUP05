package video

// DownloadVideoRepository functions for youtube videos
type DownloadVideoRepository interface {
	Download(videoID string) (error, string)
}