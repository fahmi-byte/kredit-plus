package constants

type OperatorType string

const (
	Increment OperatorType = "increment"
	Decrement OperatorType = "decrement"
)

const JWTSecret = "secret_key"

const MerchantCode = "DS19645"
const PaymentMethod = "M2"
const CallbackUrl = "https://ffcb-103-154-109-53.ngrok-free.app/api/payment-gateway-callback"

const ApiKey = "f0e4719f20e18860bf2c1ae3bbffd6c6"
const ApiUrlPaymentGateway = "https://sandbox.duitku.com/webapi/api/merchant/v2/inquiry"
