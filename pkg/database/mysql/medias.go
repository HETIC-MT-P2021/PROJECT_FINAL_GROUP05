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
	stmt, err := repo.Conn.Prepare(query.QUERY_CREATE_MEDIA)
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

	return nil
}

func (repo *mysqlMediaRepo) UpdateMedia(mediaID int, media models.Media) error {
	stmt, err := repo.Conn.Prepare(query.QUERY_UPDATE_MEDIA)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(
		media.DiscordUrl,
		media.IsArchived,
		media.UserID,
		mediaID); err != nil {
		return err
	}

	return nil
}

// UpdateIsArchivedFieldMedia to update media field called is_archived from bdd
func (repo *mysqlMediaRepo) UpdateIsArchivedFieldMedia(mediaID string, isArchived bool) error {
	stmt, err := repo.Conn.Prepare(query.QUERY_UPDATE_FIELD_IS_ARCHIVED_MEDIA)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(isArchived, mediaID); err != nil {
		return err
	}

	return nil
}