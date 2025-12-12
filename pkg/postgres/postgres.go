package postgres

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Timeout = 5
)

type PostgresConfig struct {
	Dsn string
}

type Config struct {
	Pgc PostgresConfig
	Log *logrus.Logger
}

func New(config Config) *gorm.DB {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	db, err := gorm.Open(postgres.Open(config.Pgc.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		config.Log.WithError(err).Error("failed to connect to database")
		return nil
	}

	sqlDB, _ := db.DB()
	err = sqlDB.PingContext(ctx)
	if err != nil {
		config.Log.WithError(err).Error("failed to connect to database")
		return nil
	}

	return db
}
