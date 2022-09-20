package auth

import (
	"context"
	"github.com/juicyluv/online-store/internal/domain"
)

type Repository interface {
	SignIn(ctx context.Context, request *domain.SignInRequest) (string, error)
}
