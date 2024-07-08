package repository

import (
	"context"
	"database/sql"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
)

type InstallmentRepositoryImpl struct {
}

func NewInstallmentRepository() *InstallmentRepositoryImpl {
	return &InstallmentRepositoryImpl{}
}

func (repository *InstallmentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, installments []domain.InstallmentDetail) bool {
	stmt, err := tx.Prepare("insert into installment_details(transaction_id, date, month, amount, late_charge, total_payment, payment_status, payment_date, due_date) values($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	helper.PanicIfError(err)
	defer stmt.Close()

	// Lakukan batch insert dalam transaksi
	for _, row := range installments {
		_, err = tx.Stmt(stmt).Exec(row.TransactionId, row.Date, row.Month, row.Amount, row.LateCharge, row.TotalPayment, row.PaymentStatus, row.PaymentDate, row.DueDate)
		if err != nil {
			helper.PanicIfError(err)
		}
	}

	return true
}
