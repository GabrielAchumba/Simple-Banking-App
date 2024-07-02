package services

import (
	"github.com/GabrielAchumba/Simple-Banking-App/common/conversion"
	databasePackage "github.com/GabrielAchumba/Simple-Banking-App/database/models"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/transaction/dtos"
)

type ITransactionService interface {
	CreateTransaction(createTransactionDTO dtos.CreateTransactionDTO) (dtos.CreateTransactionDTO, error)
	GetTransactions() (interface{}, error)
	GetTransaction(reference string) (interface{}, error)
}

type TransactionService struct {
	db *databasePackage.InMemoryDatabase
}

func New(_db *databasePackage.InMemoryDatabase) ITransactionService {
	return &TransactionService{
		db: _db,
	}
}

func (impl TransactionService) CreateTransaction(createTransactionDTO dtos.CreateTransactionDTO) (dtos.CreateTransactionDTO, error) {

	result, err := impl.db.CreateTransaction(createTransactionDTO)
	if err != nil {
		return dtos.CreateTransactionDTO{}, err
	}
	var createdTransactionDTO dtos.CreateTransactionDTO
	conversion.Conversion(result, &createdTransactionDTO)
	return createdTransactionDTO, err
}

func (impl TransactionService) GetTransactions() (interface{}, error) {

	result, err := impl.db.GetTransactions()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (impl TransactionService) GetTransaction(reference string) (interface{}, error) {

	result, err := impl.db.GetTransaction(reference)
	if err != nil {
		return nil, err
	}
	return result, err
}
