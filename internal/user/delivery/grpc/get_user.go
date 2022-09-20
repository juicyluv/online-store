package grpc

import (
	"context"
	"github.com/juicyluv/online-store/internal/domain/proto/user"
)

func (s *server) Test(_ context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	return &user.GetUserResponse{Name: req.Msg}, nil
}
