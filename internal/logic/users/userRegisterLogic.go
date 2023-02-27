/*
 * @Author: licat
 * @Date: 2023-02-22 14:28:35
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 14:29:57
 * @Description: licat233@gmail.com
 */
package users

import (
	"context"
	"errors"
	"strings"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (any, error) {
	key := req.SecretKey
	if key != "" {
		return nil, errors.New("secret key is not empty")
	}
	passport := strings.TrimSpace(req.UserPassport)
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

	mdUser := &model.Users{
		ActivityId: activity.Id,
		Name:       "no name",
		Passport:   passport,
	}
	_, err = l.svcCtx.UsersModel.Insert(l.ctx, mdUser)
	if err != nil {
		l.Logger.Errorf("插入用户失败，err:%v", err)
		return nil, errorx.InternalError(err)
	}
	return respx.StateResp("register success", true), nil
}
