package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type ConfigRedis struct {
	RedisAddr string
	Password  string
	Database  int
}

type Config struct {
	Rdc ConfigRedis
	Log *logrus.Logger
}

type Rdb struct {
	redis.Cmdable
}

func NewRedisClient(config Config) Rdb {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Rdc.RedisAddr,
		Password: config.Rdc.Password,
		DB:       config.Rdc.Database,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		config.Log.Error("Failed to connect to Redis", err)
	}

	return Rdb{rdb}
}
