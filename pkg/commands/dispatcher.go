package commands

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

func Dispatch(message *models.MediaProcessingQueueMessage) error {
	var command *logic.CommandImpl
	switch message.Type {
	case "audio":
		/*command = logic.NewCommandImpl(&implementation.AudioCommand{
			Message: message,
		})*/
		fmt.Println("Pas encore développé")
		break
	case "video":
		command = logic.NewCommandImpl(&implementation.VideoCommand{
			Message: message,
		})
		break
	case "image":
		/*command = logic.NewCommandImpl(&implementation.ImageCommand{
			Message: message,
		})*/
		fmt.Println("Pas encore développé")
		break
	default:
		return errors.New("Command doesn't exist")
	}
	return CommandBus.Dispatch(command)
}