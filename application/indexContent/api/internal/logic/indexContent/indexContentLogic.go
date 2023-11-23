package indexContent

import (
	"context"

	"MyRainFarm/application/indexContent/api/internal/svc"
	"MyRainFarm/application/indexContent/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexContentLogic {
	return &IndexContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexContentLogic) IndexContent() (resp *types.IndexContentResp, err error) {
	// todo: add your logic here and delete this line

	return &types.IndexContentResp{
		HotMessage:     nil,
		MessageInfo:    nil,
		ProductPicture: nil,
	}
}
