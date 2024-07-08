package service

import (
	"context"
	"kredit-plus/model/web"
)

type MerchantService interface {
	FindWithId(ctx context.Context, merchantId int) web.MerchantResponse
}
