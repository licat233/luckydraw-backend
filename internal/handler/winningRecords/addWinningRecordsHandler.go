/*
 * @Author: licat
 * @Date: 2023-02-20 23:49:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 12:26:29
 * @Description: licat233@gmail.com
 */
package winningRecords

import (
	"net/http"

	"luckydraw-backend/internal/logic/winningRecords"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddWinningRecordsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddWinningRecordsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := winningRecords.NewAddWinningRecordsLogic(r.Context(), svcCtx, r)
		resp, err := l.AddWinningRecords(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
