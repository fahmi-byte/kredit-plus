package service

import "context"

type PaymentGatewayService interface {
	PaymentProcess(ctx context.Context, bankAccount string, amount float32, paymentSuccessChan chan bool, errChan chan error)
}
