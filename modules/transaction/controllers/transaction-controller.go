package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/Simple-Banking-App/common/rest"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/transaction/dtos"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/transaction/services"
	"github.com/gin-gonic/gin"
)

type ITransactionController interface {
	CreateTransaction(ctx *gin.Context) *rest.Response
	GetTransactions(ctx *gin.Context) *rest.Response
	GetTransaction(ctx *gin.Context) *rest.Response
}

type TransactionController struct {
	transactionService services.ITransactionService
}

var _response rest.Response

func New(_transactionService services.ITransactionService) ITransactionController {
	return &TransactionController{
		transactionService: _transactionService,
	}
}

func (ctrl *TransactionController) CreateTransaction(ctx *gin.Context) *rest.Response {
	var createTransactionDTO dtos.CreateTransactionDTO

	err := ctx.BindJSON(createTransactionDTO)
	if err != nil {
		return _response.GetError(http.StatusBadRequest, err.Error())
	}

	result, err := ctrl.transactionService.CreateTransaction(createTransactionDTO)
	if err != nil {
		return _response.GetError(http.StatusBadRequest, err.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, result)
	}
}

func (ctrl *TransactionController) GetTransactions(ctx *gin.Context) *rest.Response {

	result, err := ctrl.transactionService.GetTransactions()
	if err != nil {
		return _response.GetError(http.StatusBadRequest, err.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, result)
	}
}

func (ctrl *TransactionController) GetTransaction(ctx *gin.Context) *rest.Response {

	reference := ctx.Param("reference")
	result, err := ctrl.transactionService.GetTransaction(reference)
	if err != nil {
		return _response.GetError(http.StatusBadRequest, err.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, result)
	}
}
