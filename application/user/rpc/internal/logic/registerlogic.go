package logic

import (
	"MyRainFarm/application/user/rpc/internal/code"
	"MyRainFarm/application/user/rpc/internal/model"
	"MyRainFarm/application/user/rpc/internal/svc"
	"MyRainFarm/application/user/rpc/pb"
	"MyRainFarm/pkg/encrypt"
	"MyRainFarm/pkg/jwt"
	"MyRainFarm/pkg/util"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *__.RegisterRequest) (*__.RegisterResponse, error) {
	// todo: rpc处理逻辑
	// 生成雪花id
	snow, err := util.GetSnow(l.svcCtx.Config.Apparatus)
	if err != nil {
		return nil, code.RegisterEmpty
	}
	// 密码加密
	passEnc := encrypt.EncPassword(in.Password)
	// 组合数据

	user := &model.User{
		SnowID:   snow,
		Name:     in.Username,
		Password: passEnc,
		Mobile:   in.Mobile,
		Email:    in.Email,
	}

	// 写入数据库
	err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		logx.Errorf("数据库写入错误", err)
		return nil, code.RegisterEmpty
	}

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
	return &__.RegisterResponse{
		AccessToken:  tokenStu.AccessToken,
		AccessExpire: tokenStu.AccessExpire,
	}, nil
}
