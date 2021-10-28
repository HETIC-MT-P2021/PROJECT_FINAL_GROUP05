package models

type DownloadMessage struct {
	Type 			string      `json:"type"`
	MediaLink string 			`json:"media_link"`
	Options 	Options 		`json:"options"`
}

type MediaProcessingMessage struct {
	FileName 	string      `json:"file_name"`
	Options 	Options 		`json:"options"`
}

type Options struct {
	StartInSeconds			string 			`json:"start"`
	DurationInSeconds 	string 			`json:"duration"`
}

type JobCheckerMessage struct {
	DiscordMessageID string
	MediaURL string
}