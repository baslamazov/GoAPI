package app

import (
	"GoAPI/internal/app/endpoint"
	grpcapp "GoAPI/internal/app/grpc"
	"GoAPI/internal/app/mw"
	"GoAPI/internal/config"
	"GoAPI/internal/storage/postgres"
	"log/slog"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type App struct {
	handler    *endpoint.Endpoint
	dbService  *postgres.DBService
	GRPCServer *grpcapp.App
	log        *slog.Logger
}

func New(conf *config.Config, log *slog.Logger) (*App, error) {
	app := &App{
		log: log,
	}

	app.dbService = postgres.NewDBService(conf.Database)
	app.handler = endpoint.New(app.dbService)
	app.GRPCServer = grpcapp.New(log, conf.GRPC.Port)
	return app, nil
}

// Старт апи
func (app *App) Start() error {
	router := gin.Default()
	router.Use(gin.Recovery())

	// Инициализация хранилища сессий Redis
	store, err := redis.NewStore(10, "tcp", "localhost:6380", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	router.Use(sessions.Sessions("usersession", store))

	router.POST("/login", app.handler.Auth)
	apiGroup := router.Group("/api", mw.SessionValidate)
	apiGroup.GET("/users", app.handler.GetUsers)
	go router.Run(":8083")

	app.GRPCServer.MustRun()

	return nil
}
