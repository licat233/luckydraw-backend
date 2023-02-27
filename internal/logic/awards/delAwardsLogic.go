package awards

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAwardsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAwardsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAwardsLogic {
	return &DelAwardsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAwardsLogic) DelAwards(req *types.DelAwardsReq) (any, error) {
	if err := l.svcCtx.AwardsModel.Delete(l.ctx, req.Id); err != nil {
		l.Logger.Errorf("删除奖品失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
