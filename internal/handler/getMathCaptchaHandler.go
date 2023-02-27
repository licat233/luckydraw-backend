package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luckydraw-backend/internal/logic"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
)

func GetMathCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetMathCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetMathCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
