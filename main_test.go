package main

import (
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

func setUpRoutes() (*gin.Engine, *inMemoryDatabaseModelPackage.InMemoryDatabase) {

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

	transactionService := transactionServicesPackage.New(db)
	transactionModule := transactionModulePackage.New(transactionService)
	transactionModule.RegisterRoutes(routeGroup, "/transaction")

	return _ginEngine, db

}

/*
func TestCreateAccountHandler(t *testing.T) {
	// Initialize the router and database
	r := gin.Default()
	db = NewInMemoryDatabase()
	r.POST("/accounts", createAccountHandler)

	// Create a new HTTP request
	account := map[string]interface{}{
		"number":  "123456789",
		"balance": 1000.0,
	}
	body, _ := json.Marshal(account)
	req, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(body))
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
	assert.Equal(t, account["number"], response["number"])
	assert.Equal(t, account["balance"], response["balance"])
}

*/
