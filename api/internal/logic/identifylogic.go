package logic

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdentifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdentifyLogic {
	return &IdentifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IdentifyLogic) Identify(req *types.IdentifyReq) (resp *types.IdentifyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
