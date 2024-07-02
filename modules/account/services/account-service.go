package services

import (
	"github.com/GabrielAchumba/Simple-Banking-App/common/conversion"
	databasePackage "github.com/GabrielAchumba/Simple-Banking-App/database/models"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/account/dtos"
)

type IAccountService interface {
	Create(accountDTO dtos.CreateAccountDTO) (dtos.CreateAccountDTO, error)
}

type AccountService struct {
	db *databasePackage.InMemoryDatabase
}

func New(_db *databasePackage.InMemoryDatabase) IAccountService {
	return &AccountService{
		db: _db,
	}
}

func (impl AccountService) Create(createAccountDTO dtos.CreateAccountDTO) (dtos.CreateAccountDTO, error) {

	result, err := impl.db.CreateAccount(createAccountDTO)
	var createdAccountDTO dtos.CreateAccountDTO
	conversion.SpreadOperation(createdAccountDTO, result)

	return createdAccountDTO, err
}
