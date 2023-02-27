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

type PutUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUsersLogic {
	return &PutUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutUsersLogic) PutUsers(req *types.PutUsersReq) (any, error) {
	in := &model.Users{
		Id:              req.Id,
		ActivityId:      req.ActivityId,
		AvailableAwards: req.AvailableAwards,
		Name:            req.Name,
		Passport:        req.Passport,
		Count:           req.Count,
		Total:           req.Total,
	}
	err := l.svcCtx.UsersModel.Update(l.ctx, in)
	if err != nil {
		l.Logger.Error("failed to update user, error: ", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
