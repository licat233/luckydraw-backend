package adminer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luckydraw-backend/internal/logic/adminer"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
)

func GetAdminerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAdminerListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminer.NewGetAdminerListLogic(r.Context(), svcCtx)
		resp, err := l.GetAdminerList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
