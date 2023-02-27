/*
 * @Author: licat
 * @Date: 2023-02-22 14:28:35
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 14:48:59
 * @Description: licat233@gmail.com
 */
package users

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"luckydraw-backend/common/ctxdata"
	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/common/utils"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (any, error) {
	passport := strings.TrimSpace(req.Passport)
	uuid := strings.TrimSpace(req.ActivityUuid)
	if passport == "" || uuid == "" {
		return nil, errorx.ExternalError(types.ErrParams)
	}
	activity, err := l.svcCtx.ActivityModel.FindOneByUuid(l.ctx, uuid)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("查询活动失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}
	//如果活动不存在
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(errors.New("the activity not found"))
	}
	activityId := activity.Id
	user, err := l.svcCtx.UsersModel.FindsByPassportAndActivityId(l.ctx, passport, activityId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("查询用户失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}
	var userId int64
	//如果用户不存在，不要返回错误，给予其一个jwt，存储其身份凭证
	if err == model.ErrNotFound {
		userId = 0
	} else {
		userId = user.Id
	}
	second := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), second, userId, activityId)
	if err != nil {
		l.Logger.Errorf("生成jwt失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}
	expire := time.Now().Add(time.Duration(second) * time.Second).Unix()
	data := &types.JwtToken{
		AccessToken:  token,
		AccessExpire: expire,
		RefreshAfter: expire,
	}
	return respx.SingleResp("login success", data), nil
}

func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds, userId, ActivityId int64) (string, error) {
	paltform, _ := utils.RemotePlatform(l.r)
	ip := utils.RemoteIp(l.r)
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	claims[ctxdata.CtxKeyJwtActivityId] = ActivityId
	claims[ctxdata.CtxKeyJwtPlatform] = paltform
	claims[ctxdata.CtxKeyJwtIp] = ip
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
