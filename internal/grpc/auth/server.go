package auth

import (
	"GRPC-API/protos/gen/proto"
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
	panic("non implement")
}
func (server *serverAPI) Register(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	panic("non implement")
}
