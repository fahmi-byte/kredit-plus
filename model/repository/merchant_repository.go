package repository

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
)

type MerchantRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, merchantId int) (domain.Merchant, error)
}
