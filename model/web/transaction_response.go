package web

import "time"

type TransactionResponse struct {
	Id                int
	MerchantID        int
	CustomerId        int
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

type TransactionResponseList struct {
	Id                int
	CustomerName      string
	MerchantName      string
	OTR               float32
	AdminFee          float32
	InstallmentAmount float32
	InterestAmount    float32
	AssetName         string
	InterestRate      float32
	TransactionDate   time.Time
	Tenor             string
}
