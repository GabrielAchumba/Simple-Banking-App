package services

import (
	databasePackage "github.com/GabrielAchumba/Simple-Banking-App/database/models"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/transaction/dtos"
)

type ITransactionService interface {
	CreateTransaction(createTransactionDTO dtos.CreateTransactionDTO) (interface{}, error)
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

func (impl TransactionService) CreateTransaction(createTransactionDTO dtos.CreateTransactionDTO) (interface{}, error) {

	result, err := impl.db.CreateTransaction(createTransactionDTO)
	return result, err
}

func (impl TransactionService) GetTransactions() (interface{}, error) {

	result, err := impl.db.GetTransactions()
	return result, err
}

func (impl TransactionService) GetTransaction(reference string) (interface{}, error) {

	result, err := impl.db.GetTransaction(reference)
	return result, err
}
