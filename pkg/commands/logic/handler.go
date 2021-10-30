package logic

// CommandHandler Allows to manage Command
type CommandHandler interface {
	Handle(Command) error
}