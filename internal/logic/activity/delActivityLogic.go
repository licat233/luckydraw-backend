package activity

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelActivityLogic {
	return &DelActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelActivityLogic) DelActivity(req *types.DelActivityReq) (any, error) {
	if err := l.svcCtx.ActivityModel.Delete(l.ctx, req.Id); err != nil {
		l.Logger.Errorf("删除活动失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
