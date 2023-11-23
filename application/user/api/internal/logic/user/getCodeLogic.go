package user

import (
	"MyRainFarm/application/user/api/internal/code"
	"MyRainFarm/application/user/api/internal/svc"
	"MyRainFarm/application/user/api/internal/types"
	"MyRainFarm/application/user/rpc/usercenter"
	"MyRainFarm/pkg/util"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeLogic {
	return &GetCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCodeLogic) GetCode(req *types.GetCodeReq) (err error) {
	//todo: 调用rpc 写入userPhone和code
	codeStr, err := l.svcCtx.UserRpc.GetCode(l.ctx, &usercenter.UserCodeRequest{
		UserEmail: req.Email,
	})
	//todo: 发送验证码
	err = util.SendCodeEmail(req.Email, codeStr.UserCode)
	if err != nil {
		return code.CodeKeyEmpty
	}
	return nil
}
