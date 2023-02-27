/*
 * @Author: licat
 * @Date: 2023-02-20 23:49:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 12:16:29
 * @Description: licat233@gmail.com
 */
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

type AddAwardsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAwardsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAwardsLogic {
	return &AddAwardsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAwardsLogic) AddAwards(req *types.AddAwardsReq) (any, error) {
	in := &model.Awards{
		ActivityId: req.ActivityId,
		Grade:      req.Grade,
		Name:       req.Name,
		Image:      req.Image,
		Price:      req.Price,
		Prob:       req.Prob,
		Quantity:   req.Quantity,
		Count:      req.Count,
		IsWin:      req.IsWin,
	}
	if _, err := l.svcCtx.AwardsModel.Insert(l.ctx, in); err != nil {
		l.Logger.Errorf("添加奖品失败, err: %v", err)
		return nil, errorx.InternalError(err)
	}
	return respx.New(nil), nil
}
