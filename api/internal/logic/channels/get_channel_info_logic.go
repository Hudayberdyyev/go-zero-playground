package channels

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChannelInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelInfoLogic {
	return &GetChannelInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChannelInfoLogic) GetChannelInfo(req *types.GetChannelInfoReq) (resp *types.GetChannelInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
