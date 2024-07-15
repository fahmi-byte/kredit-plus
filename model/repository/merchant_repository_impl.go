package repository

import (
	"context"
	"database/sql"
	"errors"
	"kredit-plus/constants"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"strconv"
)

type MerchantRepositoryImpl struct {
}

func NewMerchantRepository() *MerchantRepositoryImpl {
	return &MerchantRepositoryImpl{}
}

func (m *MerchantRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, merchantId int) (domain.Merchant, error) {
	SQL := "select merchant_name, bank_account from merchants where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, merchantId)
	helper.PanicIfError(err)
	defer rows.Close()

	merchant := domain.Merchant{}
	if rows.Next() {
		err := rows.Scan(&merchant.MerchantName, &merchant.BankAccount)
		helper.PanicIfError(err)
		merchant.Id = merchantId
		return merchant, nil
	} else {
		return merchant, errors.New("Merchant with id " + strconv.Itoa(merchantId) + " is Not Found")
	}
}

func (m *MerchantRepositoryImpl) BalanceUpdate(ctx context.Context, tx *sql.Tx, merchantId int, updateType constants.OperatorType, amount float32) {
	var balance float32
	query := "SELECT balance FROM merchants WHERE id = $1"
	err := tx.QueryRowContext(ctx, query, merchantId).Scan(&balance)
	helper.PanicIfError(err)

	switch updateType {
	case constants.Increment:
		balance += amount
	case constants.Decrement:
		balance -= amount
	}

	updateQuery := "UPDATE merchants SET balance = $1 WHERE id = $2"
	_, err = tx.ExecContext(ctx, updateQuery, balance, merchantId)
	helper.PanicIfError(err)
}
