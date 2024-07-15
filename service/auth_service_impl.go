package service

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"kredit-plus/exception"
	"kredit-plus/helper"
	"kredit-plus/model/repository"
	"kredit-plus/model/web"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB) *AuthServiceImpl {
	return &AuthServiceImpl{AuthRepository: authRepository, DB: DB}
}

func (service *AuthServiceImpl) AuthRegister(ctx context.Context, request web.RegisterRequest) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service.AuthRepository.CreateUser(ctx, tx, request)
}

func (service *AuthServiceImpl) AuthLogin(ctx context.Context, email string, password string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.AuthRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	helper.PanicIfError(err)

	token, err := helper.GenerateJWT(user.Id, user.Username, user.Email, user.PhoneNumber)
	helper.PanicIfError(err)

	return token
}

func (service *AuthServiceImpl) VerificationApiKey(ctx context.Context, apiKey string) bool {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	isValid := service.AuthRepository.ValidateApiKey(ctx, tx, apiKey)

	return isValid
}
