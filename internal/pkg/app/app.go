package app

import (
	"GoAPI/internal/app/endpoint"
	"GoAPI/internal/app/mw"
	"GoAPI/internal/config"
	"GoAPI/internal/storage/postgres"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type App struct {
	handler   *endpoint.Endpoint
	dbService *postgres.DBService
	//config    *config.Config
}

func New(conf *config.Config) (*App, error) {
	app := &App{}

	app.dbService = postgres.NewDBService(conf.Database)
	app.handler = endpoint.New(app.dbService)
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
	router.Run(":8083")

	return nil
}
