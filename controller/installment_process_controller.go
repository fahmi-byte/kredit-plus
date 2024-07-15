package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type InstallmentProcessController interface {
	InstallmentProcess(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
