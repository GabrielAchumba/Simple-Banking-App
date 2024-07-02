package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabrielAchumba/Simple-Banking-App/constants"
	"github.com/stretchr/testify/assert"

	"log"
	"time"

	"github.com/GabrielAchumba/Simple-Banking-App/common/config"
	"github.com/GabrielAchumba/Simple-Banking-App/common/conversion"
	"github.com/GabrielAchumba/Simple-Banking-App/common/rest"
	inMemoryDatabasePackage "github.com/GabrielAchumba/Simple-Banking-App/database"
	inMemoryDatabaseModelPackage "github.com/GabrielAchumba/Simple-Banking-App/database/models"
	accountModulePackage "github.com/GabrielAchumba/Simple-Banking-App/modules/account"
	accountServicesPackage "github.com/GabrielAchumba/Simple-Banking-App/modules/account/services"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

func setUpAccountRoutes() (*gin.Engine, *inMemoryDatabaseModelPackage.InMemoryDatabase) {

	gin.SetMode(gin.TestMode)
	_ginEngine := gin.Default()

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

	/* transactionService := transactionServicesPackage.New(db)
	transactionModule := transactionModulePackage.New(transactionService)
	transactionModule.RegisterRoutes(routeGroup) */

	return _ginEngine, db

}

func TestCreateAccountHandler(t *testing.T) {

	_ginEngine, _ := setUpAccountRoutes()
	// Create a new HTTP request
	path := "/" + constants.AccountControllerName + "/" + constants.CreateAccount
	account := map[string]interface{}{
		"balance": 1000.0,
	}

	body, _ := json.Marshal(account)
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	_ginEngine.ServeHTTP(w, req)

	// Debug print statements
	t.Log("Request Body: ", string(body))
	t.Log("Response Code: ", w.Code)
	t.Log("Response Body: ", w.Body.String())

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	var convertedResponse rest.Response
	conversion.Conversion(response, &convertedResponse)
	var data1 interface{} = convertedResponse.Data
	data2, err := conversion.ConvertInterfaceToMap(data1)
	assert.NoError(t, err)
	//var data map[string]interface{} = convertedResponse.Data
	assert.Equal(t, account["balance"], data2["balance"])
}
