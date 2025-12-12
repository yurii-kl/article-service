package app

import (
	httpapp "github.com/Article/article-service/internal/app/http"
	"github.com/Article/article-service/pkg/postgres"
	"github.com/Article/article-service/pkg/redis"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type App struct {
	HttpServer *httpapp.Server
}

func New(
	log *logrus.Logger,
	dbURI string,
	redisAddr string,
	redisPassword string,
	redisDb int,
) *App {
	db := postgres.New(postgres.Config{
		Log: log,
		Pgc: postgres.PostgresConfig{
			Dsn: dbURI,
		},
	})

	rdb := redis.NewRedisClient(redis.Config{
		Log: log,
		Rdc: redis.ConfigRedis{
			RedisAddr: redisAddr,
			Password:  redisPassword,
			Database:  redisDb,
		},
	})

	validate := validator.New()

	httpApp := httpapp.NewServer(&httpapp.ServerOption{
		Db:        db,
		RDb:       &rdb,
		Log:       log,
		Validator: validate,
	})

	return &App{
		HttpServer: httpApp,
	}
}

func (a *App) Stop() {
	a.HttpServer.Stop()
}
