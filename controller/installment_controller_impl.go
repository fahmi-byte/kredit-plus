package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"kredit-plus/constants"
	"kredit-plus/helper"
	"kredit-plus/model/web"
	"kredit-plus/service"
	"net/http"
	"strconv"
)

type InstallmentControllerImpl struct {
	InstallmentService service.InstallmentService
	Validate           *validator.Validate
}

func NewInstallmentController(installmentService service.InstallmentService, validate *validator.Validate) *InstallmentControllerImpl {
	return &InstallmentControllerImpl{InstallmentService: installmentService, Validate: validate}
}

func (controller *InstallmentControllerImpl) FindInstallmentById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	installmentProcessRequest := web.InstallmentProcessRequest{}
	installmentParam := params.ByName("installmentId")
	helper.ReadFromRequest(request, &installmentProcessRequest)
	err := controller.Validate.Struct(installmentProcessRequest)
	helper.PanicIfError(err)
	installmentId, err := strconv.Atoi(installmentParam)
	helper.PanicIfError(err)

	installmentResponse := controller.InstallmentService.FindOne(request.Context(), installmentId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   installmentResponse,
	}

	helper.WriteToReponseBody(writer, webResponse)
}

func (controller *InstallmentControllerImpl) FindInstallmentCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerParam := params.ByName("customerId")
	customerId, err := strconv.Atoi(customerParam)
	helper.PanicIfError(err)

	installmentResponse := controller.InstallmentService.FindAllInstallmentCustomer(request.Context(), customerId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   installmentResponse,
	}

	helper.WriteToReponseBody(writer, webResponse)
}

func (controller *InstallmentControllerImpl) CreatePaymentInstallment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	installmentRequest := web.InstallmentRequest{}
	helper.ReadFromRequest(request, &installmentRequest)
	err := controller.Validate.Struct(installmentRequest)
	helper.PanicIfError(err)

	installment := controller.InstallmentService.GetInstallmentForPayment(request.Context(), installmentRequest.InstallmentId)
	merchantOrderId := fmt.Sprintf("%s/%d", installment.ContractNumber, installment.Month)
	signature := helper.GenerateMD5Hash(constants.MerchantCode, merchantOrderId, strconv.Itoa(int(installment.TotalPayment)), constants.ApiKey)

	claims := request.Context().Value("claims").(*helper.Claims)
	inquiryRequest := web.PaymentGatewayInquiryRequest{
		MerchantCode:    constants.MerchantCode,
		PaymentAmount:   int(installment.TotalPayment),
		PaymentMethod:   constants.PaymentMethod,
		MerchantOrderId: merchantOrderId,
		ProductDetails:  "Pembayaran cicilan Kredit Plus",
		Email:           claims.Email,
		CustomerVaName:  claims.Name,
		CallbackUrl:     constants.CallbackUrl,
		ReturnUrl:       "http://www.contoh.com/return",
		Signature:       signature,
		ExpiryPeriod:    10,
	}

	requestBody, err := json.Marshal(inquiryRequest)
	if err != nil {
		http.Error(writer, "Failed to create request body", http.StatusInternalServerError)
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, constants.ApiUrlPaymentGateway, bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(writer, "Failed to create request", http.StatusInternalServerError)
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		http.Error(writer, "Request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Baca respons body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(writer, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Parse respons JSON
	var inquiryResp web.InquiryResponse
	if err := json.Unmarshal(respBody, &inquiryResp); err != nil {
		http.Error(writer, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   inquiryResp,
	}

	helper.WriteToReponseBody(writer, webResponse)
}
