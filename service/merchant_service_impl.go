package service

import (
	"context"
	"database/sql"
	"kredit-plus/helper"
	"kredit-plus/model/repository"
	"kredit-plus/model/web"
)

type MerchantServiceImpl struct {
	MerchantRepository repository.MerchantRepository
	DB                 *sql.DB
}

func NewMerchantService(merchantRepository repository.MerchantRepository, DB *sql.DB) *MerchantServiceImpl {
	return &MerchantServiceImpl{MerchantRepository: merchantRepository, DB: DB}
}

func (service *MerchantServiceImpl) FindWithId(ctx context.Context, merchantId int) web.MerchantResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	merchant, err := service.MerchantRepository.FindById(ctx, tx, merchantId)
	if err != nil {
		panic("NOT FOUND")
	}

	return helper.ToMerchantResponse(merchant)
}
