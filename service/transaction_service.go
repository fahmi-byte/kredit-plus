package service

import (
	"context"
	"kredit-plus/model/web"
)

type TransactionService interface {
	Create(ctx context.Context, transactionChan chan web.TransactionCreateRequest, transactionResponseChan chan web.TransactionResponse, errChan chan error)
	FindAll(ctx context.Context) []web.TransactionResponseList
}
