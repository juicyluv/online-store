package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/juicyluv/online-store/internal/domain/proto/user"
	userGrpc "github.com/juicyluv/online-store/internal/user/delivery/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

type application struct {
}

func New() *application {
	return &application{}
}

func (a *application) Run() error {
	l, err := net.Listen("tcp", ":3000")

	if err != nil {
		return err
	}

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, userGrpc.New())

	go s.Serve(l)

	conn, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	mux := runtime.NewServeMux()

	err = user.RegisterUserServiceHandler(context.Background(), mux, conn)

	if err != nil {
		return err
	}

	srv := &http.Server{Addr: ":3001", Handler: mux}

	return srv.ListenAndServe()
}
