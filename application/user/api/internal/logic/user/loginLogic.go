package user

import (
	"MyRainFarm/application/user/api/internal/code"
	"MyRainFarm/application/user/rpc/usercenter"
	"context"

	"MyRainFarm/application/user/api/internal/svc"
	"MyRainFarm/application/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	mobile := req.Mobile
	if mobile == "" {
		logx.Error("用户登录参数错误")
		return nil, code.LoginInfoEmpty
	}
	password := req.Password
	if password == "" {
		logx.Error("用户登录参数错误")
		return nil, code.LoginInfoEmpty
	}
	login, err := l.svcCtx.UserRpc.Login(l.ctx, &usercenter.LoginRequest{
		Mobile:   mobile,
		Password: password,
	})
	if err != nil {
		logx.Errorf("用户登录失败", err)
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  login.AccessToken,
		AccessExpire: login.AccessExpire,
	}, nil
}
