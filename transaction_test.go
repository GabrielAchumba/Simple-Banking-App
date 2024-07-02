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

	//accountServicesPackage "github.com/GabrielAchumba/Simple-Banking-App/modules/account/services"
	"github.com/stretchr/testify/assert"

	"log"
	"time"

	"github.com/GabrielAchumba/Simple-Banking-App/common/config"
	inMemoryDatabasePackage "github.com/GabrielAchumba/Simple-Banking-App/database"
	inMemoryDatabaseModelPackage "github.com/GabrielAchumba/Simple-Banking-App/database/models"
	accountModulePackage "github.com/GabrielAchumba/Simple-Banking-App/modules/account"
	accountServicesPackage "github.com/GabrielAchumba/Simple-Banking-App/modules/account/services"
	transactionModulePackage "github.com/GabrielAchumba/Simple-Banking-App/modules/transaction"
	transactionServicesPackage "github.com/GabrielAchumba/Simple-Banking-App/modules/transaction/services"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

func setUpTransactionRoutes() (*gin.Engine, *inMemoryDatabaseModelPackage.InMemoryDatabase) {

	_ginEngine, _ := setUpRoutes()
	gin.SetMode(gin.TestMode)

	if isProduction() != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Unable to load env file")
		}
	}

	config.Setup()
	configSettings = *config.AppSettings
	db = inMemoryDatabasePackage.DB

	_ginEngine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "*",
		RequestHeaders:  "*",
		ExposedHeaders:  "Content-Length",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	routeGroup := _ginEngine.Group("")

	accountService := accountServicesPackage.New(db)
	accountModule := accountModulePackage.New(accountService)
	accountModule.RegisterRoutes(routeGroup)

	transactionService := transactionServicesPackage.New(db)
	transactionModule := transactionModulePackage.New(transactionService)
	transactionModule.RegisterRoutes(routeGroup)

	return _ginEngine, db

}

func TestCreateTransactionHandler(t *testing.T) {
	_ginEngine, db := setUpRoutes()
	accountService := accountServicesPackage.New(db)
	//_ginEngine.POST("/transactions", createTransactionHandler)

	account := dtos.CreateAccountDTO{
		Balance: 1000.0,
	}

	// Create a test account first
	_account, err := accountService.Create(account)
	if err != nil {
		errors.Error(err.Error())
	}

	// Create a new HTTP request for credit transaction
	transaction := map[string]interface{}{
		"reference":   _account.Reference,
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
