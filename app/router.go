package app

import (
	"github.com/julienschmidt/httprouter"
	"kredit-plus/controller"
	"kredit-plus/exception"
)

func NewRouter(installmentController controller.InstallmentController, transactionController controller.TransactionController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/transactions", transactionController.FindAll)
	router.POST("/api/installment-process", installmentController.InstallmentProcess)
	router.PanicHandler = exception.ErrorHandler

	return router
}
