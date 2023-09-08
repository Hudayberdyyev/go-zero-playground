package search

import (
	"context"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SuggestionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSuggestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuggestionsLogic {
	return &SuggestionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SuggestionsLogic) Suggestions(req *types.SuggestionReq) (resp *types.SuggestionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
