package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoReactionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoReactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoReactionLogic {
	return &VideoReactionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoReactionLogic) VideoReaction(req *types.VideoReactionReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
