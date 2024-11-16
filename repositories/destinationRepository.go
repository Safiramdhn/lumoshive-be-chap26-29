package repositories

import (
	"database/sql"
	"golang-beginner-chap28/models"

	"go.uber.org/zap"
)

type DestinationRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewDestinationRepository(db *sql.DB, log *zap.Logger) *DestinationRepository {
	return &DestinationRepository{DB: db, Logger: log}
}

func (repo *DestinationRepository) GetByEvent(event models.Event) (*models.Destination, error) {
	var destination models.Destination
	sqlStatement := `SELECT id, location, description, map_url FROM destination`
	err := repo.DB.QueryRow(sqlStatement).Scan(&destination.ID, &destination.Location, &destination.Description, &destination.MapUrl)
	if err != nil {
		repo.Logger.Error("Error scanning row: " + err.Error())
		return nil, err
	}

	return &destination, nil
}
