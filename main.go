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

	installmentRepositoryImpl := repository.NewInstallmentRepository()
	customerRepositoryImpl := repository.NewCustomerRepository()
	transactionRepositoryImpl := repository.NewTransactionRepository()
	merchantRepositoryImpl := repository.NewMerchantRepository()
	paymentRepositoryImpl := repository.NewPaymentRepository()
	paymentGatewayRepositoryImpl := repository.NewPaymentGatewayRepository()

	db := app.NewDB()
	validate := NewValidator()
	installmentServiceImpl := service.NewInstallmentService(installmentRepositoryImpl, db)
	customerServiceImpl := service.NewCustomerService(customerRepositoryImpl, db)
	transactionServiceImpl := service.NewTransactionService(transactionRepositoryImpl, db)
	paymentServiceImpl := service.NewPaymentService(paymentRepositoryImpl, db)
	merchantServiceImpl := service.NewMerchantService(merchantRepositoryImpl, db)
	paymentGatewayServiceImpl := service.NewPaymentGatewayService(paymentGatewayRepositoryImpl)

	installmentControllerImpl := controller.NewInstallmentProcessController(customerServiceImpl, transactionServiceImpl, paymentServiceImpl, paymentGatewayServiceImpl, installmentServiceImpl, merchantServiceImpl, validate)
	transactionControllerImpl := controller.NewTransactionController(transactionServiceImpl)
	router := app.NewRouter(installmentControllerImpl, transactionControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router, db)
	server := NewServer(authMiddleware)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
