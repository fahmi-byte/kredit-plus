package domain

import "time"

type InstallmentDetail struct {
	Id            int
	TransactionId int
	Date          time.Time
	Month         int
	Amount        float32
	LateCharge    float32
	TotalPayment  float32
	PaymentStatus bool
	PaymentDate   time.Time
	DueDate       time.Time
}
