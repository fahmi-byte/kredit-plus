package main

import (
	"github.com/go-playground/validator/v10"
	"kredit-plus/app"
	"kredit-plus/controller"
	"kredit-plus/helper"
	"kredit-plus/middleware"
	"kredit-plus/model/repository"
	"kredit-plus/service"
	"net/http"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	authRepositoryImpl := repository.NewAuthRepository()
	installmentRepositoryImpl := repository.NewInstallmentRepository()
	customerRepositoryImpl := repository.NewCustomerRepository()
	transactionRepositoryImpl := repository.NewTransactionRepository()
	merchantRepositoryImpl := repository.NewMerchantRepository()
	paymentRepositoryImpl := repository.NewPaymentRepository()
	paymentGatewayRepositoryImpl := repository.NewPaymentGatewayRepository()

	db := app.NewDB()
	validate := NewValidator()
	authServiceImpl := service.NewAuthService(authRepositoryImpl, db)
	installmentServiceImpl := service.NewInstallmentService(installmentRepositoryImpl, db)
	customerServiceImpl := service.NewCustomerService(customerRepositoryImpl, db)
	transactionServiceImpl := service.NewTransactionService(transactionRepositoryImpl, merchantRepositoryImpl, db)
	paymentServiceImpl := service.NewPaymentService(paymentRepositoryImpl, db)
	//merchantServiceImpl := service.NewMerchantService(merchantRepositoryImpl, db)
	paymentGatewayServiceImpl := service.NewPaymentGatewayService(paymentGatewayRepositoryImpl, installmentRepositoryImpl, db)

	authControllerImpl := controller.NewAuthController(authServiceImpl, validate)
	customerControllerImpl := controller.NewCustomerController(customerServiceImpl, validate)
	installmentControllerImpl := controller.NewInstallmentController(installmentServiceImpl, validate)
	installmentProcessControllerImpl := controller.NewInstallmentProcessController(customerServiceImpl, transactionServiceImpl, paymentServiceImpl, paymentGatewayServiceImpl, installmentServiceImpl, validate)
	paymentGatewayController := controller.NewPaymentGatewayController(paymentGatewayServiceImpl)
	transactionControllerImpl := controller.NewTransactionController(transactionServiceImpl)
	router := app.NewRouter(installmentProcessControllerImpl, transactionControllerImpl, paymentGatewayController, authControllerImpl, installmentControllerImpl, customerControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router, db)
	server := NewServer(authMiddleware)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
