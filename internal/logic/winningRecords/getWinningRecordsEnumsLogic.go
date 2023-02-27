package winningRecords

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWinningRecordsEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWinningRecordsEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWinningRecordsEnumsLogic {
	return &GetWinningRecordsEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWinningRecordsEnumsLogic) GetWinningRecordsEnums(req *types.GetWinningRecordsEnumsReq) (any, error) {
	list, _, err := l.svcCtx.WinningRecordsModel.FindList(l.ctx, 0, 0, "", nil)
	if err != nil {
		l.Logger.Errorf("获取中奖记录列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	var listEnums []*types.Enum
	for _, v := range list {
		listEnums = append(listEnums, &types.Enum{
			Label: v.Ip,
			Value: v.Id,
		})
	}
	return respx.New(listEnums), nil
}
