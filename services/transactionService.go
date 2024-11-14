package services

import (
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/repositories"
)

type TransactionService struct {
	TransactionRepo repositories.TransactionRepo
}

func NewTransactionService(transactionRepo repositories.TransactionRepo) *TransactionService {
	return &TransactionService{transactionRepo}
}

func (ts *TransactionService) CreateTransaction(transactionInput models.Transaction) (*models.Transaction, error) {
	return ts.TransactionRepo.Create(transactionInput)
}
