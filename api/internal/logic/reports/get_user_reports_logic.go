package reports

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserReportsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserReportsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReportsLogic {
	return &GetUserReportsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserReportsLogic) GetUserReports(req *types.GetReportsReq) (resp *types.GetReportsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
