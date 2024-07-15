package app

import (
	"github.com/julienschmidt/httprouter"
	"kredit-plus/controller"
	"kredit-plus/exception"
)

func NewRouter(installmentProcessController controller.InstallmentProcessController, transactionController controller.TransactionController, paymentGatewayController controller.PaymentGatewayController, authController controller.AuthController, installmentController controller.InstallmentController, customerController controller.CustomerController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/login", authController.Login)
	router.GET("/api/transactions", transactionController.FindAll)
	router.GET("/api/installment-customer/:customerId", installmentController.FindInstallmentCustomer)
	router.GET("/api/installment/:installmentId", installmentController.FindInstallmentById)
	router.POST("/api/installment/payment", installmentController.CreatePaymentInstallment)
	router.POST("/api/customer", customerController.CreateCustomer)
	router.POST("/api/installment-process", installmentProcessController.InstallmentProcess)
	router.POST("/api/payment-gateway-callback", paymentGatewayController.CallbackPayment)
	router.PanicHandler = exception.ErrorHandler

	return router
}
