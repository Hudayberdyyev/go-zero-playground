package userInteractions

import (
	"net/http"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/logic/userInteractions"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserPlaylistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserPlaylistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userInteractions.NewUpdateUserPlaylistLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserPlaylist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
