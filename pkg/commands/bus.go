package commands

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/implementation/handlers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/commands/logic"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/processing/video/ffmpeg"
)

var CommandBus *logic.CommandBus

func InitCommandBus() error {
	CommandBus = logic.NewCommandBus()

	videoProcess := ffmpeg.NewVideoCommandRepository()
	err := CommandBus.AddHandler(handlers.NewVideoNoneCommandHandler(videoProcess), &implementation.VideoNoneCommand{})
	err = CommandBus.AddHandler(handlers.NewVideoGIFCommandHandler(videoProcess), &implementation.VideoGIFCommand{})
	//err = CommandBus.AddHandler(handlers.NewVideoWMCommandHandler(videoProcess), &implementation.VideoWMCommand{})
	
	return err
}