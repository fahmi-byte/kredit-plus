package service

import (
	"context"
	"kredit-plus/model/domain"
	"kredit-plus/model/web"
)

type InstallmentService interface {
	CreateInstallment(ctx context.Context, transactionChan chan web.TransactionResponse, installmentCompleteChan chan bool, errChan chan error)
	FindOne(ctx context.Context, installmentId int) domain.InstallmentDetail
	GetInstallmentForPayment(ctx context.Context, installmentId int) web.InstallmentResponsePayment
	FindAllInstallmentCustomer(ctx context.Context, customerId int) []domain.InstallmentDetail
}
