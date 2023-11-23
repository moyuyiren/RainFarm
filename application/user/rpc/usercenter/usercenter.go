// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"

	"MyRainFarm/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRequest     = __.LoginRequest
	LoginResponse    = __.LoginResponse
	RePwdRequest     = __.RePwdRequest
	RePwdResponse    = __.RePwdResponse
	RegisterRequest  = __.RegisterRequest
	RegisterResponse = __.RegisterResponse
	UserCodeRequest  = __.UserCodeRequest
	UserCodeResponse = __.UserCodeResponse
	UserInfoRequest  = __.UserInfoRequest
	UserInfoResponse = __.UserInfoResponse

	UserCenter interface {
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		RePwd(ctx context.Context, in *RePwdRequest, opts ...grpc.CallOption) (*RePwdResponse, error)
		Detail(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		GetCode(ctx context.Context, in *UserCodeRequest, opts ...grpc.CallOption) (*UserCodeResponse, error)
	}

	defaultUserCenter struct {
		cli zrpc.Client
	}
)

func NewUserCenter(cli zrpc.Client) UserCenter {
	return &defaultUserCenter{
		cli: cli,
	}
}

func (m *defaultUserCenter) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := __.NewUserCenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserCenter) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := __.NewUserCenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserCenter) RePwd(ctx context.Context, in *RePwdRequest, opts ...grpc.CallOption) (*RePwdResponse, error) {
	client := __.NewUserCenterClient(m.cli.Conn())
	return client.RePwd(ctx, in, opts...)
}

func (m *defaultUserCenter) Detail(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := __.NewUserCenterClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}

func (m *defaultUserCenter) GetCode(ctx context.Context, in *UserCodeRequest, opts ...grpc.CallOption) (*UserCodeResponse, error) {
	client := __.NewUserCenterClient(m.cli.Conn())
	return client.GetCode(ctx, in, opts...)
}