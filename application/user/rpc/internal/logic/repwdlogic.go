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

type RePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RePwdLogic {
	return &RePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RePwdLogic) RePwd(in *__.RePwdRequest) (*__.RePwdResponse, error) {
	// todo: add your logic here and delete this line
	codeKey := l.svcCtx.Redis.GetUserCodeAgain(l.ctx, in.Mobile)
	if codeKey != in.KeyCode {
		return nil, code.UserCodeAccessEmpty
	}
	// todo： 修改密码，生成新token
	//加密密码
	passEnc := encrypt.EncPassword(in.Password)
	//修改密码并获取用户snowid
	snow, err := l.svcCtx.UserModel.UpdateUserPwd(l.ctx, in.Mobile, passEnc)
	//生成token
	tokenStu, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.AuthToken.AccessSecret,
		AccessExpire: l.svcCtx.Config.AuthToken.AccessExpire,
		Fields: map[string]interface{}{
			"SnowID": snow,
		},
	})
	if err != nil {
		return nil, code.RegisterEmpty
	}
	//将token写入redis
	err = l.svcCtx.Redis.Insert(l.ctx, strconv.FormatInt(snow, 10))
	if err != nil {
		return nil, code.RegisterEmpty
	}

	return &__.RePwdResponse{
		AccessToken:  tokenStu.AccessToken,
		AccessExpire: tokenStu.AccessExpire,
	}, nil
}
