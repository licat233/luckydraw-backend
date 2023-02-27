package winningRecords

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelWinningRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelWinningRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelWinningRecordsLogic {
	return &DelWinningRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelWinningRecordsLogic) DelWinningRecords(req *types.DelWinningRecordsReq) (any, error) {
	if err := l.svcCtx.WinningRecordsModel.Delete(l.ctx, req.Id); err != nil {
		l.Logger.Errorf("删除中奖记录失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
