package repository

import (
	"context"
	"database/sql"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepository() *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{}
}

func (p PaymentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) *domain.Payment {
	//TODO implement me
	SQL := "insert into payment_merchants(merchant_id, transaction_id, amount, payment_date, status) values($1,$2,$3,$4,$5)"
	_, err := tx.ExecContext(ctx, SQL, payment.MerchantId, payment.TransactionId, payment.Amount, payment.PaymentDate, payment.Status)
	helper.PanicIfError(err)

	return &payment
}

func (p PaymentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Payment {
	SQL := "select pm.id, pm.merchant_id, m.merchant_name, transaction_id, amount, payment_date, status from payment_merchants pm " +
		"join transactions t on t.id = pm.transaction_id " +
		"join merchants m on m.id = t.merchant_id "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	payments := []domain.Payment{}
	for rows.Next() {
		payment := domain.Payment{}
		err := rows.Scan(&payment.Id, &payment.MerchantId, &payment.MerchantName, &payment.TransactionId, &payment.Amount, &payment.PaymentDate, &payment.Status)
		helper.PanicIfError(err)
		payments = append(payments, payment)
	}
	return payments
}
