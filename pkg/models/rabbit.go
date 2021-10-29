package models

import "github.com/bwmarrin/discordgo"

type DownloadQueueMessage struct {
		Type      			string  								`json:"type"`
		MediaLink 			string  								`json:"media_link"`
		Options   			Options 								`json:"options"`
		DiscordInfo 		JobCheckerQueueMessage 	`json:"discord_info"`
}

type MediaProcessingQueueMessage struct {
	FileName 				string  								`json:"file_name"`
	Options  				Options 								`json:"options"`
	DiscordInfo 		JobCheckerQueueMessage 	`json:"discord_info"`
}

type JobCheckerQueueMessage struct {
	ChannelID 				string  				`json:"channel_id"`
	MessageID  				string 					`json:"message_id"`
}

type Options struct {
	StartInSeconds    int `json:"start"`
	DurationInSeconds int `json:"duration"`
}

type DiscordSession struct {
	Session 		*discordgo.Session	`json:"session"`
	ChannelID 	string 							`json:"channel_id"`
	MessageID 	string 							`json:"message"`
}
