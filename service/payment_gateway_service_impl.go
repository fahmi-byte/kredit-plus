package service

import (
	"context"
	"errors"
	"kredit-plus/model/repository"
)

type PaymentGatewayServiceImpl struct {
	Repository repository.PaymentGatewayRepository
}

func NewPaymentGatewayService(repository repository.PaymentGatewayRepository) *PaymentGatewayServiceImpl {
	return &PaymentGatewayServiceImpl{Repository: repository}
}

func (service *PaymentGatewayServiceImpl) PaymentProcess(ctx context.Context, bankAccount string, amount float32, paymentSuccessChan chan bool, errChan chan error) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	result := service.Repository.PaymentProcess(bankAccount, amount)
	if result != nil {
		paymentSuccessChan <- true
	} else {
		errChan <- errors.New("Failed Process Payment")
	}
}
