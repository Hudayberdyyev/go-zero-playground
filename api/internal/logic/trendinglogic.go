package logic

import (
	"context"
	"github.com/Hudayberdyyev/go-zero-playground/rpc/trending/trending"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TrendingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTrendingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrendingLogic {
	return &TrendingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TrendingLogic) Trending(req *types.TrendingReq) (resp *types.TrendingResp, err error) {

	rpcResp, err := l.svcCtx.TrendingClient.GetList(l.ctx, &trending.Request{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		return nil, err
	}

	resp = &types.TrendingResp{Content: rpcResp.Content}

	return
}
