package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelSubscriptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelSubscriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelSubscriptionLogic {
	return &ChannelSubscriptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelSubscriptionLogic) ChannelSubscription(req *types.ChannelSubscriptionReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
