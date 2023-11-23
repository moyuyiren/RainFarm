package logic

import (
	"MyRainFarm/application/user/rpc/internal/code"
	"MyRainFarm/application/user/rpc/internal/model"
	"MyRainFarm/pkg/util"
	"context"

	"MyRainFarm/application/user/rpc/internal/svc"
	"MyRainFarm/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeLogic {
	return &GetCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCodeLogic) GetCode(in *__.UserCodeRequest) (*__.UserCodeResponse, error) {
	// todo: add your logic here and delete this line
	codeStr := util.RandomNumeric()
	// todo: 查询用户信息
	user := &model.User{}
	err := l.svcCtx.UserModel.GetUserInfoForEmail(l.ctx, in.UserEmail, user)
	if err != nil {
		return nil, code.UserCodeEmpty
	}
	// todo： 是否已经发送过验证码
	again := l.svcCtx.Redis.GetUserCodeAgain(l.ctx, user.Mobile)
	if again != "" {
		return nil, code.UserCodeAgainEmpty
	}
	// todo： 写入缓存
	err = l.svcCtx.Redis.InsertUserCode(l.ctx, user.Mobile, codeStr)
	if err != nil {
		return nil, code.UserCodeEmpty
	}
	return &__.UserCodeResponse{
		UserCode: codeStr,
	}, nil
}
