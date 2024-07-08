package controller

import (
	"github.com/julienschmidt/httprouter"
	"kredit-plus/helper"
	"kredit-plus/model/web"
	"kredit-plus/service"
	"net/http"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{TransactionService: transactionService}
}

func (controller *TransactionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionResponse := controller.TransactionService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToReponseBody(writer, webResponse)
}
