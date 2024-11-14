package repositories

import (
	"database/sql"
	"golang-beginner-chap28/models"

	"go.uber.org/zap"
)

type TransactionRepo struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewTransactionRepo(db *sql.DB, log *zap.Logger) *TransactionRepo {
	return &TransactionRepo{DB: db, Logger: log}
}

// Create implements MainRepo.
func (tr *TransactionRepo) Create(transactionInput models.Transaction) (*models.Transaction, error) {
	tx, err := tr.DB.Begin()
	if err != nil {
		tr.Logger.Error("Error starting transaction: ", zap.Error(err), zap.String("repository", "Transaction"))
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tr.Logger.Error("Recovered from panic: ", zap.Any("panic", p), zap.String("repository", "Transaction"))
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tr.Logger.Error("Error committing transaction: ", zap.Error(err), zap.String("repository", "Transaction"))
			tx.Rollback()
		}
	}()

	sqlStatement := `INSERT INTO transaction (event_id, customer_id, status, price) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(sqlStatement, transactionInput.EventID, transactionInput.CustomerID, "Pending", transactionInput.Price)
	if err != nil {
		tr.Logger.Error("Error executing SQL statement: ", zap.Error(err), zap.String("repository", "Transaction"))
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tr.Logger.Error("Error committing transaction: ", zap.Error(err), zap.String("repository", "Transaction"))
		return nil, err
	}
	return &transactionInput, nil
}
