package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"kredit-plus/helper"
	"kredit-plus/model/web"
	"kredit-plus/service"
	"net/http"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	Validate    *validator.Validate
}

func NewAuthController(authService service.AuthService, validate *validator.Validate) *AuthControllerImpl {
	return &AuthControllerImpl{AuthService: authService, Validate: validate}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := web.RegisterRequest{}
	helper.ReadFromRequest(request, &registerRequest)
	err := controller.Validate.Struct(registerRequest)
	helper.PanicIfError(err)

	controller.AuthService.AuthRegister(request.Context(), registerRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Create User Successfully!",
	}

	helper.WriteToReponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequest(request, &loginRequest)
	err := controller.Validate.Struct(loginRequest)
	helper.PanicIfError(err)

	token := controller.AuthService.AuthLogin(request.Context(), loginRequest.Email, loginRequest.Password)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   token,
	}

	helper.WriteToReponseBody(writer, webResponse)
}
