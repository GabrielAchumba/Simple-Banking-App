package services

import (
	"github.com/GabrielAchumba/Simple-Banking-App/common/conversion"
	databasePackage "github.com/GabrielAchumba/Simple-Banking-App/database/models"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/account/dtos"
)

type IAccountService interface {
	Create(accountDTO dtos.CreateAccountDTO) (dtos.CreateAccountDTO, error)
	GetAccount(reference string) (dtos.CreateAccountDTO, error)
	SeedAdmin()
}

type AccountService struct {
	db *databasePackage.InMemoryDatabase
}

func New(_db *databasePackage.InMemoryDatabase) IAccountService {
	return &AccountService{
		db: _db,
	}
}

func (impl AccountService) SeedAdmin() {
	admin := dtos.CreateAccountDTO{
		Balance: 2000,
	}
	impl.Create(admin)

}

func (impl AccountService) Create(createAccountDTO dtos.CreateAccountDTO) (dtos.CreateAccountDTO, error) {

	result, err := impl.db.CreateAccount(createAccountDTO)
	if err != nil {
		return dtos.CreateAccountDTO{}, err
	}
	var createdAccountDTO dtos.CreateAccountDTO
	conversion.Conversion(result, &createdAccountDTO)

	return createdAccountDTO, err
}

func (impl AccountService) GetAccount(reference string) (dtos.CreateAccountDTO, error) {

	result, err := impl.db.GetAccount(reference)
	if err != nil {
		return dtos.CreateAccountDTO{}, err
	}
	var createdAccountDTO dtos.CreateAccountDTO
	conversion.Conversion(result, &createdAccountDTO)
	return createdAccountDTO, err
}
