package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strings"
)

const (
	//EnvDev     = "development"
	EnvProd = "production"
)

type Config struct {
	// Env
	Environment string `envconfig:"ENVIRONMENT" required:"true" desc:"[production]"`
	// Service
	HttpPort string `envconfig:"HTTP_PORT" required:"true" desc:"[8080] http port"`
	// Log
	Logger   *logrus.Logger
	LogLevel string `envconfig:"LOG_LEVEL" required:"true" desc:"[info,debug] log levels"`
	// Postgres
	PostgresHost     string `envconfig:"POSTGRES_HOST" required:"true" desc:"[127.0.0.1]"`
	PostgresPort     int    `envconfig:"POSTGRES_PORT" required:"true" desc:"[5432]"`
	PostgresUser     string `envconfig:"POSTGRES_USER" required:"true" desc:"[postgres]"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true" desc:"[]"`
	PostgresSsl      string `envconfig:"POSTGRES_SSL" required:"true" desc:"[disable]"`
	PostgresDatabase string `envconfig:"POSTGRES_DATABASE" required:"true" desc:"[postgres]"`
	// Redis
	RedisHost     string `envconfig:"REDIS_HOST" required:"true" desc:"[127.0.0.1] Redis host"`
	RedisPort     int    `envconfig:"REDIS_PORT" required:"true" desc:"[6379] Redis port"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" required:"true" desc:"Redis password"`
	RedisDatabase int    `envconfig:"REDIS_DATABASE" required:"true" desc:"[0] Redis database index"`
}

func (c Config) DbConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", c.PostgresUser,
		url.QueryEscape(c.PostgresPassword), c.PostgresHost, c.PostgresPort, c.PostgresDatabase, c.PostgresSsl,
	)
}

func (c Config) RdbConnectionString() string {
	return fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort)
}

func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)
	return l
}

func Load(logger *logrus.Logger) *Config {
	cfg := &Config{
		Logger: logger,
	}
	err := envconfig.Process("", cfg)
	if err != nil {
		logger.WithError(err).Fatal("failed to process envconfig")
	}

	cfg.Logger.SetLevel(getLogLevel(cfg.LogLevel))
	return cfg
}

func getLogLevel(s string) logrus.Level {
	switch strings.ToLower(s) {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	default:
		return logrus.DebugLevel
	}
}

func GetConfig() *Config {
	logger := NewLogger()
	cfg := Load(logger)
	return cfg
}
