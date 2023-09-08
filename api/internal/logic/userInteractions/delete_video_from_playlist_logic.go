package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoFromPlaylistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoFromPlaylistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoFromPlaylistLogic {
	return &DeleteVideoFromPlaylistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVideoFromPlaylistLogic) DeleteVideoFromPlaylist(req *types.PlaylistVideoActionReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
