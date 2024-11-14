package handlers

import (
	"encoding/json"
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type TransactionHandler struct {
	TransactionService *services.TransactionService
	CustomerService    *services.CustomerService
	Logger             *zap.Logger
	Validator          *validator.Validate
}

func NewTransactionHandler(transactionService *services.TransactionService, customerService *services.CustomerService, log *zap.Logger, validator *validator.Validate) *TransactionHandler {
	return &TransactionHandler{TransactionService: transactionService, CustomerService: customerService, Logger: log, Validator: validator}
}

func (h *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonResp.SendError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
	}

	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		h.Logger.Error("Error decoding transaction data "+err.Error(), zap.String("handler", "TransactionHandler"), zap.String("function", "CreateTransactionHandler"))
		JsonResp.SendError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	customer, err := h.CustomerService.CreateCustomer(&transaction.Customer)
	if err != nil {
		h.Logger.Error("Error creating customer "+err.Error(), zap.String("handler", "TransactionHandler"), zap.String("function", "CreateTransactionHandler"))
		JsonResp.SendError(w, http.StatusInternalServerError, "Error creating customer", nil)
		return
	}
	transaction.CustomerID = customer.ID

	err = h.Validator.Struct(transaction)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range err.(validator.ValidationErrors) {
				h.Logger.Error("Validation error "+validationErr.Namespace()+":"+validationErr.Tag(), zap.Any("value", validationErr.Value()), zap.Any("condition", validationErr.Param()), zap.String("handler", "TransactionHandler"), zap.String("function", "CreateTransactionHandler"))
			}
		} else {
			h.Logger.Error("Error validating transaction "+err.Error(), zap.String("handler", "TransactionHandler"), zap.String("function", "CreateTransactionHandler"))
		}
	}
	newTransaction, err := h.TransactionService.CreateTransaction(transaction)
	if err != nil {
		h.Logger.Error("Error creating transaction "+err.Error(), zap.String("handler", "TransactionHandler"), zap.String("function", "CreateTransactionHandler"))
		JsonResp.SendError(w, http.StatusInternalServerError, "Error creating transaction", nil)
		return
	}
	JsonResp.SendCreated(w, newTransaction, "Transaction created successfully")
}
