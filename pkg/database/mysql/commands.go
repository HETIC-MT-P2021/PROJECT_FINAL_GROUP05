package mysql

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql/query"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/utils"
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
		command.CreatedAt,
		command.UpdatedAt,
		command.ServerID); err != nil {
		return err
	}

	return nil
}

var DefaultCommand []models.Command = []models.Command {
	models.Command{
		Title: "Extrait vidéo",
		Command: "video -s 'X' -d 'Y' -i 'input'",		
		IsActive: true,
		CreatedAt: utils.NewDateNow(),
		UpdatedAt: utils.NewDateNow(),
	},
	models.Command{
		Title: "Extrait audio",
		Command: "audio -s 'X' -d 'Y' -i 'input'",		
		IsActive: true,
		CreatedAt: utils.NewDateNow(),
		UpdatedAt: utils.NewDateNow(),
	},
}

func (repo *mysqlCommandRepo) CreateDefaultCommandsInServer(serverID string) error {
	for _, command := range DefaultCommand {
		stmt, err := repo.Conn.Prepare(query.QUERY_CREATE_COMMAND)
		if err != nil {
			return err
		}
		defer stmt.Close()
	
		if _, err := stmt.Exec(
			command.Title,
			command.Command,
			command.IsActive,
			serverID,
			command.CreatedAt,
			command.UpdatedAt); err != nil {
			return err
		}
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
		err = results.Scan(
			&cmd.Title,
			&cmd.Command,
			&cmd.IsActive,
			&cmd.ServerID,
			&cmd.CreatedAt,
			&cmd.UpdatedAt)
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
		&command.UpdatedAt,
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

	if _, err := stmt.Exec(isActive, utils.NewDateNow(), commandID); err != nil {
		return err
	}

	return nil
}