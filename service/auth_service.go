package service

import (
	"context"
	"database/sql"
)

type AuthService interface {
	Login()
	VerificationApiKey(ctx context.Context, tx *sql.Tx, apiKey string) bool
}
