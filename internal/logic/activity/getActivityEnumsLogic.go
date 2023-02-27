package activity

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActivityEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityEnumsLogic {
	return &GetActivityEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActivityEnumsLogic) GetActivityEnums(req *types.GetActivityEnumsReq) (any, error) {
	list, err := l.svcCtx.ActivityModel.FindAll(l.ctx)
	if err != nil {
		l.Logger.Errorf("获取活动列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	data := make([]*types.Enum, 0)
	for _, v := range list {
		data = append(data, &types.Enum{
			Label: v.Name,
			Value: v.Id,
		})
	}
	return respx.NewData(data), nil
}
