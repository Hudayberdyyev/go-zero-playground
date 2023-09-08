package feed

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LibraryLogic {
	return &LibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LibraryLogic) Library(req *types.BaseReqHeaders) (resp *types.LibraryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
