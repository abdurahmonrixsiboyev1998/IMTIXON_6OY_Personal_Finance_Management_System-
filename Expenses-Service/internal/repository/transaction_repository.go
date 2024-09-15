package repository

import (
	"database/sql"
	"expenses/internal/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(tx *models.Transaction) error {
	query := `INSERT INTO transactions (transaction_id, amount, currency, category, date, type)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, tx.TransactionID, tx.Amount, tx.Currency, tx.Category, tx.Date, tx.Type)
	return err
}

func (r *TransactionRepository) GetTransactions() ([]models.Transaction, error) {
	query := `SELECT transaction_id, amount, currency, category, date, type FROM transactions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		if err := rows.Scan(&tx.TransactionID, &tx.Amount, &tx.Currency, &tx.Category,
			&tx.Date, &tx.Type); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}
