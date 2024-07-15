package web

type InquiryResponse struct {
	MerchantCode  string `json:"merchantCode"`
	Reference     string `json:"reference"`
	PaymentUrl    string `json:"paymentUrl"`
	VaNumber      string `json:"vaNumber"`
	Amount        string `json:"amount"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}
