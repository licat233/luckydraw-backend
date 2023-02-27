/*
 * @Author: licat
 * @Date: 2023-02-22 14:28:35
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 19:47:57
 * @Description: licat233@gmail.com
 */
package adminer

import (
	"context"
	"strings"
	"time"

	"luckydraw-backend/common/ctxdata"
	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminerLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminerLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminerLoginLogic {
	return &AdminerLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminerLoginLogic) AdminerLogin(req *types.AdminerLoginReq) (any, error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.New("账号密码错误")
	}

	if !l.svcCtx.CaptchaStore.Verify(req.CaptchaId, req.Solution, true) {
		return nil, errorx.New("验证码错误")
	}

	adminer, err := l.svcCtx.AdminerModel.FindByUsername(l.ctx, req.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(err)
	}

	if err == model.ErrNotFound {
		return nil, errorx.New("用户不存在")
	}
	if adminer.Password != req.Password {
		return nil, errorx.New("用户密码不正确")
	}
	var second int64 = l.svcCtx.Config.Auth.AccessExpire
	if req.AutoLogin {
		second = 3600 * 24 * 7
	}
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, second, adminer.Id, adminer.IsSuper)
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	expire := time.Now().Add(time.Duration(second) * time.Second).Unix()
	data := &types.JwtToken{
		AccessToken:  token,
		AccessExpire: expire,
		RefreshAfter: expire,
	}
	return respx.SingleResp("登录成功", data), nil
}

func (l *AdminerLoginLogic) getJwtToken(secretKey string, seconds, adminerId, isSuper int64) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	if isSuper == 1 {
		claims["access"] = "super"
	} else {
		claims["access"] = "admin"
	}

	claims[ctxdata.CtxKeyJwtAdminerId] = adminerId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
