package commands

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

func Dispatch(message *models.MediaProcessingQueueMessage) error {
	var command *logic.CommandImpl
	switch message.Options.Filter {
	case "gif":
		command = logic.NewCommandImpl(&implementation.VideoGIFCommand{
			Message: message,
		})
		break
	case "wm":
		/*command = logic.NewCommandImpl(&implementation.ImageCommand{
			Message: message,
		})*/
		break
	default:
		command = logic.NewCommandImpl(&implementation.VideoNoneCommand{
			Message: message,
		})
		break
	}
	return CommandBus.Dispatch(command)
}