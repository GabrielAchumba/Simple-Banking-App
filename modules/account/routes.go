package arithmetics

import (
	"github.com/GabrielAchumba/Simple-Banking-App/common/rest"
	"github.com/GabrielAchumba/Simple-Banking-App/constants"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/account/controllers"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/account/services"
	"github.com/gin-gonic/gin"
)

type AccountModule struct {
	accountController controllers.IAccountController
}

func New(_accountService services.IAccountService) *AccountModule {
	accountModule := new(AccountModule)
	accountModule.accountController = controllers.New(_accountService)
	_accountService.SeedAdmin()
	return accountModule
}

func (accountModule *AccountModule) RegisterRoutes(routeGroup *gin.RouterGroup) {

	moduleRoutes := routeGroup.Group("/" + constants.AccountControllerName)

	serveHTTP := rest.ServeHTTP

	moduleRoutes.POST("/"+constants.CreateAccount, serveHTTP(accountModule.accountController.Create))
	moduleRoutes.GET("/"+constants.GetAccount, serveHTTP(accountModule.accountController.GetAccount))
}
