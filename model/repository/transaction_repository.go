package repository

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction
}
