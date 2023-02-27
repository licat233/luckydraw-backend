package awards

import (
	"context"
	"fmt"

	"luckydraw-backend/common/errorx"
	"luckydraw-backend/common/respx"
	"luckydraw-backend/internal/svc"
	"luckydraw-backend/internal/types"
	"luckydraw-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAwardsEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAwardsEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAwardsEnumsLogic {
	return &GetAwardsEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAwardsEnumsLogic) GetAwardsEnums(req *types.GetAwardsEnumsReq) (any, error) {
	in := &model.Awards{
		Id:         -1,
		ActivityId: req.ParentId,
		Uuid:       "",
		Grade:      "",
		Name:       "",
		Image:      "",
		Price:      -1,
		Prob:       -1,
		Quantity:   -1,
		Count:      -1,
		IsWin:      -1,
	}
	list, _, err := l.svcCtx.AwardsModel.FindList(l.ctx, 0, 0, "", in)
	if err != nil {
		l.Logger.Errorf("获取奖品列表失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	var listEnums []*types.Enum
	for _, v := range list {
		label := fmt.Sprintf("【%s】%s", v.Grade, v.Name)
		listEnums = append(listEnums, &types.Enum{
			Label: label,
			Value: v.Id,
		})
	}
	return respx.New(listEnums), nil
}
