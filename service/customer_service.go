package service

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
)

type CustomerService interface {
	CheckTenorLimitCustomer(ctx context.Context, customerId int, tenor string, otr float32, pin string, ch chan bool, errChan chan error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
}
