package domain

import "time"

type Transaction struct {
	Id                int
	CustomerId        int
	FullName          string
	ContractNumber    string
	MerchantId        int
	MerchantName      string
	OTR               float32
	AdminFee          float32
	InstallmentAmount float32
	InterestAmount    float32
	AssetName         string
	InterestRateId    int
	InterestRate      float32
	TransactionDate   time.Time
	Tenor             string
}
