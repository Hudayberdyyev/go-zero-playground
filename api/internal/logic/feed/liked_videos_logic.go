package feed

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikedVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikedVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikedVideosLogic {
	return &LikedVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikedVideosLogic) LikedVideos(req *types.LikedVideosReq) (resp *types.LikedVideosResp, err error) {
	// todo: add your logic here and delete this line

	return
}
