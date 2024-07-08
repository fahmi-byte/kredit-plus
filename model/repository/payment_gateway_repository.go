package repository

import "kredit-plus/model/domain"

type PaymentGatewayRepository interface {
	PaymentProcess(bankAccount string, amount float32) *domain.PaymentGateway
}
