package web

type TransactionCreateRequest struct {
	CustomerId        int
	MerchantId        int
	OTR               float32
	AdminFee          float32
	InstallmentAmount float32
	InterestAmount    float32
	InterestRateId    int
	AssetName         string
	Tenor             string
}
