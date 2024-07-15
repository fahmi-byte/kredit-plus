package service

import (
	"context"
	"database/sql"
	"errors"
	"kredit-plus/helper"
	"kredit-plus/model/repository"
)

type PaymentGatewayServiceImpl struct {
	DB                    *sql.DB
	Repository            repository.PaymentGatewayRepository
	InstallmentRepository repository.InstallmentRepository
}

func NewPaymentGatewayService(repository repository.PaymentGatewayRepository, installmentRepository repository.InstallmentRepository, DB *sql.DB) *PaymentGatewayServiceImpl {
	return &PaymentGatewayServiceImpl{Repository: repository, InstallmentRepository: installmentRepository, DB: DB}
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

func (service *PaymentGatewayServiceImpl) CallbackPayment(ctx context.Context, transaction string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service.InstallmentRepository.Update(ctx, tx, transaction)
}
