package activity

import (
	"net/http"

	"luckydraw-backend/internal/logic/activity"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetActivityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetActivityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := activity.NewGetActivityLogic(r.Context(), svcCtx)
		resp, err := l.GetActivity(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
