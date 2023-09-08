package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserPlaylistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserPlaylistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserPlaylistLogic {
	return &DeleteUserPlaylistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserPlaylistLogic) DeleteUserPlaylist(req *types.DeleteUserPlaylistReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
