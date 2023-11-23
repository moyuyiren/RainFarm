package svc

import (
	"MyRainFarm/application/user/rpc/internal/cache"
	"MyRainFarm/application/user/rpc/internal/config"
	"MyRainFarm/application/user/rpc/internal/model"
	"MyRainFarm/pkg/orm"
	"MyRainFarm/pkg/redis"
)

type ServiceContext struct {
	Config    config.Config
	DB        *orm.DB
	UserModel *model.UserModel
	Redis     *cache.CenterCache
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	rdb := redis.InitRedis(&redis.Config{
		DB:          c.RedisHost.DB,
		RedisSource: c.RedisHost.RedisSource,
		Password:    c.RedisHost.Password,
	})

	return &ServiceContext{
		Config:    c,
		DB:        db,
		UserModel: model.NewUserModel(db.DB),
		Redis:     cache.NewCenterCache(rdb),
	}
}
