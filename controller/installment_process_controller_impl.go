package controller

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"kredit-plus/helper"
	"kredit-plus/model/web"
	"kredit-plus/service"
	"net/http"
)

type InstallmentProcessControllerImpl struct {
	CustomerService       service.CustomerService
	TransactionService    service.TransactionService
	PaymentService        service.PaymentService
	PaymentGatewayService service.PaymentGatewayService
	InstallmentService    service.InstallmentService
	MerchantService       service.MerchantService
	Validate              *validator.Validate
}

func NewInstallmentProcessController(customerService service.CustomerService, transactionService service.TransactionService, paymentService service.PaymentService, paymentGatewayService service.PaymentGatewayService, installmentService service.InstallmentService, merchantService service.MerchantService, validate *validator.Validate) *InstallmentProcessControllerImpl {
	return &InstallmentProcessControllerImpl{CustomerService: customerService, TransactionService: transactionService, PaymentService: paymentService, PaymentGatewayService: paymentGatewayService, InstallmentService: installmentService, MerchantService: merchantService, Validate: validate}
}

func (controller *InstallmentProcessControllerImpl) InstallmentProcess(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	installmentProcessRequest := web.InstallmentProcessRequest{}
	helper.ReadFromRequest(request, &installmentProcessRequest)
	err := controller.Validate.Struct(installmentProcessRequest)
	helper.PanicIfError(err)
	merchant := controller.MerchantService.FindWithId(request.Context(), installmentProcessRequest.MerchantId)

	ctx, cancel := context.WithCancel(request.Context())
	defer cancel() // Pastikan cancel dipanggil untuk menghindari kebocoran context

	tenorValidChan := make(chan bool, 1)
	transactionValidChan := make(chan bool, 1)
	paymentCompleteChan := make(chan bool, 1)
	transactionChan := make(chan web.TransactionCreateRequest, 1)
	transactionResponseChan := make(chan web.TransactionResponse, 1)
	paymentSuccessChan := make(chan bool, 1)
	errChan := make(chan error, 1)

	transactionChan <- web.TransactionCreateRequest{
		CustomerId:     installmentProcessRequest.CustomerId,
		MerchantId:     installmentProcessRequest.MerchantId,
		OTR:            installmentProcessRequest.OTR,
		AdminFee:       installmentProcessRequest.AdminFee,
		AssetName:      installmentProcessRequest.AssetName,
		InterestRateId: installmentProcessRequest.InterestRateId,
		Tenor:          installmentProcessRequest.Tenor,
	}

	go helper.SafeGo(
		func() {
			controller.CustomerService.CheckTenorLimitCustomer(ctx, installmentProcessRequest.CustomerId, installmentProcessRequest.Tenor, installmentProcessRequest.OTR, installmentProcessRequest.Pin, tenorValidChan, errChan)
		})

	// Tunggu sampai tenorValidChan mengembalikan nilai
	select {
	case isValid := <-tenorValidChan:
		if isValid {
			go helper.SafeGo(
				func() {
					controller.TransactionService.Create(ctx, transactionChan, transactionResponseChan, transactionValidChan, errChan)
				})
			go helper.SafeGo(
				func() {
					controller.PaymentGatewayService.PaymentProcess(ctx, merchant.BankAccount, installmentProcessRequest.OTR, paymentSuccessChan, errChan)
				})
			go helper.SafeGo(
				func() {
					controller.PaymentService.CreatePayment(ctx, transactionResponseChan, errChan)
					controller.InstallmentService.CreateInstallment(ctx, transactionResponseChan, paymentCompleteChan, errChan)
				})
		} else {
			errChan <- errors.New("limit tenor is not enough")
		}
	case err := <-errChan:
		helper.PanicIfError(err)
		return
	}

	// Tunggu sampai operasi selesai atau terjadi error
	select {
	case err := <-errChan:
		helper.PanicIfError(err)
		return
	case <-paymentCompleteChan:
		transactionResponse := controller.TransactionService.FindAll(ctx)
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   transactionResponse[0],
		}

		helper.WriteToReponseBody(writer, webResponse)
	}

}
