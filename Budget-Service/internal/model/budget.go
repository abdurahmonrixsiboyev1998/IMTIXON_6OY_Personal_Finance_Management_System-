package model

import "time"

type Budget struct {
	ID        string    `json:"id" bson:"_id"`
	Category  string    `json:"category" bson:"category"`
	Amount    float64   `json:"amount" bson:"amount"`
	Spent     float64   `json:"spent" bson:"spent"`
	Currency  string    `json:"currency" bson:"currency"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}