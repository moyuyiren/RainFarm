package logic

import (
	"MyRainFarm/application/user/rpc/internal/code"
	"MyRainFarm/application/user/rpc/internal/model"
	"MyRainFarm/pkg/xrrcode"
	"context"

	"MyRainFarm/application/user/rpc/internal/svc"
	"MyRainFarm/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *__.UserInfoRequest) (*__.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	snowID := in.UserSnowID
	if snowID == "" {
		logx.Error("获取用户信息失败")
		return nil, code.GetUserInfoEmpty
	}
	user := &model.User{
		Name:   "",
		Mobile: "",
		Email:  "",
	}

	if err := l.svcCtx.UserModel.GetUserInfo(l.ctx, snowID, user); err != nil {
		logx.Errorf("数据库查询错误", err)
		return nil, xrrcode.ServerErr
	}

	return &__.UserInfoResponse{
		Mobile:   user.Mobile,
		Username: user.Name,
		Email:    user.Email,
	}, nil
}
