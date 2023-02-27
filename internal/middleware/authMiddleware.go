package middleware

import (
	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/jwtx"
	"luckydraw-backend/internal/config"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	Config config.Config
}

func NewAuthMiddleware(c config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		Config: c,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取token
		jwtToken := r.Header.Get("Authorization")
		if jwtToken == "" {
			httpx.Error(w, errorx.ErrAuth)
			return
		}
		//解析 token
		claims, err := jwtx.ParseToken(jwtToken, m.Config.Auth.AccessSecret)
		if err != nil {
			httpx.Error(w, errorx.InternalError(err))
			return
		}

		// 获取身份
		access := jwtx.GetClaimsValue("access", *claims)

		//获取url
		url := r.URL.Path
		if strings.HasPrefix(url, "/api/adminer/") {
			if access != "super" {
				httpx.Error(w, errorx.ErrAuth)
				return
			}
		}

		if access != "super" && access != "admin" {
			httpx.Error(w, errorx.ErrAuth)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
