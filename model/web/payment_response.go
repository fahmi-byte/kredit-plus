package web

import "time"

type PaymentResponse struct {
	Id            int
	MerchantId    int
	TransactionId int
	Amount        float32
	PaymentDate   time.Time
	Status        bool
}
