package user

import (
	"context"
	"github.com/juicyluv/online-store/internal/domain"
)

type UseCase interface {
	GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.User, error)
}
