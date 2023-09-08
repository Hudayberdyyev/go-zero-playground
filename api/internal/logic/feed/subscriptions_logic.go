package feed

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscriptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubscriptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscriptionsLogic {
	return &SubscriptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubscriptionsLogic) Subscriptions(req *types.SubscriptionsReq) (resp *types.SubscriptionsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
