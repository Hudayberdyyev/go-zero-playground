package trending

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTrendingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTrendingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTrendingListLogic {
	return &GetTrendingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTrendingListLogic) GetTrendingList(req *types.GetTrendingListReq) (resp *types.GetTrendingListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
