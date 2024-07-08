package repository

import (
	"context"
	"database/sql"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"time"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	var interestRate float32
	query := "select interest_rate from interest_rates where id = $1"
	err := tx.QueryRow(query, transaction.InterestRateId).Scan(&interestRate)
	helper.PanicIfError(err)

	transaction.InterestAmount = transaction.OTR * (interestRate / 100)

	switch transaction.Tenor {
	case "tenor_1":
		transaction.InstallmentAmount = (transaction.OTR + transaction.InterestAmount + transaction.AdminFee) / 1
	case "tenor_2":
		transaction.InstallmentAmount = (transaction.OTR + transaction.InterestAmount + transaction.AdminFee) / 2
	case "tenor_3":
		transaction.InstallmentAmount = (transaction.OTR + transaction.InterestAmount + transaction.AdminFee) / 3
	case "tenor_4":
		transaction.InstallmentAmount = (transaction.OTR + transaction.InterestAmount + transaction.AdminFee) / 4
	}

	transaction.TransactionDate = time.Now()
	var lastId int
	SQL := "INSERT INTO transactions(customer_id, merchant_id, contract_number, otr, admin_fee, installment_amount, interest_amount, asset_name, interest_rate_id, transaction_date, tenor) VALUES($1, $2, generate_new_contract_number(), $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"

	// Eksekusi query dengan menggunakan ExecContext
	err = tx.QueryRowContext(ctx, SQL, transaction.CustomerId, transaction.MerchantId, transaction.OTR, transaction.AdminFee, transaction.InstallmentAmount, transaction.InterestAmount, transaction.AssetName, transaction.InterestRateId, transaction.TransactionDate, transaction.Tenor).Scan(&lastId)
	helper.PanicIfError(err)

	transaction.Id = lastId

	return transaction
}

func (repository *TransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction {
	SQL := "select t.id, t.customer_id, c.full_name, t.merchant_id, m.merchant_name, t.otr, t.admin_fee, t.installment_amount, t.interest_amount, t.asset_name, t.interest_rate_id, ir.interest_rate, t.transaction_date, t.tenor from transactions t " +
		"join customers c on c.id = t.customer_id " +
		"join merchants m on m.id = t.merchant_id " +
		"join interest_rates ir on ir.id = t.interest_rate_id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	transactions := []domain.Transaction{}
	for rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.FullName, &transaction.MerchantId, &transaction.MerchantName, &transaction.OTR, &transaction.AdminFee, &transaction.InstallmentAmount, &transaction.InterestAmount, &transaction.AssetName, &transaction.InterestRateId, &transaction.InterestRate, &transaction.TransactionDate, &transaction.Tenor)
		helper.PanicIfError(err)
		transactions = append(transactions, transaction)
	}
	return transactions
}
