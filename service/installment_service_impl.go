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

type InstallmentServiceImpl struct {
	InstallmentRepository repository.InstallmentRepository
	DB                    *sql.DB
}

func NewInstallmentService(installmentRepository repository.InstallmentRepository, DB *sql.DB) *InstallmentServiceImpl {
	return &InstallmentServiceImpl{InstallmentRepository: installmentRepository, DB: DB}
}

func (service *InstallmentServiceImpl) CreateInstallment(ctx context.Context, transactionChan chan web.TransactionResponse, installmentCompleteChan chan bool, errChan chan error) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	requestTransaction := <-transactionChan

	var installmentCount int
	switch requestTransaction.Tenor {
	case "tenor_1":
		installmentCount = 1
	case "tenor_2":
		installmentCount = 2
	case "tenor_3":
		installmentCount = 3
	case "tenor_4":
		installmentCount = 4
	}

	installments := []domain.InstallmentDetail{}

	now := time.Now()
	lastDayOfMonth := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.UTC)

	for i := 1; i <= installmentCount; i++ {
		installments = append(installments, domain.InstallmentDetail{
			TransactionId: requestTransaction.Id,
			Amount:        requestTransaction.InstallmentAmount,
			TotalPayment:  requestTransaction.InstallmentAmount,
			PaymentStatus: false,
			Date:          time.Now(),
			Month:         i,
			DueDate:       lastDayOfMonth,
		})
		lastDayOfMonth = lastDayOfMonth.AddDate(0, 1, 0)
	}

	success := service.InstallmentRepository.Save(ctx, tx, installments)
	if success == false {
		errChan <- errors.New("Fail Insert Batch")
	}

	installmentCompleteChan <- true
}

func (service *InstallmentServiceImpl) FindOne(ctx context.Context, installmentId int) domain.InstallmentDetail {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	result, err := service.InstallmentRepository.FindById(ctx, tx, installmentId)
	helper.PanicIfError(err)

	return result
}

func (service *InstallmentServiceImpl) GetInstallmentForPayment(ctx context.Context, installmentId int) web.InstallmentResponsePayment {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.InstallmentRepository.FindById(ctx, tx, installmentId)
	helper.PanicIfError(err)

	return helper.ToInstallmentPaymentResponse(result)
}

func (service *InstallmentServiceImpl) FindAllInstallmentCustomer(ctx context.Context, customerId int) []domain.InstallmentDetail {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.InstallmentRepository.FindAllById(ctx, tx, customerId)
	helper.PanicIfError(err)

	return result
}
