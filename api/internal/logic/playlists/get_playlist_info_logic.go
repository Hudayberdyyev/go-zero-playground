package playlists

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlaylistInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPlaylistInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlaylistInfoLogic {
	return &GetPlaylistInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPlaylistInfoLogic) GetPlaylistInfo(req *types.GetPlaylistInfoReq) (resp *types.GetPlaylistInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
