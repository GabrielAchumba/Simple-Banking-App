package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabrielAchumba/Simple-Banking-App/constants"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountHandler(t *testing.T) {

	_ginEngine, _ := setUpRoutes()
	// Create a new HTTP request
	path := "/" + constants.AccountControllerName + "/" + constants.CreateAccount
	account := map[string]interface{}{
		"number":  "123456789",
		"balance": 1000.0,
	}
	body, _ := json.Marshal(account)
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	_ginEngine.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, account["number"], response["number"])
	assert.Equal(t, account["balance"], response["balance"])
}
