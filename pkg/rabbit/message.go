package rabbit

type QueueMessage struct {
	Type 			string      `json:"type"`
	MediaLink string 			`json:"media_link"`
	Options 	Options 		`json:"options"`
}

type Options struct {
	StartInSeconds			string 			`json:"start"`
	DurationInSeconds 	string 			`json:"duration"`
}