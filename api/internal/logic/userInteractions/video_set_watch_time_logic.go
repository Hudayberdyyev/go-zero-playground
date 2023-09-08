package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoSetWatchTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoSetWatchTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoSetWatchTimeLogic {
	return &VideoSetWatchTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoSetWatchTimeLogic) VideoSetWatchTime(req *types.VideoSetWatchTimeReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
