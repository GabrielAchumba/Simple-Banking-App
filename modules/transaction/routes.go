package arithmetics

import (
	"github.com/GabrielAchumba/Simple-Banking-App/common/rest"
	"github.com/GabrielAchumba/Simple-Banking-App/constants"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/transaction/controllers"
	"github.com/GabrielAchumba/Simple-Banking-App/modules/transaction/services"
	"github.com/gin-gonic/gin"
)

type TransactionModule struct {
	transactionController controllers.ITransactionController
}

func New(_transactionController services.ITransactionService) *TransactionModule {
	transactionModule := new(TransactionModule)
	transactionModule.transactionController = controllers.New(_transactionController)
	return transactionModule
}

func (transactionModule *TransactionModule) RegisterRoutes(routeGroup *gin.RouterGroup, controllerName string) {

	moduleRoutes := routeGroup.Group("/" + constants.TransactionControllerName)

	serveHTTP := rest.ServeHTTP

	moduleRoutes.POST("/"+constants.Payments, serveHTTP(transactionModule.transactionController.CreateTransaction))
	moduleRoutes.GET("/"+constants.PaymentsReference, serveHTTP(transactionModule.transactionController.GetTransaction))
	moduleRoutes.GET("/"+constants.Transactions, serveHTTP(transactionModule.transactionController.GetTransactions))
}
