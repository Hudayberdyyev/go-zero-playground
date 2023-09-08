package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoIncViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoIncViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoIncViewLogic {
	return &VideoIncViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoIncViewLogic) VideoIncView(req *types.VideoIncViewReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
