package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CustomerController interface {
	CreateCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
