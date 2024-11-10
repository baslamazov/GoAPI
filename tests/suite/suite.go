package suite

import (
	"GoAPI/internal/config"
	"GoAPI/protos/gen/proto"
	"context"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"testing"
)

type GRPCSuite struct {
	suite.Suite
	Conf       *config.Config
	AuthClient proto.AuthClient
	Context    context.Context
}

func New(t *testing.T) *GRPCSuite {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoad()

	ctx, cancelFn := context.WithCancel(context.Background())

	t.Cleanup(CleanupContext(t, cancelFn))

	conn, err := grpc.NewClient(
		cfg.Host+":"+strconv.Itoa(cfg.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		t.Fatalf("Подключение прошло херово: %v", err)
	}

	return &GRPCSuite{
		Conf:       cfg,
		AuthClient: proto.NewAuthClient(conn),
		Context:    ctx,
	}

}

// CleanupContext TODO: Заменить название и суть на завершающую функцию
func CleanupContext(t *testing.T, cancelFunc context.CancelFunc) func() {
	return func() {
		t.Helper()
		cancelFunc()
	}
}
