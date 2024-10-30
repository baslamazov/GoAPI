package grpcapp

import (
	authgrpc "GoAPI/internal/grpc/auth"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	// Запуск gRPC сервера
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic(err)
	}

	go func() {
		if err := a.gRPCServer.Serve(l); err != nil {
			log.Error("gRPC server failed", slog.String("error", err.Error()))
		}
	}()

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))
	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)
	log.Info("Stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
