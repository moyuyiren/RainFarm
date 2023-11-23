package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	DB          int
	RedisSource string
	Password    string
}

func InitRedis(conf *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.RedisSource,
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logx.Errorf("ping Error", err)
		panic(err)
	}
	return rdb
}
