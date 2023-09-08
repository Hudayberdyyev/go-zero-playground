package playlists

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlaylistVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPlaylistVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlaylistVideosLogic {
	return &GetPlaylistVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPlaylistVideosLogic) GetPlaylistVideos(req *types.GetPlaylistVideosReq) (resp *types.GetPlaylistVideosResp, err error) {
	// todo: add your logic here and delete this line

	return
}
