package userPlaylists

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPlaylistsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPlaylistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPlaylistsLogic {
	return &GetUserPlaylistsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPlaylistsLogic) GetUserPlaylists(req *types.GetUserPlaylistsReq) (resp *types.GetUserPlaylistsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
