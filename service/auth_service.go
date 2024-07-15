package service

import (
	"context"
	"kredit-plus/model/web"
)

type AuthService interface {
	AuthRegister(ctx context.Context, request web.RegisterRequest)
	AuthLogin(ctx context.Context, email string, password string) string
	VerificationApiKey(ctx context.Context, apiKey string) bool
}
