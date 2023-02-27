package awards

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luckydraw-backend/internal/logic/awards"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
)

func SelectAwardsByIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectAwardsByIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := awards.NewSelectAwardsByIdsLogic(r.Context(), svcCtx)
		resp, err := l.SelectAwardsByIds(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
