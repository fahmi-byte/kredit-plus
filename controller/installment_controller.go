package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type InstallmentController interface {
	FindInstallmentById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindInstallmentCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreatePaymentInstallment(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
