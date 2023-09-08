package reports

import (
	"net/http"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/logic/reports"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserReportsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReportsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := reports.NewGetUserReportsLogic(r.Context(), svcCtx)
		resp, err := l.GetUserReports(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
