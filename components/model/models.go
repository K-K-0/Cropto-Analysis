package model

import "time"

type Trade struct {
	ID        int       `json:"id"`
	Symbol    string    `json:"symbol"`
	Price     string    `json:"price"`
	Quantity  string    `json:"quantity"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}
