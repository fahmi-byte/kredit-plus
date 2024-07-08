package repository

import (
	"context"
	"database/sql"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() *AuthRepositoryImpl {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) Login() {
	//TODO implement me
	panic("implement me")
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
