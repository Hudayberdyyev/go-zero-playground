package playlists

import (
	"net/http"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/logic/playlists"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPlaylistVideosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPlaylistVideosReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := playlists.NewGetPlaylistVideosLogic(r.Context(), svcCtx)
		resp, err := l.GetPlaylistVideos(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
