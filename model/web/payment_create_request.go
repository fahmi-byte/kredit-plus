package web

import "time"

type PaymentCreateRequest struct {
	MerchantId    int
	TransactionId int
	Amount        float32
	PaymentDate   time.Time
	Status        bool
}
