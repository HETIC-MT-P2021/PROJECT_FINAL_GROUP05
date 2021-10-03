package commands

type CommandsRepository interface {
	CreateCommand(command Command) error
	GetCommands() ([]Command, error)
}