package users

import (
	"context"

	"luckydraw-backend/common/ctxdata"
	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentUserLogic {
	return &CurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentUserLogic) CurrentUser(req *types.CurrentUserReq) (any, error) {
	// 一、先判断该用户是否已经注册
	userId := ctxdata.GetUidFromCtx(l.ctx)
	activityId := ctxdata.GetAidFromCtx(l.ctx)
	if activityId == 0 {
		activity, err := l.svcCtx.ActivityModel.FindOneByUuid(l.ctx, req.ActivityUuid)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询活动失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			activityId = 0
		} else {
			activityId = activity.Id
		}
	}
	if activityId == 0 {
		return nil, nil
	}
	isRegistered := false
	var err error
	var user *model.Users
	if userId == 0 && req.Passport != "" && activityId != 0 {
		user, err = l.svcCtx.UsersModel.FindsByPassportAndActivityId(l.ctx, req.Passport, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询用户失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			userId = 0
		} else {
			isRegistered = true
			userId = user.Id
		}
	}
	if userId == 0 {
		//那就提示网络拥堵，抽奖人数过多，请稍后再试
		return nil, nil
	}

	if user == nil {
		user, err = l.svcCtx.UsersModel.FindsByIdAndActivityId(l.ctx, userId, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询用户失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		isRegistered = err == nil
	}
	if !isRegistered {
		return nil, nil
	}
	userInfo := &types.PublicUser{
		Id:         user.Id,
		ActivityId: user.ActivityId,
		Passport:   user.Passport,
		Count:      user.Count,
		Total:      user.Total,
	}
	return respx.SingleResp("success", userInfo), nil
}
