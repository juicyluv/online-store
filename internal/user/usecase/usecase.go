package usecase

import "github.com/juicyluv/online-store/internal/user"

type usecase struct {
	userRepo user.Repository
}

func NewUserUseCase() user.UseCase {
	return &usecase{}
}
