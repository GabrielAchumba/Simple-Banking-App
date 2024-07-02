package database

import (
	"github.com/GabrielAchumba/Simple-Banking-App/database/models"
)

var DB = &models.InMemoryDatabase{
	Accounts:     make(map[string]*models.Account),
	Transactions: make(map[string]*models.Transaction),
}
