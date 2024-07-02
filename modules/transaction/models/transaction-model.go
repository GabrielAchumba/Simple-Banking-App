package models

import (
	"time"
)

type TransactionModel struct {
	ID              int       `json:"id"`
	AccountID       int       `json:"accountID"`
	Amount          float64   `json:"amount"`
	Type            string    `json:"type"`
	TransactionDate time.Time `json:"transactionDate"`
	Description     string    `json:"description"`
	Reference       string    `json:"reference"`
}
