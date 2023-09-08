package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendFeedBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendFeedBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendFeedBackLogic {
	return &SendFeedBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendFeedBackLogic) SendFeedBack(req *types.SendFeedbackReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
