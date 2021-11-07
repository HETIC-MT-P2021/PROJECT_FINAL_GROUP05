package consumers

import (
	"encoding/json"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord/carlos"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/media_download/video/ytb"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
	"github.com/kkdai/youtube/v2"
)

type CallMediaProcessingProducer struct {
	Body []byte
}

func NewCallMediaProcessingProducer() *CallMediaProcessingProducer {
	return &CallMediaProcessingProducer{}
}

func (producer *CallMediaProcessingProducer) SetBody(body []byte) {
	producer.Body = body
}

func (producer *CallMediaProcessingProducer) Execute() error {
	var downloadMessage *models.DownloadQueueMessage
	err := json.Unmarshal(producer.Body, &downloadMessage)
	if err != nil {
		return err
	}

	download := ytb.NewDownloadRepository(youtube.Client{})
	err, fileName := download.Download(downloadMessage.MediaLink)
	if err != nil {
		return err
	}

	message := models.MediaProcessingQueueMessage{
		Type: downloadMessage.Type,
		FileName: fileName,
		Options: downloadMessage.Options,
		DiscordInfo: downloadMessage.DiscordInfo,
	}

	producers.MediaProcessingProducer(message)

	return nil
}

type MediaProcessingAction struct {
	Body []byte
}

func NewMediaProcessingAction() *MediaProcessingAction {
	return &MediaProcessingAction{}
}

func (producer *MediaProcessingAction) Execute() error {
	var processingMessage *models.MediaProcessingQueueMessage
	err := json.Unmarshal(producer.Body, &processingMessage)
	if err != nil {
		return err
	}

	err = commands.Dispatch(processingMessage)
	if err != nil {
		return err
	}
	
	carlosBot := carlos.NewDiscordRepository()
	err = carlosBot.InitBotWithoutHandler()
	if err != nil {
		return err
	}

	carlosBot.ChannelID = processingMessage.DiscordInfo.ChannelID
	carlosBot.ServerID = processingMessage.DiscordInfo.ServerID
	carlosBot.MessageID = processingMessage.DiscordInfo.MessageID
	message, err := carlosBot.UpdateCarlosIsProcessingMessage(processingMessage.FileName)

	mediaRepo := mysql.NewMediaRepository(database.DB)
	now := utils.NewDateNow()
	media := models.Media{
		DiscordUrl: message.Attachments[0].URL,
		IsArchived: true,
		UserID: message.Author.ID,
		ServerID: processingMessage.DiscordInfo.ServerID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return mediaRepo.CreateMedia(media)
}

func (producer *MediaProcessingAction) SetBody(body []byte) {
	producer.Body = body
}