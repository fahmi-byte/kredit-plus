package controller

import (
	"context"
	"fmt"
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
	Validate              *validator.Validate
}

func NewInstallmentProcessController(customerService service.CustomerService, transactionService service.TransactionService, paymentService service.PaymentService, paymentGatewayService service.PaymentGatewayService, installmentService service.InstallmentService, validate *validator.Validate) *InstallmentProcessControllerImpl {
	return &InstallmentProcessControllerImpl{CustomerService: customerService, TransactionService: transactionService, PaymentService: paymentService, PaymentGatewayService: paymentGatewayService, InstallmentService: installmentService, Validate: validate}
}

func (controller *InstallmentProcessControllerImpl) InstallmentProcess(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	installmentProcessRequest := web.InstallmentProcessRequest{}
	helper.ReadFromRequest(request, &installmentProcessRequest)
	err := controller.Validate.Struct(installmentProcessRequest)
	helper.PanicIfError(err)
	request.Close = true

	ctx, cancel := context.WithCancel(request.Context())
	defer cancel() // Pastikan cancel dipanggil untuk menghindari kebocoran context

	tenorValidChan := make(chan bool, 1)
	transactionChan := make(chan web.TransactionCreateRequest, 1)
	transactionResponseChan := make(chan web.TransactionResponse, 1)
	installmentCompleteChan := make(chan bool, 1)
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

	select {
	case <-tenorValidChan:
		fmt.Println("tenornya berhasil lanjutkan")
		go helper.SafeGo(
			func() {
				controller.TransactionService.Create(ctx, transactionChan, transactionResponseChan, errChan)
			})
		go helper.SafeGo(
			func() {
				controller.InstallmentService.CreateInstallment(ctx, transactionResponseChan, installmentCompleteChan, errChan)
			})
	case err := <-errChan:
		fmt.Println("masuk ke kondisi error chanel")
		helper.PanicIfError(err)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Success",
	}

	helper.WriteToReponseBody(writer, webResponse)
}
