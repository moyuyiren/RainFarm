package logic

import (
	"MyRainFarm/application/user/rpc/internal/code"
	"MyRainFarm/pkg/encrypt"
	"MyRainFarm/pkg/jwt"
	"context"
	"strconv"

	"MyRainFarm/application/user/rpc/internal/svc"
	"MyRainFarm/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *__.LoginRequest) (*__.LoginResponse, error) {
	// todo: add your logic here and delete this line
	passEnc := encrypt.EncPassword(in.Password)
	snow, err := l.svcCtx.UserModel.GetUserPassword(l.ctx, in.Mobile, passEnc)
	if err != nil {
		return nil, code.LoginInfoEmpty
	}
	tokenStu, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.AuthToken.AccessSecret,
		AccessExpire: l.svcCtx.Config.AuthToken.AccessExpire,
		Fields: map[string]interface{}{
			"SnowID": snow,
		},
	})
	if err != nil {
		return nil, code.LoginEmpty
	}
	err = l.svcCtx.Redis.Insert(l.ctx, strconv.FormatInt(snow, 10))
	if err != nil {
		return nil, code.LoginEmpty
	}

	return &__.LoginResponse{
		AccessToken:  tokenStu.AccessToken,
		AccessExpire: tokenStu.AccessExpire,
	}, nil
}
