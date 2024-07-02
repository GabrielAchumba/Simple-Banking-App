package config

import (
	"fmt"
	"os"
)

type Settings struct {
	Database struct {
		DatabaseNme        string
		DatebaseConnection string
		InMemoryDatabase   struct {
			ArithmeticsTable map[int]interface{}
		}
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

	AppSettings.Tables.Arithmetics = os.Getenv("ARITHMETICS")

	AppSettings.Server.Port = os.Getenv("PORT")

	AppSettings.Database.InMemoryDatabase.ArithmeticsTable = make(map[int]interface{})

	fmt.Println("Appsetings loaded successfully")
}
