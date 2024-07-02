package models

import (
	"time"

	"github.com/GabrielAchumba/Simple-Banking-App/common/conversion"
	"github.com/GabrielAchumba/Simple-Banking-App/common/errors"
	"github.com/GabrielAchumba/Simple-Banking-App/utils"
)

type Transaction struct {
	ID              string
	Amount          float64
	Type            string
	TransactionDate string
	Description     string
	Reference       string
}

// Create
func (db *InMemoryDatabase) CreateTransaction(model interface{}) (*Transaction, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()
	uid := utils.UniqueId()

	var transaction Transaction
	conversion.Conversion(model, &transaction)

	account, exists := db.Accounts[transaction.Reference]
	if !exists {
		return nil, errors.Error("account not found")
	}

	if transaction.Type == "debit" {
		account.Balance -= transaction.Amount
	} else if transaction.Type == "credit" {
		account.Balance += transaction.Amount
	} else {
		return nil, errors.Error("invalid transaction type")
	}

	transaction.ID = uid
	transaction.TransactionDate = time.Now().Format(time.RFC3339)

	db.Transactions[uid] = &transaction
	db.Accounts[transaction.Reference] = account
	return &transaction, nil
}

// Read One
func (db *InMemoryDatabase) GetTransaction(id string) (*Transaction, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	transaction, exists := db.Transactions[id]
	if !exists {
		return nil, errors.Error("transaction not found")
	}
	return transaction, nil
}

// Read All
func (db *InMemoryDatabase) GetTransactions() (map[string]*Transaction, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	return db.Transactions, nil
}

// Update
/* func (db *InMemoryDatabase) UpdateTransaction(id string, model interface{}) (*Transaction, error) {

	db.Mu.Lock()
	defer db.Mu.Unlock()

	var transactionConvereted Transaction
	conversion.Conversion(model, &transactionConvereted)

	transaction, exists := db.Transactions[id]
	if !exists {
		return nil, errors.Error("transaction not found")
	}

	conversion.SpreadOperation(transaction, transactionConvereted)

	return transaction, nil
} */

// Delete
func (db *InMemoryDatabase) DeleteTransaction(reference string) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	_, exists := db.Accounts[reference]
	if !exists {
		return errors.Error("account not found")
	}

	delete(db.Accounts, reference)
	return nil
}
