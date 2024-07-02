package config

import (
	"fmt"
	"os"
)

type Settings struct {
	Database struct {
		DatabaseNme        string
		DatebaseConnection string
	}

	Tables struct {
		Arithmetics string
	}

	Server struct {
		Port string
	}
}

var AppSettings = &Settings{}

func Setup() {

	AppSettings.Database.DatabaseNme = os.Getenv("DATABASENAME")
	AppSettings.Database.DatebaseConnection = os.Getenv("DATABASECONNECTION")

	AppSettings.Tables.Arithmetics = os.Getenv("ACCOUNT")
	AppSettings.Tables.Arithmetics = os.Getenv("TRANSACTION")

	AppSettings.Server.Port = os.Getenv("PORT")

	fmt.Println("Appsetings loaded successfully")
}
