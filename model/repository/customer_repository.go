package repository

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
)

type CustomerRepository interface {
	CheckTenorLimit(ctx context.Context, tx *sql.Tx, customerId int, tenor string, otr float32, pin string) error
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
	CreateCustomer(ctx context.Context, tx *sql.Tx, customer domain.Customer)
}
