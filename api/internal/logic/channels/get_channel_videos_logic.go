package channels

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChannelVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelVideosLogic {
	return &GetChannelVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChannelVideosLogic) GetChannelVideos(req *types.GetChannelVideosReq) (resp *types.GetChannelVideosResp, err error) {
	// todo: add your logic here and delete this line

	return
}
