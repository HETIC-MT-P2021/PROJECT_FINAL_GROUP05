package models

type DownloadQueueMessage struct {
	Type      string  `json:"type"`
	MediaLink string  `json:"media_link"`
	Options   Options `json:"options"`
}

type MediaProcessingQueueMessage struct {
	FileName string  `json:"file_name"`
	Options  Options `json:"options"`
}

type Options struct {
	StartInSeconds    int `json:"start"`
	DurationInSeconds int `json:"duration"`
}

type JobCheckerMessage struct {
	DiscordMessageID string
	MediaURL         string
}
