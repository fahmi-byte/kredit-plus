package service

import (
	"context"
	"database/sql"
	"fmt"
	"kredit-plus/constants"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"kredit-plus/model/repository"
	"kredit-plus/model/web"
)

type TransactionServiceImpl struct {
	MerchantRepository    repository.MerchantRepository
	TransactionRepository repository.TransactionRepository
	DB                    *sql.DB
}

func NewTransactionService(transactionRepository repository.TransactionRepository, merchantRepository repository.MerchantRepository, DB *sql.DB) *TransactionServiceImpl {
	return &TransactionServiceImpl{TransactionRepository: transactionRepository, MerchantRepository: merchantRepository, DB: DB}
}

func (t *TransactionServiceImpl) Create(ctx context.Context, transactionChan chan web.TransactionCreateRequest, transactionResponseChan chan web.TransactionResponse, errChan chan error) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	fmt.Println("masuk kedalam function transaction")
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	fmt.Println("masuk kedalam function transaction")
	requestTransaction := <-transactionChan
	fmt.Println("data transaction berhasil dioper dari channel")

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
	t.MerchantRepository.BalanceUpdate(ctx, tx, transactionResult.MerchantId, constants.Increment, transactionResult.OTR)
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
}

func (t *TransactionServiceImpl) FindAll(ctx context.Context) []web.TransactionResponseList {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactions := t.TransactionRepository.FindAll(ctx, tx)

	return helper.ToTransactionListResponse(transactions)
}
