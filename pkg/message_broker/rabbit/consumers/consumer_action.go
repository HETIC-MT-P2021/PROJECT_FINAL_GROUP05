package consumers

import (
	"encoding/json"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/message_broker/rabbit/producers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

type CallMediaProcessingProducer struct {
	Message models.MediaProcessingMessage
}

func NewCallMediaProcessingProducer(message models.MediaProcessingMessage) *CallMediaProcessingProducer {
	return &CallMediaProcessingProducer{
		Message: message,
	}
}

func (producer *CallMediaProcessingProducer) Execute(downloadMessageByte []byte) error {
	var downloadMessage *models.DownloadMessage
	err := json.Unmarshal(downloadMessageByte, &downloadMessage)
	if err != nil {
		return err
	}

	message := producer.Message
	message.Options = downloadMessage.Options
	producers.MediaProcessingProducer(message)

	return nil
}