package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	accountServicesPackage "github.com/GabrielAchumba/Simple-Banking-App/modules/account/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionHandler(t *testing.T) {
	_ginEngine, db := setUpRoutes()
	accountService := accountServicesPackage.New(db)
	_ginEngine.POST("/transactions", createTransactionHandler)

	account := map[string]interface{}{
		"number":  "123456789",
		"balance": 1000.0,
	}

	// Create a test account first
	account := db.CreateAccount("123456789", 1000.0)

	// Create a new HTTP request for debit transaction
	transaction := map[string]interface{}{
		"account_id": account.ID,
		"type":       "debit",
		"amount":     100.0,
	}
	body, _ := json.Marshal(transaction)
	req, _ := http.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, transaction["account_id"], int(response["account_id"].(float64)))
	assert.Equal(t, transaction["type"], response["type"])
	assert.Equal(t, transaction["amount"], response["amount"])

	// Verify the account balance is updated
	updatedAccount, err := db.GetAccount(account.ID)
	assert.NoError(t, err)
	assert.Equal(t, 900.0, updatedAccount.Balance)

	// Create a new HTTP request for credit transaction
	transaction = map[string]interface{}{
		"account_id": account.ID,
		"type":       "credit",
		"amount":     200.0,
	}
	body, _ = json.Marshal(transaction)
	req, _ = http.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Serve the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, transaction["account_id"], int(response["account_id"].(float64)))
	assert.Equal(t, transaction["type"], response["type"])
	assert.Equal(t, transaction["amount"], response["amount"])

	// Verify the account balance is updated
	updatedAccount, err = db.GetAccount(account.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1100.0, updatedAccount.Balance)
}
