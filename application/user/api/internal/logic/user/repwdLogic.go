package user

import (
	"MyRainFarm/application/user/api/internal/code"
	"MyRainFarm/application/user/api/internal/svc"
	"MyRainFarm/application/user/api/internal/types"
	"MyRainFarm/application/user/rpc/usercenter"
	"MyRainFarm/pkg/xrrcode"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type RepwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRepwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RepwdLogic {
	return &RepwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RepwdLogic) Repwd(req *types.ResetPwdReq) (resp *types.ResetPwdResp, err error) {
	// todo: add your logic here and delete this line
	mobile := req.Mobile
	if mobile == "" {
		return nil, code.RePasswdEmpty
	}
	Accesscode := req.AssessCode
	if Accesscode == "" {
		return nil, code.RePasswdEmpty
	}
	passwd := req.Password
	if passwd == "" {
		return nil, code.RePasswdEmpty
	}
	rePwdRpc, err := l.svcCtx.UserRpc.RePwd(l.ctx, &usercenter.RePwdRequest{
		Mobile:   mobile,
		KeyCode:  Accesscode,
		Password: passwd,
	})
	if err != nil {
		return nil, xrrcode.ServerErr
	}
	return &types.ResetPwdResp{
		AccessToken:  rePwdRpc.AccessToken,
		AccessExpire: rePwdRpc.AccessExpire,
	}, nil
}
