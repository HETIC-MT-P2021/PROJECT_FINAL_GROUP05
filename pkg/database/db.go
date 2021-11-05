package database

import "database/sql"

var (
	DB *sql.DB
	CommandRepo CommandRepository
	MediaRepo MediaRepository
	ServerRepo ServerRepository
)