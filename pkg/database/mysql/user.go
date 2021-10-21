package mysql

import (
	"database/sql"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/query"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

// mysqlUserRepo struct for db connection
type mysqlUserRepo struct {
	DbConn *sql.DB
}

// NewUserRepositoryImpl creates new mysqlUserRepo
func NewUserRepositoryImpl(dbConn *sql.DB) *mysqlUserRepo {
	return &mysqlUserRepo{
		DbConn: dbConn,
	}
}

// GetUserFromUsername to retrieve user from bdd
func (r *mysqlUserRepo) GetUserFromUsername(username string) (*models.User, error) {
	var (
		Email    string
		Password string
		RoleInt  uint8
	)

	stmt, err := r.DbConn.Prepare(query.QUERY_FIND_USERS_BY_USERNAME)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&Email, &Password, &RoleInt)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    Email,
		Password: Password,
		Role:     models.Role(RoleInt),
	}

	return user, nil
}

// CreateAccount to create an account with role operator
func (r *mysqlUserRepo) CreateAccount(userInput models.RequestRegister) (err error) {
	tx, err := r.DbConn.Begin()
	if err != nil {
		return err
	}

	stmt, es := tx.Prepare(query.QUERY_CREATE_ACCOUNT)
	if es != nil {
		return es
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userInput.Email, userInput.Username, userInput.Password); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}