package consumers

import (
	"encoding/json"
	"fmt"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/video/ffmpeg"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/media_download/video/ytb"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
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
		FileName: fileName,
		Options: downloadMessage.Options,
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

	processing := ffmpeg.NewCommandRepository()
	err = processing.MakeVideoClip(processingMessage)
	if err != nil {
		return err
	}
	fmt.Println("Here is discord message update")
	return nil
}

func (producer *MediaProcessingAction) SetBody(body []byte) {
	producer.Body = body
}