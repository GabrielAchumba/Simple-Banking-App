package services

import (
	databasePackage "github.com/GabrielAchumba/Simple-Banking-App/database/models"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/account/dtos"
)

type IAccountService interface {
	Create(accountDTO dtos.CreateAccountDTO) (interface{}, error)
}

type AccountService struct {
	db *databasePackage.InMemoryDatabase
}

func New(_db *databasePackage.InMemoryDatabase) IAccountService {
	return &AccountService{
		db: _db,
	}
}

func (impl AccountService) Create(createAccountDTO dtos.CreateAccountDTO) (interface{}, error) {

	result, err := impl.db.CreateAccount(createAccountDTO)
	return result, err
}
