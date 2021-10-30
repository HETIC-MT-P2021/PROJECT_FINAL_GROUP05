package mysql

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql/query"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

type mysqlCommandRepo struct {
	Conn *sql.DB
}

// NewCommandRepository creates new mysqlCommandRepo
func NewCommandRepository(conn *sql.DB) *mysqlCommandRepo {
	return &mysqlCommandRepo{
		Conn: conn,
	}
}

// CreateCommand to create command from bdd
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

	if _, err := stmt.Exec(
		command.Title,
		command.Command,
		command.IsActive,
		command.ServerID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetCommands to retrieve commands from bdd
func (db *mysqlCommandRepo) GetCommands() ([]models.Command, error) {
	results, err := db.Conn.Query(query.QUERY_FIND_COMMANDS)
	if err != nil {
		return []models.Command{}, err
	}

	commands := []models.Command{}

	for results.Next() {
		var cmd models.Command
		err = results.Scan(&cmd.Title, &cmd.Command, &cmd.IsActive, &cmd.ServerID)
		if err != nil {
			return []models.Command{}, err
		}
		commands = append(commands, cmd)
}

	return commands, nil
}