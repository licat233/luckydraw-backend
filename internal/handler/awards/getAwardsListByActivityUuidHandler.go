package awards

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luckydraw-backend/internal/logic/awards"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
)

func GetAwardsListByActivityUuidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAwardsListByActivityUuidReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := awards.NewGetAwardsListByActivityUuidLogic(r.Context(), svcCtx)
		resp, err := l.GetAwardsListByActivityUuid(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
