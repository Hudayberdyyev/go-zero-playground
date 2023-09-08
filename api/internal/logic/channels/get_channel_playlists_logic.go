package channels

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelPlaylistsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChannelPlaylistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelPlaylistsLogic {
	return &GetChannelPlaylistsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChannelPlaylistsLogic) GetChannelPlaylists(req *types.GetChannelPlaylistsReq) (resp *types.GetChannelPlaylistsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
