package authgrpc

import (
	"GoAPI/protos/gen/proto"
	"context"
	"google.golang.org/grpc"
)

type serverAPI struct {
	proto.AuthServer
}

func (server *serverAPI) CreateUser(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	// TODO: Вызвать валидацию и создание пользователя в бд
	return &proto.CreateResponse{
		UserId: 1,
	}, nil
}
func (server *serverAPI) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// TODO: Вызвать валидацию и результат авторизации

	return &proto.LoginResponse{Token: "token"}, nil
}
func Register(gRPC *grpc.Server) {
	proto.RegisterAuthServer(gRPC, &serverAPI{})
}
