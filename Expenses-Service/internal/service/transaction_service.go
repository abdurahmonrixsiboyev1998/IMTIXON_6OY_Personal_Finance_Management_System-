package service

import (
	"expenses/internal/models"
	"expenses/internal/repository"
	"time"

	"github.com/google/uuid"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) LogIncome(amount float64, currency, category string, date time.Time) (*models.Transaction, error) {
	tx := &models.Transaction{
		TransactionID: uuid.New().String(),
		Amount:        amount,
		Currency:      currency,
		Category:      category,
		Date:          date,
		Type:          "income",
	}
	err := s.repo.CreateTransaction(tx)
	return tx, err
}

func (s *TransactionService) LogExpense(amount float64, currency, category string, date time.Time) (*models.Transaction, error) {
	tx := &models.Transaction{
		TransactionID: uuid.New().String(),
		Amount:        amount,
		Currency:      currency,
		Category:      category,
		Date:          date,
		Type:          "expense",
	}
	err := s.repo.CreateTransaction(tx)
	return tx, err
}

func (s *TransactionService) GetTransactions() ([]models.Transaction, error) {
	return s.repo.GetTransactions()
}