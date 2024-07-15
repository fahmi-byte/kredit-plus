package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"kredit-plus/helper"
	"kredit-plus/model/domain"
	"kredit-plus/model/web"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() *AuthRepositoryImpl {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, request web.RegisterRequest) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	SQL := "INSERT INTO users (username, email, password, phone_number, role_id) VALUES($1, $2, $3, $4, $5)"
	_, err = tx.ExecContext(ctx, SQL, request.Username, request.Email, hashedPassword, request.PhoneNumber, request.RoleId)
	helper.PanicIfError(err)
}

func (repository *AuthRepositoryImpl) ValidateApiKey(ctx context.Context, tx *sql.Tx, apiKey string) bool {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM merchants WHERE api_key = $1)"
	err := tx.QueryRow(query, apiKey).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

func (repository *AuthRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, username, email, password, phone_number FROM users where email = $1 LIMIT 1"
	err := tx.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user with email " + email + "Not Found")
		} else {
			panic(err)
		}
	}
	return user, nil
}
