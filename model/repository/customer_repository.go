package repository

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
)

type CustomerRepository interface {
	CheckTenorLimit(ctx context.Context, tx *sql.Tx, customerId int, tenor string, otr float32, pin string) bool
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
}
