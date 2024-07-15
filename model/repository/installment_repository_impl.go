package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"strconv"
	"time"
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

func (repository *InstallmentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transaction string) {
	transactionId, err := helper.ExtractNumber(transaction)
	date := time.Now()
	helper.PanicIfError(err)

	var installmentId int
	var totalPayment float32
	var tenor string
	var customerId int
	query := "SELECT id.id, id.total_payment, t.tenor, t.customer_id FROM installment_details id JOIN transactions t ON id.transaction_id = t.id WHERE transaction_id = $1  AND payment_status = false ORDER BY month LIMIT 1"
	err = tx.QueryRowContext(ctx, query, transactionId).Scan(&installmentId, &totalPayment, &tenor, &customerId)
	helper.PanicIfError(err)

	SQL := "UPDATE installment_details SET payment_status = true, payment_date = $1 WHERE id = $2"
	_, err = tx.ExecContext(ctx, SQL, date, installmentId)
	helper.PanicIfError(err)

	sqlUpdateTenor := fmt.Sprintf("UPDATE tenor_customers SET %s = %s + $1 WHERE customer_id = $2", tenor, tenor)
	_, err = tx.ExecContext(ctx, sqlUpdateTenor, totalPayment, customerId)
	helper.PanicIfError(err)
}

func (repository *InstallmentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, installmentId int) (domain.InstallmentDetail, error) {
	installment := domain.InstallmentDetail{}
	SQL := "select id.id, id.transaction_id, t.contract_number, id.date, id.month, id.amount, id.late_charge, id.total_payment, id.payment_status, id.payment_date, id.due_date from installment_details id join transactions t ON id.transaction_id = t.id where id.id = $1"
	err := tx.QueryRowContext(ctx, SQL, installmentId).Scan(&installment.Id, &installment.TransactionId, &installment.ContractNumber, &installment.Date, &installment.Month, &installment.Amount, &installment.LateCharge, &installment.TotalPayment, &installment.PaymentStatus, &installment.PaymentDate, &installment.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return installment, errors.New("Installment with id " + strconv.Itoa(installmentId) + "Not Found")
		} else {
			panic(err)
		}
	}
	return installment, nil
}

func (repository *InstallmentRepositoryImpl) FindAllById(ctx context.Context, tx *sql.Tx, customerId int) ([]domain.InstallmentDetail, error) {
	SQL := "select id.id, id.transaction_id, id.date, id.month, id.amount, id.late_charge, id.total_payment, id.payment_status, id.payment_date, id.due_date from installment_details id join transactions t ON id.transaction_id = t.id where t.customer_id = $1"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	installments := []domain.InstallmentDetail{}
	for rows.Next() {
		installment := domain.InstallmentDetail{}
		err := rows.Scan(&installment.Id, &installment.TransactionId, &installment.Date, &installment.Month, &installment.Amount, &installment.LateCharge, &installment.TotalPayment, &installment.PaymentStatus, &installment.PaymentDate, &installment.DueDate)
		helper.PanicIfError(err)
		installments = append(installments, installment)
	}
	return installments, nil
}
