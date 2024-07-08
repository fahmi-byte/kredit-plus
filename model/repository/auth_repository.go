package repository

import (
	"context"
	"database/sql"
)

type AuthRepository interface {
	Login()
	ValidateApiKey(ctx context.Context, tx *sql.Tx, apiKey string) bool
}
