package users

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/common/utils"
	"luckydraw-backend/internal/logic/awards"
	"luckydraw-backend/internal/logic/tools"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersListLogic {
	return &GetUsersListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersListLogic) GetUsersList(req *types.GetUsersListReq) (any, error) {
	in := &model.Users{
		Id:              req.Id,
		ActivityId:      req.ActivityId,
		AvailableAwards: req.AvailableAwards,
		Name:            req.Name,
		Passport:        req.Passport,
		Count:           req.Count,
		Total:           req.Total,
	}
	list, total, err := l.svcCtx.UsersModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Errorf("获取奖品列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	var data []*types.UserDetail
	var awardLogic = awards.NewSelectAwardsByIdsLogic(l.ctx, l.svcCtx)
	for _, user := range list {
		activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, user.ActivityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询活动失败，err:%v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			activity = nil
		}
		ids := utils.HandlerIdList(user.AvailableAwards)
		availableAwards, err := awardLogic.SelectAwards(ids)
		if err != nil {
			return nil, err
		}
		act := tools.ActivityToResp(activity)
		if act == nil {
			continue
		}
		data = append(data, &types.UserDetail{
			User:            tools.UserToResp(user),
			Activity:        act,
			AvailableAwards: availableAwards,
		})
	}
	return respx.NewListData(data, total, req.PageSize, req.Page), nil
}
