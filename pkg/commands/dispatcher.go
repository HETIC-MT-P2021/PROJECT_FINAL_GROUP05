package commands

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/discord/carlos"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

type DispatchParams struct {
	Message *models.MediaProcessingQueueMessage
	DiscordBot *carlos.CarlosBot
}

func Dispatch(params DispatchParams) error {
	var command *logic.CommandImpl
	switch params.Message.Options.Filter {
	case "gif":
		command = logic.NewCommandImpl(&implementation.VideoGIFCommand{
			Message: params.Message,
			DiscordBot: params.DiscordBot,
		})
		break
	/*case "wm":
		command = logic.NewCommandImpl(&implementation.VideoWMCommand{
			Message: params.Message,
			DiscordBot: params.DiscordBot,
		})
		break*/
	default:
		command = logic.NewCommandImpl(&implementation.VideoNoneCommand{
			Message: params.Message,
			DiscordBot: params.DiscordBot,
		})
		break
	}
	return CommandBus.Dispatch(command)
}