package repository

import (
	"context"
	"database/sql"
	"kredit-plus/model/domain"
	"kredit-plus/model/web"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, request web.RegisterRequest)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	ValidateApiKey(ctx context.Context, tx *sql.Tx, apiKey string) bool
}
