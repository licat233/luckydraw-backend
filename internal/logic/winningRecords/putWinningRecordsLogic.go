package winningRecords

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutWinningRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutWinningRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutWinningRecordsLogic {
	return &PutWinningRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutWinningRecordsLogic) PutWinningRecords(req *types.PutWinningRecordsReq) (any, error) {
	in := &model.WinningRecords{
		Id:         req.Id,
		UserId:     req.UserId,
		AwardId:    req.AwardId,
		ActivityId: req.ActivityId,
		Ip:         req.Ip,
		Platform:   req.Platform,
	}
	err := l.svcCtx.WinningRecordsModel.Update(l.ctx, in)
	if err != nil {
		l.Logger.Error("failed to update WinningRecords, error: ", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
