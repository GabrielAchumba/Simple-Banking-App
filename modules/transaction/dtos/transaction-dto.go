package dtos

import (
	"time"
)

type CreateTransactionDTO struct {
	Reference       string    `json:"reference"`
	Amount          float64   `json:"amount"`
	Type            string    `json:"type"`
	TransactionDate time.Time `json:"transactionDate"`
	Description     string    `json:"description"`
}
