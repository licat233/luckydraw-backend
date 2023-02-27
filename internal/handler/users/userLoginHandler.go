/*
 * @Author: licat
 * @Date: 2023-02-22 14:28:35
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 14:30:23
 * @Description: licat233@gmail.com
 */
package users

import (
	"net/http"

	"luckydraw-backend/internal/logic/users"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := users.NewUserLoginLogic(r.Context(), svcCtx, r)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
