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
func (repo *mysqlCommandRepo) CreateCommand(command models.Command) error {
	stmt, err := repo.Conn.Prepare(query.QUERY_CREATE_COMMAND)
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

	return nil
}

// GetCommands to retrieve commands from bdd
func (repo *mysqlCommandRepo) GetCommands() ([]models.Command, error) {
	results, err := repo.Conn.Query(query.QUERY_FIND_COMMANDS)
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

// UpdateCommand to update command info from bdd
func (repo *mysqlCommandRepo) UpdateCommand(commandID string, command models.Command) error {
	stmt, err := repo.Conn.Prepare(query.QUERY_UPDATE_COMMAND)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(
		command.Title,
		command.Command,
		command.IsActive,
		command.ServerID,
		commandID); err != nil {
		return err
	}

	return nil
}

// UpdateIsActiveFieldCommand to update command field called is_active from bdd
func (repo *mysqlCommandRepo) UpdateIsActiveFieldCommand(commandID string, isActive bool) error {
	stmt, err := repo.Conn.Prepare(query.QUERY_UPDATE_FIELD_IS_ACTIVE_COMMAND)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(isActive, commandID); err != nil {
		return err
	}

	return nil
}