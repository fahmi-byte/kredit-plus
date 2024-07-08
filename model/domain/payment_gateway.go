package domain

import "time"

type PaymentGateway struct {
	Status          string
	TransactionId   string
	Amount          float32
	TransactionTime time.Time
}
