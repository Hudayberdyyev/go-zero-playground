package trending

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGenresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGenresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGenresLogic {
	return &GetGenresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGenresLogic) GetGenres(req *types.GetGenresReq) (resp *types.GetGenresResp, err error) {
	// todo: add your logic here and delete this line

	return
}
