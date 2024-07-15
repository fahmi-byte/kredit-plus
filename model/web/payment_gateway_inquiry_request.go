package web

type PaymentGatewayInquiryRequest struct {
	MerchantCode    string `json:"merchantCode"`
	PaymentAmount   int    `json:"paymentAmount"`
	PaymentMethod   string `json:"paymentMethod"`
	MerchantOrderId string `json:"merchantOrderId"`
	ProductDetails  string `json:"productDetails"`
	Email           string `json:"email"`
	CustomerVaName  string `json:"customerVaName"`
	CallbackUrl     string `json:"callbackUrl"`
	ReturnUrl       string `json:"returnUrl"`
	Signature       string `json:"signature"`
	ExpiryPeriod    int    `json:"expiryPeriod"`
}
