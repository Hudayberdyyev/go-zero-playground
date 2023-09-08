package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserPlaylistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserPlaylistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserPlaylistLogic {
	return &CreateUserPlaylistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserPlaylistLogic) CreateUserPlaylist(req *types.CreateUserPlaylistReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
