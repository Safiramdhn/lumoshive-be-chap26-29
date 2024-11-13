package repositories

import (
	"database/sql"
	"fmt"
	"golang-beginner-chap28/models"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Repository[T any] interface {
	GetTourData(dateFilter time.Time, sortByPrice string, page, pageSize int) ([]models.TourData, int, error)
}

type repository[T any] struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewRepository[T any](db *sql.DB, logger *zap.Logger) Repository[T] {
	return &repository[T]{DB: db, Logger: logger}
}

// GetTourData implements Repository.
func (r *repository[T]) GetTourData(dateFilter time.Time, sortByPrice string, page int, pageSize int) ([]models.TourData, int, error) {
	var whereConditions []string
	args := []interface{}{}
	var totalItems int
	index := 1

	// Add date filter if present
	if !dateFilter.IsZero() {
		whereConditions = append(whereConditions, fmt.Sprintf("e.date = $%d", index))
		args = append(args, dateFilter)
		index++
	}

	// Combine conditions to form the WHERE clause
	whereClause := ""
	if len(whereConditions) > 0 {
		whereClause = " WHERE " + strings.Join(whereConditions, " AND ")
	}

	// Start SQL query
	sqlStatement := `SELECT t.id, e.id, e.date, e.title, e.price, d.id, d.name, d.location,
                    COALESCE(AVG(r.rating), 0) AS average_rating,
                    COUNT(t.id) AS total_transactions 
                    FROM event e 
                    JOIN transaction t ON t.event_id = e.id 
                    JOIN destination d ON e.destination_id = d.id 
                    LEFT JOIN review r ON r.transaction_id = t.id
                    ` + whereClause + `
                    GROUP BY t.id, e.id, d.id`

	// Add sorting by price if specified
	if sortByPrice == "asc" {
		sqlStatement += " ORDER BY e.price ASC"
	} else if sortByPrice == "desc" {
		sqlStatement += " ORDER BY e.price DESC"
	}

	// Add pagination using LIMIT and OFFSET
	sqlStatement += fmt.Sprintf(" LIMIT $%d OFFSET $%d", index, index+1)
	args = append(args, pageSize, (page-1)*pageSize)

	// Execute the query
	r.Logger.Info("Executing query select", zap.String("query", sqlStatement))
	rows, err := r.DB.Query(sqlStatement, args...)
	if err != nil {
		r.Logger.Error("Error executing query "+err.Error(), zap.String("repository", "tourDara"), zap.String("query", sqlStatement))
		return nil, 0, err
	}
	defer rows.Close()

	var result []models.TourData
	for rows.Next() {
		var transaction models.Transaction
		var event models.Event
		var destination models.Destination
		var avgRating float64
		var totalTransactions int

		err := rows.Scan(&transaction.ID, &event.ID, &event.Date, &event.Title, &event.Price,
			&destination.ID, &destination.Name, &destination.Location,
			&avgRating, &totalTransactions)
		if err != nil {
			r.Logger.Error("Error scanning row "+err.Error(), zap.String("repository", "tourData"))
			return nil, 0, err
		}

		result = append(result, models.TourData{
			Transaction:       transaction,
			Event:             event,
			Destination:       destination,
			AverageRating:     avgRating,
			TotalTransactions: totalTransactions,
		})
	}
	r.Logger.Info("Executing query select", zap.String("query", "SELECT COUNT(*) FROM event"))
	err = r.DB.QueryRow("SELECT COUNT(*) FROM event").Scan(&totalItems)
	if err != nil {
		r.Logger.Error("Error counting total items "+err.Error(), zap.String("repository", "tourData"), zap.String("query", "SELECT COUNT(*) FROM event"))
		return nil, 0, err
	}

	return result, totalItems, nil
}
