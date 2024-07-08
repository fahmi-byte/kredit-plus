package domain

import "time"

type Payment struct {
	Id            int
	MerchantId    int
	MerchantName  string
	TransactionId int
	Amount        float32
	PaymentDate   time.Time
	Status        bool
}
