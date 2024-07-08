package repository

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
)

type InstallmentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, installments []domain.InstallmentDetail) bool
}
