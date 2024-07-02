package models

import (
	"github.com/GabrielAchumba/Simple-Banking-App/common/conversion"
	"github.com/GabrielAchumba/Simple-Banking-App/common/errors"
	"github.com/GabrielAchumba/Simple-Banking-App/utils"
)

type Account struct {
	Reference string  `json:"reference"`
	Balance   float64 `json:"balance"`
}

// Create
func (db *InMemoryDatabase) CreateAccount(model interface{}) (*Account, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()
	accountNumber := utils.UniqueId()

	var account Account
	err := conversion.Conversion(model, &account)
	if err != nil {
		return nil, errors.Error("wrong account model")
	}
	account.Reference = accountNumber

	db.Accounts[accountNumber] = &account
	return &account, nil
}

// Read
func (db *InMemoryDatabase) GetAccount(reference string) (*Account, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	account, exists := db.Accounts[reference]
	if !exists {
		return nil, errors.Error("account not found")
	}
	return account, nil
}

// Update
func (db *InMemoryDatabase) UpdateAccount(reference string, model interface{}) (*Account, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	var accountConvereted Transaction
	conversion.Conversion(model, &accountConvereted)

	account, exists := db.Accounts[reference]
	if !exists {
		return nil, errors.Error("account not found")
	}

	conversion.SpreadOperation(account, accountConvereted)

	return account, nil
}

// Delete
func (db *InMemoryDatabase) DeleteAccount(reference string) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	_, exists := db.Accounts[reference]
	if !exists {
		return errors.Error("account not found")
	}

	delete(db.Accounts, reference)
	return nil
}
