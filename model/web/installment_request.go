package web

type InstallmentRequest struct {
	InstallmentId int `validate:"required" json:"installment_id"`
}
