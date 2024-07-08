package service

import (
	"context"
	"kredit-plus/model/web"
)

type InstallmentService interface {
	CreateInstallment(ctx context.Context, transactionChan chan web.TransactionResponse, paymentCompleteChan chan bool, errChan chan error)
}
