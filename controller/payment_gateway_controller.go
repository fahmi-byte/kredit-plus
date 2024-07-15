package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PaymentGatewayController interface {
	CallbackPayment(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
