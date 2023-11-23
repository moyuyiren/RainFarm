package user

import (
	"MyRainFarm/application/user/rpc/usercenter"
	"context"

	"MyRainFarm/application/user/api/internal/svc"
	"MyRainFarm/application/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail() (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	userinfo, err := l.svcCtx.UserRpc.Detail(l.ctx, &usercenter.UserInfoRequest{UserSnowID: l.ctx.Value("SnowID").(string)})
	if err != nil {
		logx.Errorf("获取用户信息失败", err)
	}
	return &types.UserInfoResp{
		UserName: userinfo.Username,
		Mobile:   userinfo.Mobile,
		Email:    userinfo.Email,
	}, nil
}
