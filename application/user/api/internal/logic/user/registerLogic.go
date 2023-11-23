package user

import (
	"MyRainFarm/application/user/api/internal/code"
	"MyRainFarm/application/user/rpc/usercenter"
	"MyRainFarm/pkg/encrypt"
	"context"
	"strings"

	"MyRainFarm/application/user/api/internal/svc"
	"MyRainFarm/application/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: 注册信息处理
	req.UserName = strings.TrimSpace(req.UserName)
	if len(req.UserName) == 0 {
		return nil, code.RegisterUserNameEmpty
	}
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	}
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		req.Password = encrypt.EncPassword(req.Password)
	}
	req.Email = strings.TrimSpace(req.Email)
	if len(req.Email) == 0 {
		return nil, code.RegisterEmailEmpty
	}
	// todo: rpc处理
	rpcresp, err := l.svcCtx.UserRpc.Register(l.ctx, &usercenter.RegisterRequest{
		Mobile:   req.Mobile,
		Username: req.UserName,
		Password: req.Password,
		Email:    req.Email,
	})
	if rpcresp == nil && err != nil {
		logx.Errorf("Rpc 用户注册《《UserRpc.Register》》调用失败", err)
		return nil, err
	}
	return &types.RegisterResp{
		AccessToken:  rpcresp.AccessToken,
		AccessExpire: rpcresp.AccessExpire,
	}, nil
	//return nil, xrrcode.ServerErr
}
