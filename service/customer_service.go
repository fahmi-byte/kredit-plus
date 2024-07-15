package service

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
	"kredit-plus/model/web"
)

type CustomerService interface {
	CheckTenorLimitCustomer(ctx context.Context, customerId int, tenor string, otr float32, pin string, validTenorChan chan bool, errChan chan error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
	CreateCustomer(ctx context.Context, customer web.CustomerRequest)
}
