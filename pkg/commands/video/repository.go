package video

import "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/rabbit"

type CommandRepository interface {
	MakeVideoClip(opt rabbit.Options) error
}