package repositories

import (
	"database/sql"
	"golang-beginner-chap28/models"

	"go.uber.org/zap"
)

type EventRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewEventRepository(db *sql.DB, log *zap.Logger) *EventRepository {
	return &EventRepository{DB: db, Logger: log}
}

func (repo *EventRepository) GetById(id int) (*models.Event, error) {
	var event models.Event
	query := `SELECT id, title, description, photo_url, price, date FROM events WHERE id = $1`
	err := repo.DB.QueryRow(query, id).Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.PhotoUrl, &event.Price, &event.Date)
	if err != nil {
		repo.Logger.Error("Error scanning row: " + err.Error())
		return nil, err
	}

	return &event, nil
}
