package userPlaylists

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPlaylistsByVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPlaylistsByVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPlaylistsByVideoLogic {
	return &GetUserPlaylistsByVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPlaylistsByVideoLogic) GetUserPlaylistsByVideo(req *types.GetUserPlaylistsByVideoReq) (resp *types.GetUserPlaylistsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
