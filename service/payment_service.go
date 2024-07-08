package service

import (
	"context"
	"kredit-plus/model/web"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, transactionChan chan web.TransactionResponse, errChan chan error)
	FindAll(ctx context.Context) []web.TransactionResponseList
}
