package service

import (
	"context"
	"database/sql"
	"kredit-plus/helper"
	"kredit-plus/model/repository"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB) *AuthServiceImpl {
	return &AuthServiceImpl{AuthRepository: authRepository, DB: DB}
}

func (service *AuthServiceImpl) Login() {
	//TODO implement me
	panic("implement me")
}

func (service *AuthServiceImpl) VerificationApiKey(ctx context.Context, tx *sql.Tx, apiKey string) bool {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	isValid := service.AuthRepository.ValidateApiKey(ctx, tx, apiKey)

	return isValid
}
