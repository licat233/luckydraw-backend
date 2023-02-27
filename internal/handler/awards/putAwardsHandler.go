package awards

import (
	"net/http"

	"luckydraw-backend/internal/logic/awards"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PutAwardsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PutAwardsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := awards.NewPutAwardsLogic(r.Context(), svcCtx)
		resp, err := l.PutAwards(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
