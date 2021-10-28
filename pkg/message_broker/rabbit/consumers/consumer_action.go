package consumers

import (
	"encoding/json"
	"fmt"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

type CallMediaProcessingProducer struct {
	MediaProcessingMessage models.MediaProcessingMessage
	Body []byte
}

func NewCallMediaProcessingProducer(message models.MediaProcessingMessage) *CallMediaProcessingProducer {
	return &CallMediaProcessingProducer{
		MediaProcessingMessage: message,
	}
}

func (producer *CallMediaProcessingProducer) SetBody(body []byte) {
	producer.Body = body
}

func (producer *CallMediaProcessingProducer) Execute() error {
	var downloadMessage *models.DownloadMessage
	err := json.Unmarshal(producer.Body, &downloadMessage)
	if err != nil {
		return err
	}

	message := producer.MediaProcessingMessage
	message.Options = downloadMessage.Options
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
	fmt.Println("Here is discord message update")
	return nil
}

func (producer *MediaProcessingAction) SetBody(body []byte) {
	producer.Body = body
}