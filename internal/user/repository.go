package user

import (
	"context"
	"github.com/juicyluv/online-store/internal/domain"
)

type Repository interface {
	GetUserById(ctx context.Context, userId int64) (*domain.User, error)
}
