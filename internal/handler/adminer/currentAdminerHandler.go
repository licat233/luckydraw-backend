package adminer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luckydraw-backend/internal/logic/adminer"
	"luckydraw-backend/internal/svc"
)

func CurrentAdminerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminer.NewCurrentAdminerLogic(r.Context(), svcCtx)
		resp, err := l.CurrentAdminer()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
