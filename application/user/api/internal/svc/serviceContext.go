package svc

import (
	"MyRainFarm/application/user/api/internal/config"
	"MyRainFarm/application/user/rpc/usercenter"
	"MyRainFarm/pkg/interceptors"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:  c,
		UserRpc: usercenter.NewUserCenter(userRPC),
	}
}
