package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewPostgresDB() *sql.DB {
	log := IntiLogger()
	connStr := "user=postgres dbname=travel sslmode=disable password=postgres host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Error opening database connection: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Error("Error connecting to the database: " + err.Error())
	}

	return db
}
