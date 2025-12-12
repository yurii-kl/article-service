package main

import (
	_ "github.com/Article/article-service/docs"
	"github.com/Article/article-service/internal/app"
	"github.com/Article/article-service/pkg/config"
	"os"
	"os/signal"
	"syscall"
)

// @title           Article Service API
// @version         1.0
// @description     Article management service API
// @termsOfService  http://swagger.io/terms/
// @BasePath        /api/v1
// @schemes         http

func main() {
	cfg := config.Load(config.NewLogger())
	log := cfg.Logger

	application := app.New(
		log,
		cfg.DbConnectionString(),
		cfg.RdbConnectionString(),
		cfg.RedisPassword,
		cfg.RedisDatabase)

	go func() {
		err := application.HttpServer.Run()
		if err != nil {
			log.WithError(err).Fatal("http server error")
			return
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	log.Info("Gracefully stopped")
}
