package repository

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
)

type InstallmentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, installments []domain.InstallmentDetail) bool
	Update(ctx context.Context, tx *sql.Tx, transaction string)
	FindById(ctx context.Context, tx *sql.Tx, installmentId int) (domain.InstallmentDetail, error)
	FindAllById(ctx context.Context, tx *sql.Tx, customerId int) ([]domain.InstallmentDetail, error)
}
