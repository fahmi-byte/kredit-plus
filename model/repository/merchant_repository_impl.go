package repository

import (
	"context"
	"database/sql"
	"errors"
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
