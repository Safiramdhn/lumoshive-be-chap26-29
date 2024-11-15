package repositories

import (
	"database/sql"
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/utils"
	"time"

	"go.uber.org/zap"
)

type TourPlanRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewTourPlanRepository(db *sql.DB, log *zap.Logger) *TourPlanRepository {
	return &TourPlanRepository{DB: db, Logger: log}
}

var current = time.Now()

func (repo *TourPlanRepository) GetByEvent(event models.Event) ([]models.TourPlan, error) {
	var tourPlans []models.TourPlan
	repo.Logger.Info("Get tour plan by event id", zap.Duration("duration", time.Since(current)))
	query := `SELECT id, activities, description FROM tour_plan WHERE event_id = $1`
	rows, err := repo.DB.Query(query, event.ID)
	if err != nil {
		repo.Logger.Error("Error scanning row: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var activities []byte
		tourPlan := models.TourPlan{}
		err := rows.Scan(&tourPlan.ID, &activities, &tourPlan.Description)
		if err != nil {
			repo.Logger.Error("Error scanning row: " + err.Error())
			return nil, err
		}
		tourPlan.Activities = utils.ParseActivities(activities)
		tourPlans = append(tourPlans, tourPlan)
	}

	return tourPlans, nil
}
