package logic

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/rpc/trending/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/rpc/trending/trending"

	"github.com/zeromicro/go-zero/core/logx"
)

var content = []string{
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
	"trending",
}

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListLogic) GetList(in *trending.Request) (*trending.Response, error) {
	return &trending.Response{
		Total:   20,
		Content: content[:in.Limit],
	}, nil
}
