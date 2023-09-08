package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPlaylistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPlaylistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPlaylistLogic {
	return &UpdateUserPlaylistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPlaylistLogic) UpdateUserPlaylist(req *types.UpdateUserPlaylistReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
