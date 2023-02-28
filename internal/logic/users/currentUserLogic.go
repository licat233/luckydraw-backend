package users

import (
	"context"
	"strings"

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

	passport := strings.TrimSpace(req.Passport)
	activityUuid := strings.TrimSpace(req.ActivityUuid)
	if activityId < 1 && activityUuid == "" {
		return nil, errorx.New("錯誤的請求參數")
	}
	if passport == "" && userId < 1 {
		return respx.New(NewVisitor(activityId, passport)), nil
	}
	var err error
	var activity *model.Activity
	//如果activityId不存在，则通過activityUuid查询
	if activityId < 1 {
		activity, err = l.svcCtx.ActivityModel.FindOneByUuid(l.ctx, activityUuid)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询活动失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
	} else {
		//否则通过activityId查询
		activity, err = l.svcCtx.ActivityModel.FindOne(l.ctx, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询活动失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
	}
	if err == model.ErrNotFound || activity == nil {
		return nil, errorx.New("活動已不存在")
	}

	if activity.Status != 1 {
		return nil, errorx.New("Sorry！此活動已結束，感謝你的關注，請添加官方客服，獲取最新活動。")
	}

	activityId = activity.Id

	var user *model.Users
	//如果userId不存在，则通过passport查询
	if userId < 1 {
		user, err = l.svcCtx.UsersModel.FindsByPassportAndActivityId(l.ctx, passport, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询用户失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
	} else {
		user, err = l.svcCtx.UsersModel.FindsByIdAndActivityId(l.ctx, userId, activityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询用户失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
	}

	if err == model.ErrNotFound || user == nil {
		//如果没有注册，那就提示抽奖次数为0
		return respx.New(NewVisitor(activityId, passport)), nil
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

func NewVisitor(activityId int64, passport string) *types.PublicUser {
	return &types.PublicUser{
		Id:         0,
		ActivityId: activityId,
		Passport:   passport,
		Count:      0,
		Total:      0,
	}
}
