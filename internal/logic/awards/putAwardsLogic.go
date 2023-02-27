package awards

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutAwardsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutAwardsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutAwardsLogic {
	return &PutAwardsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutAwardsLogic) PutAwards(req *types.PutAwardsReq) (any, error) {
	in := &model.Awards{
		Id:         req.Id,
		ActivityId: req.ActivityId,
		Uuid:       req.Uuid,
		Grade:      req.Grade,
		Name:       req.Name,
		Image:      req.Image,
		Price:      req.Price,
		Prob:       req.Prob,
		Quantity:   req.Quantity,
		Count:      req.Count,
		IsWin:      req.IsWin,
	}

	err := l.svcCtx.AwardsModel.Update(l.ctx, in)
	if err != nil {
		l.Logger.Error("failed to update award, error: ", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
