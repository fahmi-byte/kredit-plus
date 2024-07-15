package repository

import (
	"context"
	"database/sql"
	"kredit-plus/constants"
	"kredit-plus/model/domain"
)

type MerchantRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, merchantId int) (domain.Merchant, error)
	BalanceUpdate(ctx context.Context, tq *sql.Tx, merchantId int, updateType constants.OperatorType, amount float32)
}
