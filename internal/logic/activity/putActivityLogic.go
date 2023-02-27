package activity

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutActivityLogic {
	return &PutActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutActivityLogic) PutActivity(req *types.PutActivityReq) (any, error) {
	in := &model.Activity{
		Id:     req.Id,
		Uuid:   req.Uuid,
		Name:   req.Name,
		Status: req.Status,
	}

	err := l.svcCtx.ActivityModel.Update(l.ctx, in)
	if err != nil {
		l.Logger.Error("failed to update activity, error: ", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
