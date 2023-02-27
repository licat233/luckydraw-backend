package handler

import (
	"net/http"

	"luckydraw-backend/internal/logic"
	"luckydraw-backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRemoteIpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetRemoteIpLogic(r.Context(), svcCtx, r)
		resp, err := l.GetRemoteIp()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
