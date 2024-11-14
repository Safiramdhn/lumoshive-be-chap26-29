package repositories

import (
	"database/sql"
	"golang-beginner-chap28/models"

	"go.uber.org/zap"
)

type CustomerRepo struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewCustomerRepo(db *sql.DB, log *zap.Logger) *CustomerRepo {
	return &CustomerRepo{DB: db, Logger: log}
}

// Create implements MainRepo.
func (c *CustomerRepo) Create(customerInput models.Customer) (*models.Customer, error) {
	tx, err := c.DB.Begin()
	if err != nil {
		c.Logger.Error("Failed to start transaction", zap.Error(err))
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			c.Logger.Error("Recovered from panic: ", zap.Any("panic", p), zap.String("repository", "Transaction"))
			tx.Rollback()
			panic(p)
		} else if err != nil {
			c.Logger.Error("Error committing transaction: ", zap.Error(err), zap.String("repository", "Transaction"))
			tx.Rollback()
		}
	}()

	sqlStatement := `INSERT INTO customer (name, email, phone_number) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(sqlStatement, customerInput.Name, customerInput.Email, customerInput.Phone).Scan(&customerInput.ID)
	if err != nil {
		c.Logger.Error("Error inserting customer: ", zap.Error(err), zap.String("repository", "CustomerRepo"), zap.String("function", "Create"), zap.String("query", sqlStatement))
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		c.Logger.Error("Error committing transaction: ", zap.Error(err), zap.String("repository", "CustomerRepo"), zap.String("function", "Create"))
		return nil, err
	}
	return &customerInput, nil
}
