package handler

import (
	"encoding/json"
	"expenses/internal/models"
	"expenses/internal/service"
	"net/http"
	"time"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) LogIncome(w http.ResponseWriter, r *http.Request) {
	var req models.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	tx, err := h.service.LogIncome(req.Amount, req.Currency, req.Category, date)
	if err != nil {
		http.Error(w, "Failed to log income", http.StatusInternalServerError)
		return
	}

	res := models.TransactionResponse{
		Message:       "Income logged successfully",
		TransactionId: tx.TransactionID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TransactionHandler) LogExpense(w http.ResponseWriter, r *http.Request) {
	var req models.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	tx, err := h.service.LogExpense(req.Amount, req.Currency, req.Category, date)
	if err != nil {
		http.Error(w, "Failed to log expense", http.StatusInternalServerError)
		return
	}

	res := models.TransactionResponse{
		Message:       "Expense logged successfully",
		TransactionId: tx.TransactionID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.GetTransactions()
	if err != nil {
		http.Error(w, "Failed to get transactions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}