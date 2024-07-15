package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"kredit-plus/helper"
	"kredit-plus/model/web"
	"kredit-plus/service"
	"net/http"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
	Validate        *validator.Validate
}

func NewCustomerController(customerService service.CustomerService, validate *validator.Validate) *CustomerControllerImpl {
	return &CustomerControllerImpl{CustomerService: customerService, Validate: validate}
}

func (controller *CustomerControllerImpl) CreateCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerRequest := web.CustomerRequest{}
	helper.ReadFromRequest(request, &customerRequest)
	err := controller.Validate.Struct(customerRequest)
	helper.PanicIfError(err)

	controller.CustomerService.CreateCustomer(request.Context(), customerRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Success Create Customer!",
	}

	helper.WriteToReponseBody(writer, webResponse)
}
