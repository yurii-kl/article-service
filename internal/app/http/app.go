package httpapp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Article/article-service/internal/app/http/container"
	articleHttp "github.com/Article/article-service/internal/article/transport/http"
	"github.com/Article/article-service/pkg/config"
	"github.com/Article/article-service/pkg/metrics"
	"github.com/Article/article-service/pkg/redis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Server struct {
	engine    *gin.Engine
	cfg       *config.Config
	validator *validator.Validate
	db        *gorm.DB
	rdb       *redis.Rdb
	log       *logrus.Logger
}

type ServerOption struct {
	Db        *gorm.DB
	RDb       *redis.Rdb
	Validator *validator.Validate
	Log       *logrus.Logger
}

func NewServer(option *ServerOption) *Server {
	return &Server{
		db:        option.Db,
		rdb:       option.RDb,
		engine:    gin.Default(),
		cfg:       config.GetConfig(),
		validator: option.Validator,
		log:       option.Log,
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

var server *http.Server

func (s *Server) Run() error {
	_ = s.engine.SetTrustedProxies(nil)

	s.engine.Use(cors.Default())
	s.engine.Use(gin.Recovery())
	s.engine.Use(metrics.PrometheusMiddleware())

	if s.cfg.Environment == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	api := s.engine.Group("/api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.GET("/private/health", healthCheck)
	api.GET("/private/metrics", gin.WrapH(promhttp.Handler()))
	v1 := api.Group("/v1")

	s.Routes(v1)

	port := fmt.Sprintf(":%s", s.cfg.HttpPort)
	s.log.WithField("port", port).Info("Starting server")

	server = &http.Server{
		Addr:              port,
		Handler:           s.engine,
		ReadHeaderTimeout: 5 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.WithField("op", "httpapp.Run").Error("ListenAndServe: ", err)
		return err
	}
	return nil
}

func (s *Server) Routes(v1 *gin.RouterGroup) {
	articleContainer := container.NewArticleContainer(s.db, s.rdb, s.log, s.validator)
	articleHttp.RegisterRoutes(v1, articleContainer)
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	const op = "httpapp.Stop"

	s.log.WithField("op", op).Info("stopping http server")

	err := server.Shutdown(ctx)
	if err != nil {
		s.log.WithError(err).Error("failed to shutdown server gracefully")
		return
	}

	<-ctx.Done()
	s.log.Info("timeout of 5 seconds.")

	s.log.Info("http server stopped")
}
