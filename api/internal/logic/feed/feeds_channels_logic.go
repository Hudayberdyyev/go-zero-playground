package feed

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedsChannelsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedsChannelsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedsChannelsLogic {
	return &FeedsChannelsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedsChannelsLogic) FeedsChannels(req *types.FeedChannelsReq) (resp *types.FeedChannelsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
