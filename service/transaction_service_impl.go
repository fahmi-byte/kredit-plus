package service

import (
	"context"
	"database/sql"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"kredit-plus/model/repository"
	"kredit-plus/model/web"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *sql.DB
}

func NewTransactionService(transactionRepository repository.TransactionRepository, DB *sql.DB) *TransactionServiceImpl {
	return &TransactionServiceImpl{TransactionRepository: transactionRepository, DB: DB}
}

func (t *TransactionServiceImpl) Create(ctx context.Context, transactionChan chan web.TransactionCreateRequest, transactionResponseChan chan web.TransactionResponse, transactionValidChan chan bool, errChan chan error) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	requestTransaction := <-transactionChan

	transaction := domain.Transaction{
		CustomerId:     requestTransaction.CustomerId,
		MerchantId:     requestTransaction.MerchantId,
		OTR:            requestTransaction.OTR,
		AdminFee:       requestTransaction.AdminFee,
		AssetName:      requestTransaction.AssetName,
		InterestRateId: requestTransaction.InterestRateId,
		Tenor:          requestTransaction.Tenor,
	}
	transactionResult := t.TransactionRepository.Save(ctx, tx, transaction)
	transactionResponseChan <- web.TransactionResponse{
		Id:                transactionResult.Id,
		MerchantID:        transactionResult.MerchantId,
		CustomerId:        transactionResult.CustomerId,
		Tenor:             transactionResult.Tenor,
		OTR:               transactionResult.OTR,
		AdminFee:          transactionResult.AdminFee,
		AssetName:         transactionResult.AssetName,
		InterestRateId:    transactionResult.InterestRateId,
		InterestAmount:    transactionResult.InterestAmount,
		InstallmentAmount: transactionResult.InstallmentAmount,
		TransactionDate:   transactionResult.TransactionDate,
	}
	transactionValidChan <- true
}

func (t *TransactionServiceImpl) FindAll(ctx context.Context) []web.TransactionResponseList {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactions := t.TransactionRepository.FindAll(ctx, tx)

	return helper.ToTransactionListResponse(transactions)
}
