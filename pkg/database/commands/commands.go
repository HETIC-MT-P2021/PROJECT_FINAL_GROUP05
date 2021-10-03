package commands

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/query"
)

type Command struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Command string `json:"command"`
	IsChecked bool `json:"is_checked"`
}

type CommandsImpl struct {
	Conn *sql.DB
}

func NewCommandsRepository(conn *sql.DB) *CommandsImpl {
	return &CommandsImpl{
		Conn: conn,
	}
}

func (db *CommandsImpl) CreateCommand(command Command) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query.QUERY_CREATE_COMMAND)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(command.Title, command.Command, command.IsChecked); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *CommandsImpl) GetCommands() ([]Command, error) {
	results, err := db.Conn.Query(query.QUERY_FIND_COMMANDS)
	if err != nil {
		return []Command{}, err
	}

	commands := []Command{}

	for results.Next() {
		var cmd Command
		err = results.Scan(&cmd.Title, &cmd.Command, &cmd.IsChecked)
		if err != nil {
			return []Command{}, err
		}
		commands = append(commands, cmd)
}

	return commands, nil
}