package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabrielAchumba/Simple-Banking-App/common/errors"
	"github.com/GabrielAchumba/Simple-Banking-App/constants"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/account/dtos"
	accountServicesPackage "github.com/GabrielAchumba/Simple-Banking-App/modules/account/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionHandler(t *testing.T) {
	_ginEngine, db := setUpRoutes()
	accountService := accountServicesPackage.New(db)
	//_ginEngine.POST("/transactions", createTransactionHandler)

	account := dtos.CreateAccountDTO{
		Number:  "",
		Balance: 1000.0,
	}

	// Create a test account first
	_account, err := accountService.Create(account)
	if err != nil {
		errors.Error(err.Error())
	}

	// Create a new HTTP request for credit transaction
	transaction := map[string]interface{}{
		"reference":   _account.Number,
		"type":        "credit",
		"amount":      300.0,
		"description": "Great to credit my account",
	}

	body, _ := json.Marshal(transaction)
	path := "/" + constants.TransactionControllerName + "/" + constants.Payments
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	_ginEngine.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)
	assert.Equal(t, transaction["reference"], int(response["reference"].(float64)))
	assert.Equal(t, transaction["type"], response["type"])
	assert.Equal(t, transaction["amount"], response["amount"])

}
