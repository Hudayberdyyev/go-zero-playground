package userInteractions

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendReportLogic {
	return &SendReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendReportLogic) SendReport(req *types.SendReportReq) (resp *types.MessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
