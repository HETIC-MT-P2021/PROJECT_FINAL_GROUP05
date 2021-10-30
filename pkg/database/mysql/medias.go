package mysql

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/database/mysql/query"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

type mysqlMediaRepo struct {
	Conn *sql.DB
}

// NewMediaRepository creates new mysqlMediaRepo
func NewMediaRepository(conn *sql.DB) *mysqlMediaRepo {
	return &mysqlMediaRepo{
		Conn: conn,
	}
}

func (repo *mysqlMediaRepo) CreateMedia(media models.Media) error {
	tx, err := repo.Conn.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query.QUERY_CREATE_MEDIA)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(
		media.DiscordUrl,
		media.IsArchived,
		media.UserID,
		media.ServerID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}