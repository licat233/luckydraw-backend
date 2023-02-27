package activity

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityLogic {
	return &AddActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityLogic) AddActivity(req *types.AddActivityReq) (any, error) {
	in := &model.Activity{
		Uuid:   uuid.NewString(),
		Name:   req.Name,
		Status: req.Status,
	}
	_, err := l.svcCtx.ActivityModel.Insert(l.ctx, in)
	if err != nil {
		l.Logger.Error("failed to insert Activity, error: ", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
