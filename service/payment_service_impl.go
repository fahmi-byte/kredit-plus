package service

import (
	"context"
	"database/sql"
	"errors"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"kredit-plus/model/repository"
	"kredit-plus/model/web"
	"time"
)

type PaymentServiceImpl struct {
	PaymentRepository repository.PaymentRepository
	DB                *sql.DB
}

func NewPaymentService(paymentRepository repository.PaymentRepository, DB *sql.DB) *PaymentServiceImpl {
	return &PaymentServiceImpl{PaymentRepository: paymentRepository, DB: DB}
}

func (service *PaymentServiceImpl) CreatePayment(ctx context.Context, transactionChan chan web.TransactionResponse, errChan chan error) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	requestTransaction := <-transactionChan

	payment := domain.Payment{
		MerchantId:    requestTransaction.MerchantID,
		TransactionId: requestTransaction.Id,
		Amount:        requestTransaction.OTR,
		PaymentDate:   time.Now(),
		Status:        true,
	}
	paymentResponse := service.PaymentRepository.Save(ctx, tx, payment)
	if paymentResponse == nil {
		errChan <- errors.New("Create Payment Fail")
	}
	transactionChan <- requestTransaction
}

func (service *PaymentServiceImpl) FindAll(ctx context.Context) []web.TransactionResponseList {
	//TODO implement me
	panic("implement me")
}
