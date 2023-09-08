package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportYoutubePlaylistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportYoutubePlaylistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportYoutubePlaylistLogic {
	return &ImportYoutubePlaylistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportYoutubePlaylistLogic) ImportYoutubePlaylist(req *types.ImportYoutubePlaylistReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
