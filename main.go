package main

import (
	"log"
	"net/http"
	"os"
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

var (
	configSettings config.Settings
	ginEngine      *gin.Engine
	db             *inMemoryDatabaseModelPackage.InMemoryDatabase
)

func isProduction() string {
	appEnv := os.Getenv("APP_ENV")
	return appEnv
}

func init() {

	gin.SetMode(gin.ReleaseMode)
	ginEngine = gin.Default()

	if isProduction() != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Unable to load env file")
		}
	}

	config.Setup()
	configSettings = *config.AppSettings
	db = inMemoryDatabasePackage.DB

}

func SetUpModules() {

	ginEngine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "*",
		RequestHeaders:  "*",
		ExposedHeaders:  "Content-Length",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	routeGroup := ginEngine.Group("")

	accountService := accountServicesPackage.New(db)
	accountModule := accountModulePackage.New(accountService)
	accountModule.RegisterRoutes(routeGroup)

	transactionService := transactionServicesPackage.New(db)
	transactionModule := transactionModulePackage.New(transactionService)
	transactionModule.RegisterRoutes(routeGroup, "/transaction")

}

func main() {

	SetUpModules()

	port := configSettings.Server.Port

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      ginEngine,
		ReadTimeout:  time.Second * 600,
		WriteTimeout: time.Second * 1200,
	}

	log.Println("Banking App APIs running on port: " + port)

	log.Fatal(server.ListenAndServe())
}
