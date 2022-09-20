package grpc

import "github.com/juicyluv/online-store/internal/domain/proto/user"

type server struct {
	user.UnimplementedUserServiceServer
}

func New() *server {
	return &server{}
}
