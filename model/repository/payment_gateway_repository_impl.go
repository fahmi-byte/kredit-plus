package repository

import (
	"kredit-plus/model/domain"
	"time"
)

type PaymentGatewayRepositoryImpl struct {
}

func NewPaymentGatewayRepository() *PaymentGatewayRepositoryImpl {
	return &PaymentGatewayRepositoryImpl{}
}

func (p PaymentGatewayRepositoryImpl) PaymentProcess(bankAccount string, amount float32) *domain.PaymentGateway {
	return &domain.PaymentGateway{
		Status:          "Success",
		TransactionId:   "TX21230",
		Amount:          amount,
		TransactionTime: time.Now(),
	}
}
