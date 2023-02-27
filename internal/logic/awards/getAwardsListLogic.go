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

type GetAwardsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAwardsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAwardsListLogic {
	return &GetAwardsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAwardsListLogic) GetAwardsList(req *types.GetAwardsListReq) (any, error) {
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
	data, total, err := l.svcCtx.AwardsModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Errorf("获取奖品列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	var list []*types.Awards
	for _, v := range data {
		list = append(list, tools.AwardsToResp(v))
	}
	return respx.NewListData(list, total, req.PageSize, req.Page), nil
}
