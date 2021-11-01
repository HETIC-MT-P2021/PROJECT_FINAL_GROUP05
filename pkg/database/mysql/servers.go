package mysql

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql/query"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

type mysqlServerRepo struct {
	Conn *sql.DB
	CommandRepo database.CommandRepository
}

// NewServerRepository creates new mysqlServerRepo
func NewServerRepository(conn *sql.DB, repo database.CommandRepository) *mysqlServerRepo {
	return &mysqlServerRepo{
		Conn: conn,
		CommandRepo: repo,
	}
}

// CreateServer to create server from bdd
func (repo *mysqlServerRepo) CreateServer(server models.Server) error {
	stmt, err := repo.Conn.Prepare(query.QUERY_CREATE_SERVER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(
		server.ID,
		server.ServerName,
		server.CreatedAt); err != nil {
		return err
	}

	err = repo.CommandRepo.CreateDefaultCommandsInServer(server.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetServers to retrieve servers from bdd
func (repo *mysqlServerRepo) GetServers() ([]models.Server, error) {
	results, err := repo.Conn.Query(query.QUERY_FIND_SERVERS)
	if err != nil {
		return []models.Server{}, err
	}

	servers := []models.Server{}
	for results.Next() {
		var server models.Server
		err = results.Scan(
			&server.ID, 
			&server.ServerName, &server.CreatedAt)
		if err != nil {
			return []models.Server{}, err
		}
		servers = append(servers, server)
	}

	return servers, nil
}

// GetServer to retrieve server with media and commands data from bdd
func (repo *mysqlServerRepo) GetServer(serverID string) (*models.ServerCommandsAndMedias, error) {
	rows, err := repo.Conn.Query(query.QUERY_FIND_SERVER_WITH_MADIA_AND_CMDS, serverID)
	if err != nil {
		return &models.ServerCommandsAndMedias{}, err
	}
	server := &models.ServerCommandsAndMedias{}
	for rows.Next() {
    command := &models.Command{}
    media := &models.Media{}
    err = rows.Scan(
      &server.ServerName,
      &server.CreatedAt,
			&command.Title,
			&command.Command,
			&command.IsActive,
			&media.DiscordUrl,
			&media.IsArchived,
			&media.UserID,
    )
    server.Commands = append(server.Commands, *command)
    server.Medias = append(server.Medias, *media)
  }

	return server, nil
}

func (repo *mysqlServerRepo) GetServerMedias(serverID string) ([]models.Media, error) {
	rows, err := repo.Conn.Query(query.QUERY_FIND_SERVER_MEDIAS, serverID)
	if err != nil {
		return []models.Media{}, err
	}

	medias := []models.Media{}
	for rows.Next() {
		media := models.Media{}
    err = rows.Scan(
			&media.DiscordUrl,
			&media.IsArchived,
			&media.UserID,
    )
		medias = append(medias, media)
  }

	return medias, nil
}

func (repo *mysqlServerRepo) GetServerCommands(serverID string) ([]models.Command, error) {
	rows, err := repo.Conn.Query(query.QUERY_FIND_SERVER_COMMANDS, serverID)
	if err != nil {
		return []models.Command{}, err
	}

	commands := []models.Command{}
	for rows.Next() {
		command := models.Command{}
    err = rows.Scan(
			&command.Title,
			&command.Command,
			&command.IsActive,
    )
		commands = append(commands, command)
  }

	return commands, nil
}
