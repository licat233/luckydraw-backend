package awards

import (
	"context"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/logic/tools"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAwardsByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectAwardsByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectAwardsByIdsLogic {
	return &SelectAwardsByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectAwardsByIdsLogic) SelectAwardsByIds(req *types.SelectAwardsByIdsReq) (any, error) {
	data, err := l.SelectAwards(req.Ids)
	if err != nil {
		return nil, err
	}
	return respx.New(data), nil
}

func (l *SelectAwardsByIdsLogic) SelectAwards(ids []int64) ([]*types.Awards, error) {
	var list []*types.Awards
	if len(ids) == 0 {
		return list, nil
	}
	for _, id := range ids {
		data, err := l.svcCtx.AwardsModel.FindOne(l.ctx, id)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Errorf("查询奖品失败, err: %v", err)
			return nil, errorx.InternalError(err)
		}
		if err == model.ErrNotFound {
			continue
		}
		list = append(list, tools.AwardsToResp(data))
	}
	return list, nil
}
