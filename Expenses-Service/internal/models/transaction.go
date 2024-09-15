package models

import "time"

type Transaction struct {
	TransactionID string    `json:"transactionId"`
	Type          string    `json:"type"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Category      string    `json:"category"`
	Date          time.Time `json:"date"`
}

type TransactionRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Category string  `json:"category"`
	Date     string  `json:"date"`
}

type TransactionResponse struct {
	Message       string `json:"message"`
	TransactionId string `json:"transactionId"`
}
