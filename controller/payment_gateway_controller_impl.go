package controller

import (
	"github.com/julienschmidt/httprouter"
	"kredit-plus/helper"
	"kredit-plus/model/web"
	"kredit-plus/service"
	"net/http"
)

type PaymentGatewayControllerImpl struct {
	PaymentGatewayService service.PaymentGatewayService
}

func NewPaymentGatewayController(paymentGatewayService service.PaymentGatewayService) *PaymentGatewayControllerImpl {
	return &PaymentGatewayControllerImpl{PaymentGatewayService: paymentGatewayService}
}

func (controller *PaymentGatewayControllerImpl) CallbackPayment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaction := request.FormValue("merchantOrderId")
	controller.PaymentGatewayService.CallbackPayment(request.Context(), transaction)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Success",
	}

	helper.WriteToReponseBody(writer, webResponse)
}
