package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
)

const (
	attemptsDBConnexion = 5
	waitForConnexion    = 5
)

// Config for DB connection
type Config struct {
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"MYSQL_DATABASE"`
	DbUser     string `env:"MYSQL_USER"`
	DbPassword string `env:"MYSQL_PASSWORD"`
	DbConn     *sql.DB
}

// Connect connection to database
func Connect() (*sql.DB, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("%+v", err)
	}
	dsn := cfg.DbUser + ":" + cfg.DbPassword + "@" + cfg.DbHost + "/" + cfg.
		DbName + "?parseTime=true&charset=utf8"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	for index := 1; index <= attemptsDBConnexion; index++ {
		err = db.Ping()
		if err != nil {
			if index < attemptsDBConnexion {
				log.Printf("db connection failed, %d retry : %v", index, err)
				time.Sleep(waitForConnexion * time.Second)
			}
			continue
		} else {
			return db, nil
		}
	}

	return nil, nil
}