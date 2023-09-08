package reports

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReportTypesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReportTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReportTypesLogic {
	return &GetReportTypesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReportTypesLogic) GetReportTypes(req *types.GetReportTypesReq) (resp *types.GetReportTypesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
