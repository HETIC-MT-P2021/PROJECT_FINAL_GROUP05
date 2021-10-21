package mysql

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql/query"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

type mysqlCommandRepo struct {
	Conn *sql.DB
}

// NewCommandsRepository creates new mysqlCommandRepo
func NewCommandsRepository(conn *sql.DB) *mysqlCommandRepo {
	return &mysqlCommandRepo{
		Conn: conn,
	}
}

// GetUserFromUsername to create command from bdd
func (db *mysqlCommandRepo) CreateCommand(command models.Command) error {
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

// GetUserFromUsername to retrieve commands from bdd
func (db *mysqlCommandRepo) GetCommands() ([]models.Command, error) {
	results, err := db.Conn.Query(query.QUERY_FIND_COMMANDS)
	if err != nil {
		return []models.Command{}, err
	}

	commands := []models.Command{}

	for results.Next() {
		var cmd models.Command
		err = results.Scan(&cmd.Title, &cmd.Command, &cmd.IsChecked)
		if err != nil {
			return []models.Command{}, err
		}
		commands = append(commands, cmd)
}

	return commands, nil
}