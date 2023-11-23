package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Apparatus int64
	AuthToken struct {
		AccessSecret string
		AccessExpire int64
	}
	DB struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
	RedisHost struct {
		RedisSource string
		Password    string
		DB          int
	}
}
