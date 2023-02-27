package users

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUsersLogic {
	return &AddUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUsersLogic) AddUsers(req *types.AddUsersReq) (any, error) {
	in := &model.Users{
		ActivityId:      req.ActivityId,
		AvailableAwards: req.AvailableAwards,
		Name:            req.Name,
		Passport:        req.Passport,
		Count:           req.Count,
		Total:           req.Total,
	}
	if _, err := l.svcCtx.UsersModel.Insert(l.ctx, in); err != nil {
		l.Logger.Errorf("添加用户失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
