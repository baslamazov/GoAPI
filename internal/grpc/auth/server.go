package authgrpc

import (
	"GoAPI/protos/gen/proto"
	"context"

	"google.golang.org/grpc"
)

type serverAPI struct {
	proto.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	proto.RegisterAuthServer(gRPC, &serverAPI{})
}
func (server *serverAPI) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	return &proto.LoginResponse{Token: "token"}, nil
}
func (server *serverAPI) SingUp(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	panic("non implement")
}
