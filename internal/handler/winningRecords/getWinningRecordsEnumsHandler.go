package winningRecords

import (
	"net/http"

	"luckydraw-backend/internal/logic/winningRecords"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetWinningRecordsEnumsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetWinningRecordsEnumsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := winningRecords.NewGetWinningRecordsEnumsLogic(r.Context(), svcCtx)
		resp, err := l.GetWinningRecordsEnums(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
