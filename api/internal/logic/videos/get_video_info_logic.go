package videos

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoInfoLogic {
	return &GetVideoInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoInfoLogic) GetVideoInfo(req *types.GetVideoInfoReq) (resp *types.GetVideoInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
