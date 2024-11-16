package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewPostgresDB() *sql.DB {
	log := IntiLogger()
	start := time.Now()
	configViper := InitViper()

	dbName := configViper.GetString("DB_NAME")
	dbUser := configViper.GetString("DB_USER")
	dbPassword := configViper.GetString("DB_PASSWORD")
	dbHost := configViper.GetString("HOST")

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", dbUser, dbName, dbPassword, dbHost)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Error opening database connection: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Error("Error connecting to the database: " + err.Error())
	}

	log.Info("Connected to database", zap.String("config_type", "viper"),
		zap.Duration("duration", time.Since(start)),
	)
	return db
}
