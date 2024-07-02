package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/Simple-Banking-App/common/rest"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/account/dtos"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/account/services"
	"github.com/gin-gonic/gin"
)

type IAccountController interface {
	Create(ctx *gin.Context) *rest.Response
}

type AccountController struct {
	accountService services.IAccountService
}

var _response rest.Response

func New(_accountService services.IAccountService) IAccountController {
	return &AccountController{
		accountService: _accountService,
	}
}

func (ctrl *AccountController) Create(ctx *gin.Context) *rest.Response {
	var createAccountDTO dtos.CreateAccountDTO

	err := ctx.BindJSON(createAccountDTO)
	if err != nil {
		return _response.GetError(http.StatusBadRequest, err.Error())
	}

	result, err := ctrl.accountService.Create(createAccountDTO)
	if err != nil {
		return _response.GetError(http.StatusBadRequest, err.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, result)
	}
}
