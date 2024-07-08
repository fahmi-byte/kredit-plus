package web

type InstallmentProcessRequest struct {
	CustomerId     int     `validate:"required" json:"customer_id"`
	MerchantId     int     `validate:"required" json:"merchant_id"`
	Pin            string  `validate:"required" json:"pin"`
	OTR            float32 `validate:"required" json:"otr"`
	AdminFee       float32 `validate:"required" json:"admin_fee"`
	InterestRateId int     `validate:"required" json:"interest_rate_id"`
	AssetName      string  `validate:"required" json:"asset_name"`
	Tenor          string  `validate:"required" json:"tenor"`
}
