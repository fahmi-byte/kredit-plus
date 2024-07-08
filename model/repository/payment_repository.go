package repository

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
)

type PaymentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) *domain.Payment
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Payment
}
