package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoToPlaylistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVideoToPlaylistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoToPlaylistLogic {
	return &AddVideoToPlaylistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVideoToPlaylistLogic) AddVideoToPlaylist(req *types.PlaylistVideoActionReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
